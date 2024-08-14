package service

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"context"
)

type BudgetService struct {
	storage storage.StorageI
	pb.UnimplementedBudgetServiceServer
}

func NewBudgetService(storage storage.StorageI) *BudgetService {
	return &BudgetService{storage: storage}
}

func (s *BudgetService) Create(ctx context.Context, req *pb.BudgetCReq) (*pb.Void, error) {
	return s.storage.Budget().Create(req)
}

func (s *BudgetService) GetByID(ctx context.Context, req *pb.ByID) (*pb.BudgetGRes, error) {
	return s.storage.Budget().GetByID(req)
}

func (s *BudgetService) Update(ctx context.Context, req *pb.BudgetUReq) (*pb.Void, error) {
	return s.storage.Budget().Update(req)
}

func (s *BudgetService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.storage.Budget().Delete(req)
}

func (s *BudgetService) GetAll(ctx context.Context, req *pb.BudgetGAreq) (*pb.BudgetGARes, error) {
	return s.storage.Budget().GetAll(req)
}
