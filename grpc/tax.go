package grpc

import (
	"context"
	pb "productCatalog/proto"

	log "go-micro.dev/v4/logger"
)

func (e *PricingRpcHandler) CalculateTax(ctx context.Context, in *pb.CalculateTaxRequest, out *pb.CalculateTaxResponse) error {
	log.Info("handler - CalculateTax")
	defer panicRecover()

	result, err := e.TaxService.CalculateTaxPb(in)
	if err != nil {
		return err
	}

	out.Id = result.ID
	out.AmountTotal = float64(result.AmountTotal) / 100.
	out.TaxAmount = float64(result.TaxAmountExclusive) / 100.
	return nil
}
