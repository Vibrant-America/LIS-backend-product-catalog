package service

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"
	"github.com/stripe/stripe-go/v74"
	"productCatalog/ent"
	"productCatalog/model"
	pb "productCatalog/proto"

	"github.com/grissom1/go-micro/v4/util/log"
	"github.com/stripe/stripe-go/v74/tax/calculation"
)

type IStripeService interface {
	CalculateTax(*model.CalculateTaxRequest) (*stripe.TaxCalculation, error)
	CalculateTaxPb(md *pb.CalculateTaxRequest) (*stripe.TaxCalculation, error)
}

type StripeService struct {
	client *ent.Client
	cache  *redis.Client
	ctx    context.Context
	apiKey string
}

func NewStripeService(client *ent.Client, cache *redis.Client, apiKey string) IStripeService {
	stripe.Key = apiKey
	return &StripeService{
		client: client,
		cache:  cache,
		ctx:    context.Background(),
		apiKey: apiKey,
	}
}

func (s StripeService) CalculateTaxPb(md *pb.CalculateTaxRequest) (*stripe.TaxCalculation, error) {
	params := &stripe.TaxCalculationParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		CustomerDetails: &stripe.TaxCalculationCustomerDetailsParams{
			Address: &stripe.AddressParams{
				Country:    stripe.String(md.CustomerDetails.Country),
				State:      stripe.String(md.CustomerDetails.State),
				PostalCode: stripe.String(md.CustomerDetails.PostalCode),
				City:       stripe.String(md.CustomerDetails.City),
				Line1:      stripe.String(md.CustomerDetails.Line1),
			},
			AddressSource: stripe.String(md.CustomerDetails.AddressType),
		},
		LineItems:    []*stripe.TaxCalculationLineItemParams{},
		ShippingCost: &stripe.TaxCalculationShippingCostParams{Amount: stripe.Int64(int64(md.ShippingFee * 100))},
	}
	for _, item := range md.LineItems {
		params.LineItems = append(params.LineItems, &stripe.TaxCalculationLineItemParams{
			Amount:      stripe.Int64(int64(item.Amount * 100)),
			Quantity:    stripe.Int64(item.Quantity),
			Reference:   stripe.String(item.Reference),
			TaxCode:     stripe.String(item.TaxCode),
			TaxBehavior: stripe.String(item.TaxBehavior),
		})
	}
	params.AddExpand("line_items")

	result, err := calculation.New(params)
	if err != nil {
		log.Error(err.Error())
		sentry.CaptureMessage(err.Error())
		return nil, err
	}
	return result, nil
}

func (s StripeService) CalculateTax(md *model.CalculateTaxRequest) (*stripe.TaxCalculation, error) {
	params := &stripe.TaxCalculationParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		CustomerDetails: &stripe.TaxCalculationCustomerDetailsParams{
			Address: &stripe.AddressParams{
				Country:    stripe.String(md.CustomerDetails.Country),
				State:      stripe.String(md.CustomerDetails.State),
				PostalCode: stripe.String(md.CustomerDetails.PostalCode),
				City:       stripe.String(md.CustomerDetails.City),
				Line1:      stripe.String(md.CustomerDetails.Line1),
			},
			AddressSource: stripe.String(md.CustomerDetails.AddressType),
		},
		LineItems:    []*stripe.TaxCalculationLineItemParams{},
		ShippingCost: &stripe.TaxCalculationShippingCostParams{Amount: stripe.Int64(int64(md.ShippingFee * 100))},
	}
	for _, item := range md.LineItems {
		params.LineItems = append(params.LineItems, &stripe.TaxCalculationLineItemParams{
			Amount:      stripe.Int64(int64(item.Amount * 100)),
			Quantity:    stripe.Int64(item.Quantity),
			Reference:   stripe.String(item.Reference),
			TaxCode:     stripe.String(item.TaxCode),
			TaxBehavior: stripe.String(item.TaxBehavior),
		})
	}
	params.AddExpand("line_items")

	result, err := calculation.New(params)
	if err != nil {
		log.Error(err.Error())
		sentry.CaptureMessage(err.Error())
		return nil, err
	}
	return result, nil
}
