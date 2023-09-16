package main

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type HashRateReport struct {
	attempts      uint64
	secondElapsed float64
}

type Miner struct {
	Argon2ID

	account string

	httpClient *http.Client
	ctx        context.Context
	sync.RWMutex

	attempts uint64

	startTime                time.Time
	memoryUpdate             chan uint32
	initMemoryDifficultyDone chan struct{}
}

type MinerOption func(m *Miner)

func NewMiner(account string, opts ...MinerOption) *Miner {
	m := &Miner{
		httpClient:               http.DefaultClient,
		ctx:                      context.Background(),
		Argon2ID:                 GetDefaultArgon2ID(),
		account:                  account,
		memoryUpdate:             make(chan uint32),
		initMemoryDifficultyDone: make(chan struct{}),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (m *Miner) Banner() {
	logrus.Infof("--------User Configuration--------")
	logrus.Infof("Memory Difficulty: %d", m.memory)
	logrus.Infof("Time Difficulty: %d", m.time)
	logrus.Infof("Threads: %d", m.thread)
	logrus.Infof("account: %s", m.account)
	logrus.Infof("salt: %s", string(m.salt))
	logrus.Infof("keyLength: %d", m.keyLen)
	logrus.Infof("----------------------------------")
}

type MemoryDifficultyResponse struct {
	Difficulty string `json:"difficulty"`
}

func (m *Miner) Background() {
	tick := time.NewTicker(time.Second * 5)

	go func() {
		first := true
		for {
			select {
			case <-m.ctx.Done():
				logrus.Infof("G.Background exited by context signal")
				return
			default:
				memoryDifficulty, err := GetLatestMemoryDifficulty()
				if err != nil {
					logrus.Errorf("failed to fetch latest memory difficulty: %+v", err)
					continue
				}
				if m.HasMemoryDifficultyUpdate(memoryDifficulty) {
					m.setMemoryDifficulty(memoryDifficulty)
					if first {
						first = false
						logrus.Infof("init memory difficulty from api done, start mining...")
						close(m.initMemoryDifficultyDone)
						m.startTime = time.Now()
					}
				}

				<-time.After(time.Second * 5)
			}
		}
	}()

	for {
		select {
		case <-m.ctx.Done():
			logrus.Infof("G.Background exited by context signal")
			return
		case <-tick.C:
			logrus.Infof("speed: %.2f hash/s", m.GetHashRate())
		case memory := <-m.memoryUpdate:
			if m.HasMemoryDifficultyUpdate(memory) {
				logrus.Infof("update memory by api: %d", memory)
				m.setMemoryDifficulty(memory)
			}
		}
	}
}

func (m *Miner) setMemoryDifficulty(memoryDifficulty uint32) {
	m.Lock()
	m.memory = memoryDifficulty
	m.Unlock()
}

func (m *Miner) HasMemoryDifficultyUpdate(memoryDifficultyToValidate uint32) bool {
	m.RLock()
	update := m.memory != memoryDifficultyToValidate
	m.RUnlock()

	return update
}

func (m *Miner) solve(id Argon2ID, name string) {
	key := NewRandomKey()

	for attempts := uint64(0); ; {
		attempts++
		key.AddOne()

		hash := id.HashStr(key.JackBytes())
		if strings.Contains(hash, defaultTargetXEN11) {
			m.ReportHashRate(attempts % hashRateReportInterval)
			// use goroutine to start next block instantly
			go func() {
				logrus.Infof("%s hex: %s", name, key.Hex())
				logrus.Infof("%s signature: %s", name, hash)
				err := Verify(Payload{
					HashToVerify: id.Signature(key.JackBytes()),
					Key:          key.Hex(),
					Account:      m.account,
				})
				if err != nil {
					logrus.Errorf("failed to verify: %+v", err)
				}
			}()
			return
		}

		if attempts%hashRateReportInterval == 0 {
			m.ReportHashRate(hashRateReportInterval)

			select {
			case <-m.ctx.Done():
				return
			default:
				if m.HasMemoryDifficultyUpdate(id.memory) {
					return
				}
			}
		}
	}
}

func (m *Miner) ReportHashRate(attempts uint64) {
	m.Lock()
	m.attempts += attempts
	m.Unlock()
}

func (m *Miner) Start() {
	go m.Background()

	f := func(idx int) {
		<-m.initMemoryDifficultyDone

		for true {
			id := m.GetArgon2ID()
			name := fmt.Sprintf("p%d", idx)

			if !id.Valid() {
				logrus.Infof("[%s]waiting for memory difficulty init", name)
				<-time.After(time.Second)
				continue
			}

			m.solve(id, name)
		}
	}

	logrus.Infof("start %d process...", runtime.NumCPU())
	for i := 0; i < goroutineCount; i++ {
		go f(i)
	}
}

func (m *Miner) GetArgon2ID() Argon2ID {
	m.RLock()
	m.RUnlock()

	return m.Argon2ID
}

func (m *Miner) GetHashRate() float64 {
	m.RLock()
	defer m.RUnlock()

	cost := time.Now().Sub(m.startTime)
	speed := float64(m.attempts) / cost.Seconds()
	return speed
}
