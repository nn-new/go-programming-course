package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"wordcount/config"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func main() {
	cfg := config.LoadConfig(".")
	// Set up a channel for handling Ctrl-C, etc
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

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

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{cfg.KafkaURL},
		Topic:    "paragraphs",
		GroupID:  "consumer-group-paragraphs-1",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Dialer:   dialer,
	})

	defer reader.Close()

	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{cfg.KafkaURL},
		Topic:    "words",
		Dialer:   dialer,
		Balancer: &kafka.Hash{},
	})

	defer func() {
		if err := kafkaWriter.Close(); err != nil {
			log.Fatalf("fatal error to close writer: %s\n", err)
		}
	}()

	run := true
	for run {
		select {
		case sig := <-quit:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			localCtx, localCancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer localCancel()
			m, err := reader.ReadMessage(localCtx)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Printf("Consumed event from topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			recordValue := m.Value

			words, err := json.Marshal(strings.Split(string(recordValue), " "))
			if err != nil {
				fmt.Println(err.Error())
			}

			msg := kafka.Message{
				Key:   m.Key,
				Value: words,
			}

			err = kafkaWriter.WriteMessages(localCtx, msg)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
