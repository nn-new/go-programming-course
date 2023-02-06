package message

import "github.com/segmentio/kafka-go"

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
