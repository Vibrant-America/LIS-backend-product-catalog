package common

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	capi "github.com/hashicorp/consul/api"
	//"go-micro.dev/v4/config"
	log "go-micro.dev/v4/logger"
)

//func GetRedisConfigFromConsul(config config.Config, path ...string) *redis.Client {
//	redisConfig := &RedisConfig{}
//	err := config.Get(path...).Scan(redisConfig)
//	if err != nil {
//		log.Fatalf("failed to get redis config: %v", err)
//	}
//
//	return redis.NewFailoverClient(&redis.FailoverOptions{
//		MasterName:    redisConfig.MasterName,
//		SentinelAddrs: redisConfig.Nodes,
//	})
//}

type RedisConfig struct {
	MasterName       string   `json:"master_name"`
	SentinelAddrs    []string `json:"sentinel_addrs"`
	Password         string   `json:"password,omitempty"`
	SentinelPassword string   `json:"sentinel_password,omitempty"`
}

var RedisParameters RedisConfig

func CapiGetRedisConfig(client *capi.Client, path string) *redis.Client {

	val, _, err := client.KV().Get(path, nil)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(val.Value, &RedisParameters)
	if err != nil {
		log.Error(err)
	}
	redisOpt := redis.FailoverOptions{
		MasterName:    RedisParameters.MasterName,
		SentinelAddrs: RedisParameters.SentinelAddrs,
	}
	if KuberEnv.RunEnv == "Aks" || KuberEnv.RunEnv == "Akstaging" {
		redisOpt.Password = RedisParameters.Password
		redisOpt.SentinelPassword = RedisParameters.SentinelPassword
	}

	return redis.NewFailoverClient(&redisOpt)
}
