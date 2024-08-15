package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_PORT           string
	BUDGET_SERVICE_PORT string

	MONGO_URI                          string
	MONGO_DB_NAME                      string
	MONGO_TRANSACTIONS_COLLECTION_NAME string
	MONGO_ACCOUNTS_COLLECTION_NAME     string
	MONGO_CATEGORIES_COLLECTION_NAME   string
	MONGO_GOALS_COLLECTION_NAME        string
	MONGO_REPORTS_COLLECTION_NAME      string
	MONGO_BUDGETS_COLLECTION_NAME      string

	KAFKA_BROKER string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.AUTH_PORT = cast.ToString(coalesce("AUTH_PORT", ":8088"))
	config.BUDGET_SERVICE_PORT = cast.ToString(coalesce("BUDGET_SERVICE_PORT", ":50052"))

	config.MONGO_URI = cast.ToString(coalesce("MONGO_URI", "mongodb://localhost:27017"))
	config.MONGO_DB_NAME = cast.ToString(coalesce("MONGO_DB_NAME", "lingua_learning"))
	config.MONGO_TRANSACTIONS_COLLECTION_NAME = cast.ToString(coalesce("MONGO_TRANSACTIONS_COLLECTION_NAME", "transactions"))
	config.MONGO_ACCOUNTS_COLLECTION_NAME = cast.ToString(coalesce("MONGO_ACCOUNTS_COLLECTION_NAME", "accounts"))
	config.MONGO_CATEGORIES_COLLECTION_NAME = cast.ToString(coalesce("MONGO_CATEGORIES_COLLECTION_NAME", "categories"))
	config.MONGO_GOALS_COLLECTION_NAME = cast.ToString(coalesce("MONGO_GOALS_COLLECTION_NAME", "goals"))
	config.MONGO_REPORTS_COLLECTION_NAME = cast.ToString(coalesce("MONGO_REPORTS_COLLECTION_NAME", "reports"))
	config.MONGO_BUDGETS_COLLECTION_NAME = cast.ToString(coalesce("MONGO_BUDGETS_COLLECTION_NAME", "budgets"))

	config.KAFKA_BROKER = cast.ToString(coalesce("KAFKA_BROKER", "localhost:9092"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
