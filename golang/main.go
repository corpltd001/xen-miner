package main

import (
	"flag"
	"os"

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
var threads int

func init() {
	initLogrus()

	flag.StringVar(&account, "account", defaultAccount, "beneficiary address")
	flag.IntVar(&threads, "threads", systemThreadCount, "Number of threads")
	flag.Parse()
}

func main() {
	miner := NewMiner(account, threads)
	miner.memory = 0
	logrus.Infof("start miner: %s", miner.account)

	go miner.Start()

	select {}
}
