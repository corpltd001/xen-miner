package main

import (
	"encoding/binary"
	"encoding/hex"
	"math"
	"math/rand"
)

const KeyLength = 64

type Key struct {
	data [KeyLength]byte
}

func (k *Key) Hex() string {
	return hex.EncodeToString(k.data[:])
}

func (k *Key) SetBytes(data []byte) {
	if len(data) > KeyLength {
		data = data[:KeyLength]
	}
	k.data = [KeyLength]byte{}
	copy(k.data[:], data[:])
}

func (k *Key) JackBytes() []byte {
	return []byte(k.Hex())
}

func (k *Key) AddOne() {
	for i := KeyLength - 1; i >= 0; i-- {
		carry := k.data[i] == math.MaxUint8
		k.data[i]++

		if !carry {
			return
		}
	}
}

func NewRandomKey() *Key {
	value := make([]byte, KeyLength)
	for i := 0; i < KeyLength/8; i++ {
		r := rand.Uint64()
		binary.BigEndian.PutUint64(value[i*8:(i+1)*8], r)
	}

	ans := &Key{}
	ans.SetBytes(value)

	return ans
}
