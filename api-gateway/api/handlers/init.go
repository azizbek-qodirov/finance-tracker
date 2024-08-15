package handlers

import (
	pb "gateway-service/genprotos" // Update with your actual package path

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Account     pb.AccountServiceClient
	Budget      pb.BudgetServiceClient
	Category    pb.CategoryServiceClient
	Goal        pb.GoalServiceClient
	Report      pb.ReportServiceClient
	Transaction pb.TransactionServiceClient
}

func NewHandler(conn *grpc.ClientConn) *HTTPHandler {
	return &HTTPHandler{
		Account:     pb.NewAccountServiceClient(conn),
		Budget:      pb.NewBudgetServiceClient(conn),
		Category:    pb.NewCategoryServiceClient(conn),
		Goal:        pb.NewGoalServiceClient(conn),
		Report:      pb.NewReportServiceClient(conn),
		Transaction: pb.NewTransactionServiceClient(conn),
	}
}
