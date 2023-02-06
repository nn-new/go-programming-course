package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"
	"wordcount/config"
	"wordcount/domain/message"

	lorem "github.com/drhodes/golorem"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func main() {
	cfg := config.LoadConfig(".")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mechanism := plain.Mechanism{
		Username: cfg.KafkaUser,
		Password: cfg.KafkaPassword,
	}

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.KafkaURL},
		Topic:    "paragraphs",
		Dialer:   dialer,
		Balancer: &kafka.Hash{},
	})

	defer func() {
		if err := kafkaWriter.Close(); err != nil {
			log.Fatalf("fatal error to close writer: %s\n", err)
		}
	}()

	paragraphs := map[string]string{}
	for i := 0; i < 5; i++ {
		paragraph := lorem.Paragraph(10, 100)
		paragraphs[paragraph[:10]] = paragraph
	}

	msgs := message.CreateMessageWithMap(paragraphs)

	err := kafkaWriter.WriteMessages(ctx, msgs...)
	if err != nil {
		log.Fatalln(err)
	}
}
