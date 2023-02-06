package main

import (
	"context"
	"crypto/tls"
	"kafka-go-getting-started/config"
	"kafka-go-getting-started/domain/message"
	"log"
	"time"

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
		Topic:    "notification",
		Dialer:   dialer,
		Balancer: &kafka.Hash{},
	})

	defer func() {
		if err := kafkaWriter.Close(); err != nil {
			log.Fatalf("fatal error to close writer: %s\n", err)
		}
	}()

	// Example Send message with JSON
	// {
	// 	"ordertime": 1497014222380,
	// 	"orderid": 18,
	// 	"itemid": "Item_184",
	// 	"address": {
	// 		"city": "Mountain View",
	// 		"state": "CA",
	// 		"zipcode": 94041
	// 	}
	// }

	// msg, err := message.CreateOrderMessage(shop.MockOrder)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// err = kafkaWriter.WriteMessages(ctx, msg)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	msgs := message.CreateMessageNoKey([]string{"This is 1st.", "This is 2nd"})
	// msgs := message.CreateMessageWithMap(map[string]string{
	// 	"1": "This is 1st.",
	// 	"2": "This is 2nd",
	// })

	// msgs, err := message.CreateMessageWithDupKey(
	// 	[]string{"3", "3"},
	// 	[]string{"add to cart", "order confirm"},
	// )
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := kafkaWriter.WriteMessages(ctx, msgs...)
	if err != nil {
		log.Fatalln(err)
	}
}
