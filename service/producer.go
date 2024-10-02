package service

import (
	"context"
	"github.com/IBM/sarama"
	"go-micro.dev/v4/util/log"
	"productCatalog/global"
	"strconv"
)

type IProducerService interface {
	SendGeneralEvent([]byte) error
}

type ProducerService struct {
	producer sarama.SyncProducer
	ctx      context.Context
}

func NewProducerService(producer sarama.SyncProducer) IProducerService {
	return &ProducerService{
		//consumerGroup: consumerGroup,
		producer: producer,
		ctx:      context.Background(),
	}
}

func (s ProducerService) SendGeneralEvent(message []byte) error {
	// Publish the message to Kafka
	partition, offset, err := s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: global.KafkaGeneralEventTopic,
		Value: sarama.StringEncoder(message),
	})
	if err != nil {
		log.Fatalf("Failed to send message to Kafka: %v", err)
	}
	log.Info("Message is stored in topic /" + global.KafkaGeneralEventTopic + "/partition/" + string(partition) + "/offset/" + strconv.FormatInt(offset, 10))
	return err
}
