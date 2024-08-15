package kfk

import (
	"context"
	"fmt"
	"log"

	"budget-service/storage"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader  *kafka.Reader
	Storage storage.StorageI
}

func NewKafkaConsumer(kafkaBroker, topic, groupID string, storage storage.StorageI) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return &KafkaConsumer{Reader: reader, Storage: storage}
}

func (kc *KafkaConsumer) ConsumeMessages(ctx context.Context) {
	for {
		m, err := kc.Reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Error reading Kafka message: %v", err)
			break
		}

		fmt.Printf("Message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		switch m.Topic {
		case "account-updates":
			kc.handleAccountUpdate(m.Value)
		case "account-balance-updates":
			kc.handleAccountBalanceUpdate(m.Value)
		case "budget-create":
			kc.handleBudgetCreate(m.Value)
		case "budget-update":
			kc.handleBudgetUpdate(m.Value)
		case "budget-delete":
			kc.handleBudgetDelete(m.Value)
		case "category-create":
			kc.handleCategoryCreate(m.Value)
		case "category-update":
			kc.handleCategoryUpdate(m.Value)
		case "category-delete":
		case "goal-create":
			kc.handleGoalCreate(m.Value)
		case "goal-update":
			kc.handleGoalUpdate(m.Value)
		case "goal-delete":
			kc.handleGoalDelete(m.Value)
			kc.handleCategoryDelete(m.Value)
		case "transaction-create":
			kc.handleTransactionCreate(m.Value)
		case "transaction-delete":
			kc.handleTransactionDelete(m.Value)
		}
	}
}

func (kc *KafkaConsumer) Close() error {
	return kc.Reader.Close()
}
