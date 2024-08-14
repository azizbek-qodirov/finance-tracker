package main

import (
	"log"
	"net"

	cf "budget-service/config"
	"budget-service/storage"

	service "budget-service/service"

	pb "budget-service/genprotos"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()

	db, err := storage.NewPostgresStorage(config)
	em.CheckErr(err)
	defer db.Close()

	listener, err := net.Listen("tcp", config.BUDGET_SERVICE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	accountService := service.NewAccountService(db)
	budgetService := service.NewBudgetService(db)
	categoryService := service.NewCategoryService(db)
	goalService := service.NewGoalService(db)
	reportService := service.NewReportService(db)
	transactionService := service.NewTransactionService(db)

	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, accountService)
	pb.RegisterBudgetServiceServer(s, budgetService)
	pb.RegisterCategoryServiceServer(s, categoryService)
	pb.RegisterGoalServiceServer(s, goalService)
	pb.RegisterReportServiceServer(s, reportService)
	pb.RegisterTransactionServiceServer(s, transactionService)

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
