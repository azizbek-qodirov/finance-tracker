package storage

import (
	"context"
	"fmt"

	"budget-service/config"
	"budget-service/storage/managers"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cfg = config.Load()
)

type Storage struct {
	MongoClient *mongo.Client

	AccountS     AccountI
	BudgetS      BudgetI
	CategoryS    CategoryI
	GoalS        GoalI
	ReportS      ReportI
	TransactionS TransactionI
}

func NewDatabaseStorage(config config.Config) (*Storage, *redis.Client, error) {
	// #################     MONGODB CONNECTION     ###################### //
	clientOptions := options.Client().ApplyURI(config.MONGO_URI).SetAuth(options.Credential{
		Username: "root",
		Password: "root",
	})
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, nil, err
	}
	fmt.Println("Successfully connected to the database mongodb!!!")

	// #################     REDIS CONNECTION     ###################### //
	rclient := redis.NewClient(&redis.Options{
		Addr: cfg.REDIS_HOST + ":" + cfg.REDIS_PORT, // "localhost:6379"
		DB:   0,
	})
	fmt.Println("Successfully connected to the database redis!!!")

	return &Storage{
		MongoClient: client,
	}, rclient, nil
}

func (s *Storage) Account() AccountI {
	if s.AccountS == nil {
		s.AccountS = managers.NewAccountManager(s.MongoClient, cfg.MONGO_DB_NAME, cfg.MONGO_ACCOUNTS_COLLECTION_NAME)
	}
	return s.AccountS
}

func (s *Storage) Budget() BudgetI {
	if s.BudgetS == nil {
		s.BudgetS = managers.NewBudgetManager(s.MongoClient, cfg.MONGO_DB_NAME, cfg.MONGO_BUDGETS_COLLECTION_NAME)
	}
	return s.BudgetS
}

func (s *Storage) Category() CategoryI {
	if s.CategoryS == nil {
		s.CategoryS = managers.NewCategoryManager(s.MongoClient, cfg.MONGO_DB_NAME, cfg.MONGO_CATEGORIES_COLLECTION_NAME)
	}
	return s.CategoryS
}

func (s *Storage) Goal() GoalI {
	if s.GoalS == nil {
		s.GoalS = managers.NewGoalManager(s.MongoClient, cfg.MONGO_DB_NAME, cfg.MONGO_GOALS_COLLECTION_NAME)
	}
	return s.GoalS
}

func (s *Storage) Report() ReportI {
	if s.ReportS == nil {
		s.ReportS = managers.NewReportManager(s.MongoClient, cfg.MONGO_DB_NAME, cfg.MONGO_TRANSACTIONS_COLLECTION_NAME, cfg.MONGO_BUDGETS_COLLECTION_NAME, cfg.MONGO_GOALS_COLLECTION_NAME)
	}
	return s.ReportS
}

func (s *Storage) Transaction() TransactionI {
	if s.TransactionS == nil {
		s.TransactionS = managers.NewTransactionManager(s.MongoClient, cfg.MONGO_DB_NAME, cfg.MONGO_TRANSACTIONS_COLLECTION_NAME)
	}
	return s.TransactionS
}

func (s *Storage) Close() {
	s.MongoClient.Disconnect(context.Background())
}
