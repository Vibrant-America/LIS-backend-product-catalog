package global

import (
	"os"
	"sync"
	"time"
)

func PreserveTime() time.Time { return time.Date(1970, 1, 1, 0, 0, 1, 0, &time.Location{}) } // The first second in Epoch to represent deleted time

const LisServiceBase = "https://www.vibrant-america.com/lisapi"

const (
	TaskCodeTest = "test-"
)

var ChargingServiceBase string
var OrderServiceBase string
var LisCoreServiceBase string
var KafkaGeneralEventTopic string

const KafkaServiceAddr = "192.168.60.9:9095"

var StripeMutex sync.Mutex

func InitServiceAddress() {
	if os.Getenv("BILLING_ENV") == "Aks" {
		KafkaGeneralEventTopic = "lis-general-events"
		ChargingServiceBase = LisServiceBase + "/v1/charging"
		OrderServiceBase = "http://192.168.60.6/v1/portal/order/"
		LisCoreServiceBase = "192.168.60.6:30276"
	} else if os.Getenv("BILLING_ENV") == "Production" {
		KafkaGeneralEventTopic = "lis-general-events"
		ChargingServiceBase = LisServiceBase + "/v1/charging"
		OrderServiceBase = "http://192.168.60.6/v1/portal/order/"
		LisCoreServiceBase = "192.168.60.6:30276"
	} else {
		KafkaGeneralEventTopic = "staging_lis-general-events"
		ChargingServiceBase = LisServiceBase + "/v1/charging/staging"
		OrderServiceBase = "http://192.168.60.6/v1/portal/order/"
		LisCoreServiceBase = "192.168.60.6:30280"
	}
}
