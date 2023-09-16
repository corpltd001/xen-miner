package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

func initLogrus() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
	logrus.Infoln("init logger done")
}

var account string
var workerCount int

func init() {
	initLogrus()

	flag.StringVar(&account, "account", defaultAccount, "Beneficiary address")
	flag.IntVar(&workerCount, "worker", runtime.NumCPU(), "Number of thread to miner, use all cpu core if not set")
	flag.Parse()

	if workerCount <= 0 {
		logrus.Fatalf("invalid worker setting: %d, should >= 1", workerCount)
	}
}

func main() {
	miner := NewMiner(account, MinerOptionWithCount(workerCount))
	miner.memory = 0
	logrus.Infof("start miner: %s", miner.account)

	go miner.Start()

	select {}
}
