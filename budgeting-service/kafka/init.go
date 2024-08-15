package kfk

import (
	cf "budget-service/config"
	"budget-service/storage"
	"context"
)

func InitKafka(config *cf.Config, db *storage.Storage) {
	// %--------- Account consumers
	accountUpdateConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "account-updates", "account-update-group", db)
	go accountUpdateConsumer.ConsumeMessages(context.Background())
	defer accountUpdateConsumer.Close()

	accountBalanceConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "account-balance-updates", "account-balance-update-group", db)
	go accountBalanceConsumer.ConsumeMessages(context.Background())
	defer accountBalanceConsumer.Close()

	// %--------- Budget consumers
	budgetCreateConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "budget-create", "budget-create-group", db)
	go budgetCreateConsumer.ConsumeMessages(context.Background())
	defer budgetCreateConsumer.Close()

	budgetUpdateConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "budget-update", "budget-update-group", db)
	go budgetUpdateConsumer.ConsumeMessages(context.Background())
	defer budgetUpdateConsumer.Close()

	budgetDeleteConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "budget-delete", "budget-delete-group", db)
	go budgetDeleteConsumer.ConsumeMessages(context.Background())
	defer budgetDeleteConsumer.Close()

	// %--------- Category consumers

	categoryCreateConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "category-create", "category-create-group", db)
	go categoryCreateConsumer.ConsumeMessages(context.Background())
	defer categoryCreateConsumer.Close()

	categoryUpdateConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "category-update", "category-update-group", db)
	go categoryUpdateConsumer.ConsumeMessages(context.Background())
	defer categoryUpdateConsumer.Close()

	categoryDeleteConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "category-delete", "category-delete-group", db)
	go categoryDeleteConsumer.ConsumeMessages(context.Background())
	defer categoryDeleteConsumer.Close()

	// %--------- Transaction consumers
	transactionCreateConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "transaction-create", "transaction-create-group", db)
	go transactionCreateConsumer.ConsumeMessages(context.Background())
	defer transactionCreateConsumer.Close()

	transactionDeleteConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "transaction-delete", "transaction-delete-group", db)
	go transactionDeleteConsumer.ConsumeMessages(context.Background())
	defer transactionDeleteConsumer.Close()

	// %--------- Goal consumers
	goalCreateConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "goal-create", "goal-create-group", db)
	go goalCreateConsumer.ConsumeMessages(context.Background())
	defer goalCreateConsumer.Close()

	goalUpdateConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "goal-update", "goal-update-group", db)
	go goalUpdateConsumer.ConsumeMessages(context.Background())
	defer goalUpdateConsumer.Close()

	goalDeleteConsumer := NewKafkaConsumer(config.KAFKA_BROKER, "goal-delete", "goal-delete-group", db)
	go goalDeleteConsumer.ConsumeMessages(context.Background())
	defer goalDeleteConsumer.Close()
}
