package service

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"context"
)

type TransactionService struct {
	storage storage.StorageI
	pb.UnimplementedTransactionServiceServer
}

func NewTransactionService(storage storage.StorageI) *TransactionService {
	return &TransactionService{storage: storage}
}

func (s *TransactionService) Create(ctx context.Context, req *pb.TransactionCReq) (*pb.Void, error) {
	return s.storage.Transaction().Create(req)
}

func (s *TransactionService) GetByID(ctx context.Context, req *pb.ByID) (*pb.TransactionGRes, error) {
	return s.storage.Transaction().GetByID(req)
}

func (s *TransactionService) GetAll(ctx context.Context, req *pb.TransactionGAReq) (*pb.TransactionGARes, error) {
	return s.storage.Transaction().GetAll(req)
}

func (s *TransactionService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.storage.Transaction().Delete(req)
}
