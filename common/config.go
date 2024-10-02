package common

import (
	"encoding/json"
	capi "github.com/hashicorp/consul/api"
	log "go-micro.dev/v4/logger"
)

// !Consul config for go micro consul config plugin has to be JSON format. This is not required by Consul, but by Go-Micro framework
//func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
//	consulSource := consul.NewSource(
//		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
//		consul.WithPrefix(prefix),
//		//default prefix /micro/config
//		consul.StripPrefix(true),
//		consul.WithToken(KuberEnv.ConsulToken),
//	)
//	consulConfig, err := config.NewConfig()
//	if err != nil {
//		return consulConfig, err
//	}
//	err = consulConfig.Load(consulSource)
//	return consulConfig, err
//}

func CapiGetSecrets(client *capi.Client, path string) {
	val, _, err := client.KV().Get(path, nil)
	if err != nil {
		log.Error(err)
	}

	var mySecret struct {
		Secret             string `json:"secret"`
		SecretStaging      string `json:"secret_staging"`
		InternalJWT        string `json:"internalJWT"`
		StagingInternalJWT string `json:"stagingInternalJWT"`
		StripeLiveKeyVA    string `json:"stripeLiveKeyVA"`
		StripeTestKeyVA    string `json:"stripeTestKeyVA"`
		StripeLiveKeyZb    string `json:"stripeLiveKeyZb"`
		StripeTestKeyZb    string `json:"stripeTestKeyZb"`
	}

	err = json.Unmarshal(val.Value, &mySecret)
	if err != nil {
		log.Error(err)
	}

	if KuberEnv.RunEnv == "Aks" || KuberEnv.RunEnv == "Production" {
		JWTSecret = mySecret.Secret
		InternalJWT = mySecret.InternalJWT

		StripeKeyZb = mySecret.StripeLiveKeyZb
		StripeKeyVA = mySecret.StripeLiveKeyVA
	} else {
		JWTSecret = mySecret.SecretStaging
		InternalJWT = mySecret.StagingInternalJWT

		StripeKeyZb = mySecret.StripeTestKeyZb
		StripeKeyVA = mySecret.StripeTestKeyVA
	}

}
