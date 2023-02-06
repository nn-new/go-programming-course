package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"kafka-go-getting-started/config"
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

	conn, err := dialer.DialContext(ctx, "tcp", cfg.KafkaURL)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		// fmt.Printf("topic: %s, partition: %d\n", p.Topic, p.ID)
		m[p.Topic] = struct{}{}
	}

	for k := range m {
		fmt.Println(k)
	}
}
