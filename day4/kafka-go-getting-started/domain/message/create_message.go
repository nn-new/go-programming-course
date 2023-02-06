package message

import (
	"encoding/json"
	"errors"
	"kafka-go-getting-started/domain/shop"

	"github.com/segmentio/kafka-go"
)

func CreateMessageNoKey(input []string) []kafka.Message {
	msgs := []kafka.Message{}
	for _, v := range input {
		msgs = append(msgs, kafka.Message{
			Value: []byte(v),
		})
	}
	return msgs
}

func CreateMessageWithMap(input map[string]string) []kafka.Message {
	msgs := []kafka.Message{}
	for k, v := range input {
		msgs = append(msgs, kafka.Message{
			Key:   []byte(k),
			Value: []byte(v),
		})
	}
	return msgs
}

func CreateMessageWithDupKey(keys []string, values []string) ([]kafka.Message, error) {
	if len(keys) != len(values) {
		return nil, errors.New("keys/values length not match")
	}
	msgs := []kafka.Message{}
	for i, k := range keys {
		msgs = append(msgs, kafka.Message{
			Key:   []byte(k),
			Value: []byte(values[i]),
		})
	}
	return msgs, nil
}

func CreateOrderMessage(o shop.Order) (kafka.Message, error) {
	out, err := json.Marshal(&o)
	if err != nil {
		return kafka.Message{}, nil
	}
	return kafka.Message{
		Key:   []byte(o.OrderID.String()),
		Value: out,
	}, nil
}
