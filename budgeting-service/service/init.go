package service

import (
	"budget-service/storage"

	pb "budget-service/genprotos"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

func InitServer(db *storage.Storage, redis *redis.Client) *grpc.Server {
	accountService := NewAccountService(db, redis)
	budgetService := NewBudgetService(db)
	categoryService := NewCategoryService(db)
	goalService := NewGoalService(db)
	reportService := NewReportService(db)
	transactionService := NewTransactionService(db)

	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, accountService)
	pb.RegisterBudgetServiceServer(s, budgetService)
	pb.RegisterCategoryServiceServer(s, categoryService)
	pb.RegisterGoalServiceServer(s, goalService)
	pb.RegisterReportServiceServer(s, reportService)
	pb.RegisterTransactionServiceServer(s, transactionService)

	return s
}
