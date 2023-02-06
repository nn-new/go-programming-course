package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
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
		Topic:    "words",
		GroupID:  "consumer-group-words-1",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		Dialer:   dialer,
	})

	defer reader.Close()

	mapWord := map[string]int{}

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

			words := []string{}
			err = json.Unmarshal(recordValue, &words)
			if err != nil {
				fmt.Println(err.Error())
			}

			for _, v := range words {
				if v != "" {
					rd := strings.ReplaceAll(v, ".", "")
					rb := strings.ReplaceAll(rd, "\n", "")
					mapWord[rb] = mapWord[rb] + 1
				}
			}

			for k, v := range mapWord {
				fmt.Printf("%s : %d\n", k, v)
			}
		}
	}
}
