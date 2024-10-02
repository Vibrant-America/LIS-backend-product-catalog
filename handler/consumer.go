package handler

import (
	"errors"
	"github.com/IBM/sarama"
	"github.com/getsentry/sentry-go"
	log "go-micro.dev/v4/logger"
	"productCatalog/global"
)

func (p ProductCatalog) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (p ProductCatalog) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (p ProductCatalog) ConsumeClaim(_ sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for _ = range claim.Messages() {
		// DON'T LET ANY ERROR INTERRUPT THIS FOR LOOP!!!!!!!!!!!!!
	}
	return nil
}

func (p ProductCatalog) StartConsumer() {
	defer p.Consumer.consumerGroup.Close()
	for {
		err := p.Consumer.consumerGroup.Consume(p.Consumer.ctx, []string{global.KafkaGeneralEventTopic}, p)
		if err != nil {
			log.Fatalf("Error in consumer group: %v", err)
			sentry.CaptureException(errors.New("Error in consumer group: " + err.Error()))
		}
	}
}
