package producer_user_create

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	producer *kafka.Writer
}

func New() *Producer {
	producer := &kafka.Writer{
		Addr:         kafka.TCP("0.0.0.0:9092"),
		Topic:        "user.create",
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
	return &Producer{
		producer: producer,
	}
}

func (p *Producer) Create(ctx context.Context, message []byte) {
	err := p.producer.WriteMessages(ctx, kafka.Message{
		Value: message,
	})
	if err != nil {
		log.Println("failed to: ", err)
	}
}
