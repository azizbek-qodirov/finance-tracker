package kfk

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func ProduceKafkaMessage(topic string, key string, message interface{}, kafkaBroker string) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message to JSON: %v", err)
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: messageBytes,
	}

	err = w.WriteMessages(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("failed to write message to Kafka: %v", err)
	}

	fmt.Printf("Kafka message produced successfully to topic: %s with key: %s \n", topic, key)

	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to close Kafka writer: %v", err)
	}

	return nil
}
