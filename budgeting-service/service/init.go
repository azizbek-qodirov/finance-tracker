package service

import (
	"budget-service/storage"

	pb "budget-service/genprotos"

	"google.golang.org/grpc"
)

func InitServer(db *storage.Storage) *grpc.Server {
	accountService := NewAccountService(db)
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
