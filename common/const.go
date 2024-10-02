package common

var JWTSecret string
var InternalJWT string
var StripeKeyZb string
var StripeKeyVA string

//var ConsulSecret struct {
//	Secret             string `json:"secret"`
//	SecretStaging      string `json:"secret_staging"`
//	InternalJWT        string `json:"internalJWT"`
//	StagingInternalJWT string `json:"stagingInternalJWT"`
//	StripeLiveKeyVA    string `json:"stripeLiveKeyVA"`
//	StripeTestKeyVA    string `json:"stripeTestKeyVA"`
//	StripeLiveKeyZb    string `json:"stripeLiveKeyZb"`
//	StripeTestKeyZb    string `json:"stripeTestKeyZb"`
//}
//
//func GetConsulSecret(config config.Config, path string) {
//	err := config.Get(path).Scan(&ConsulSecret)
//	if err != nil {
//		log.Error(err)
//	}
//}
