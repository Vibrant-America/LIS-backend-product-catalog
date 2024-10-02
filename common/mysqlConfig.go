package common

import (
	"encoding/json"
	capi "github.com/hashicorp/consul/api"
	log "go-micro.dev/v4/logger"
)

type MySQLConfig struct {
	Host              string `json:"host,omitempty"`
	User              string `json:"user,omitempty"`
	Pwd               string `json:"password,omitempty"`
	Database          string `json:"database,omitempty"`
	Port              int    `json:"port,omitempty"`
	SecondaryHost     string `json:"secondary_host,omitempty"`
	SecondaryUser     string `json:"secondary_user,omitempty"`
	SecondaryPwd      string `json:"secondary_password,omitempty"`
	SecondaryDatabase string `json:"secondary_database,omitempty"`
	SecondaryPort     int    `json:"secondary_port,omitempty"`
	ThirdHost         string `json:"third_host,omitempty"`
	ThirdUser         string `json:"third_user,omitempty"`
	ThirdPwd          string `json:"third_password,omitempty"`
	ThirdDatabase     string `json:"third_database,omitempty"`
	ThirdPort         int    `json:"third_port,omitempty"`
	AutoMigration     bool   `json:"autoMigration,omitempty"`
}

//func GetMySqlConfigFromConsul(config config.Config, path ...string) *MySQLConfig {
//	mysqlConfig := &MySQLConfig{}
//	err := config.Get(path...).Scan(mysqlConfig)
//	log.Info("consule config:", string(config.Get(path...).Bytes()))
//	if err != nil {
//		log.Fatal("reading consul config fail: ", path)
//	}
//	return mysqlConfig
//}

func CapiGetMySqlConfig(client *capi.Client, path string) *MySQLConfig {
	val, _, err := client.KV().Get(path, nil)
	if err != nil {
		log.Error(err)
	}

	var mysqlInfo MySQLConfig
	err = json.Unmarshal(val.Value, &mysqlInfo)
	if err != nil {
		log.Error(err)
	}
	return &mysqlInfo
}
