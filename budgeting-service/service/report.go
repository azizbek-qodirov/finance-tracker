package service

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"context"
)

type ReportService struct {
	storage storage.StorageI
	pb.UnimplementedReportServiceServer
}

func NewReportService(storage storage.StorageI) *ReportService {
	return &ReportService{storage: storage}
}

func (s *ReportService) GetSpendings(ctx context.Context, req *pb.SpendingGReq) (*pb.SpendingGRes, error) {
	return s.storage.Report().GetSpendings(req)
}

func (s *ReportService) GetIncomes(ctx context.Context, req *pb.IncomeGReq) (*pb.IncomeGRes, error) {
	return s.storage.Report().GetIncomes(req)
}

// Uncomment and implement when ready
// func (s *ReportService) BudgetPerformance(ctx context.Context, req *pb.BudgetPerReq) (*pb.BudgetPerGet, error) {
// 	return s.storage.Report().BudgetPerformance(req)
// }

// func (s *ReportService) GoalProgress(ctx context.Context, req *pb.GoalProgresReq) (*pb.GoalProgresGet, error) {
// 	return s.storage.Report().GoalProgress(req)
// }
