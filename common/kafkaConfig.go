package common

import (
	"encoding/json"
	capi "github.com/hashicorp/consul/api"
	log "go-micro.dev/v4/logger"
)

type KafkaConfig struct {
	User    string   `json:"user,omitempty"`
	Pwd     string   `json:"password,omitempty"`
	Brokers []string `json:"brokers,omitempty"`
}

func CapiGetKafkaConfig(client *capi.Client, path string) *KafkaConfig {
	val, _, err := client.KV().Get(path, nil)
	if err != nil {
		log.Error(err)
	}

	var kafkaInfo KafkaConfig
	err = json.Unmarshal(val.Value, &kafkaInfo)
	if err != nil {
		log.Error(err)
	}
	return &kafkaInfo
}
