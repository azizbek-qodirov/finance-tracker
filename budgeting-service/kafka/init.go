package kfk

import (
	cf "budget-service/config"
	"budget-service/storage"
)

func InitKafka(config *cf.Config, db *storage.Storage) {
	manager := NewKafkaConsumerManager()

	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "account-updates", "account-update-group", AccountUpdateHandler(db))
	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "account-balance-updates", "account-balance-update-group", AccountBalanceUpdateHandler(db))

	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "budget-create", "budget-create-group", BudgetCreateHandler(db))
	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "budget-update", "budget-update-group", BudgetUpdateHandler(db))
	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "budget-delete", "budget-delete-group", BudgetDeleteHandler(db))

	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "category-create", "category-create-group", CategoryCreateHandler(db))
	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "category-update", "category-update-group", CategoryUpdateHandler(db))
	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "category-delete", "category-delete-group", CategoryDeleteHandler(db))

	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "transaction-create", "transaction-create-group", TransactionCreateHandler(db))
	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "transaction-delete", "transaction-delete-group", TransactionDeleteHandler(db))

	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "goal-create", "goal-create-group", GoalCreateHandler(db))
	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "goal-update", "goal-update-group", GoalUpdateHandler(db))
	manager.RegisterConsumer([]string{config.KAFKA_BROKER}, "goal-delete", "goal-delete-group", GoalDeleteHandler(db))
}
