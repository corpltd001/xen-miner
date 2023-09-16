package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
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

func init() {
	initLogrus()

	flag.StringVar(&account, "account", defaultAccount, "beneficiary address")
	flag.Parse()
}

func main() {
	miner := NewMiner(account)
	miner.memory = 0
	logrus.Infof("start miner: %s", miner.account)

	go miner.Start()

	select {}
}
