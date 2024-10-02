package common

import (
	"os"

	log "go-micro.dev/v4/logger"
)

var KuberEnv struct {
	RunEnv           string //Staging or Production
	ConsulConfigAddr string
	ConsulPrefix     string
	ConsulToken      string
}

func InitKuberEnv() {
	KuberEnv.RunEnv = os.Getenv("BILLING_ENV")
	KuberEnv.ConsulConfigAddr = os.Getenv("CONSUL_ADDR")
	KuberEnv.ConsulPrefix = os.Getenv("CONSUL_PREFIX")
	KuberEnv.ConsulToken = os.Getenv("CONSUL_TOKEN")
	log.Info("Env:", KuberEnv.RunEnv)
	log.Info("Consul Addr:", KuberEnv.ConsulConfigAddr)
	log.Info("Consul Prefix:", KuberEnv.ConsulPrefix)
	log.Info("Consul Token:", KuberEnv.ConsulToken)
}
