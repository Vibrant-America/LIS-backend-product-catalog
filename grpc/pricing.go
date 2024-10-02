package grpc

import (
	"github.com/getsentry/sentry-go"
	"productCatalog/service"
	"time"
)

type PricingRpcHandler struct {
	TaxService service.IStripeService
}

func panicRecover() {
	err := recover()
	if err != nil {
		sentry.CurrentHub().Recover(err)
		sentry.Flush(time.Second * 5)
		panic(err)

	}
}
