package main

import "runtime"

const (
	defaultAccount = "0xF120007d00480034fAf40000e1727C7809734b20"

	defaultTargetXEN11      = "XEN11"
	defaultTimeDifficulty   = 1
	defaultMemoryDifficulty = 8
	defaultThread           = 1
	defaultKeyLength        = 64

	hashRateReportInterval = 10_000
)

var (
	defaultSalt = []byte("XEN10082022XEN")

	goroutineCount = runtime.NumCPU()
)
