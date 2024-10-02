package handler

import (
	"context"
	"github.com/IBM/sarama"
	"productCatalog/service"
)

type ProductCatalog struct {
	MerchandiseService service.IMerchandiseService
	Producer           service.IProducerService
	Consumer           *Consumer
}

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
	ctx           context.Context
}

func NewConsumer(consumerGroup sarama.ConsumerGroup) *Consumer {
	return &Consumer{
		consumerGroup: consumerGroup,
		ctx:           context.Background(),
	}
}
