package service

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"context"
)

type AccountService struct {
	storage storage.StorageI
	pb.UnimplementedAccountServiceServer
}

func NewAccountService(storage storage.StorageI) *AccountService {
	return &AccountService{storage: storage}
}

func (s *AccountService) GetAccount(ctx context.Context, req *pb.ByUserID) (*pb.AccountGRes, error) {
	return s.storage.Account().GetAccount(req)
}

func (s *AccountService) GetBalance(ctx context.Context, req *pb.ByUserID) (*pb.AccountBalanceGRes, error) {
	return s.storage.Account().GetBalance(req)
}

func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.AccountUReq) (*pb.Void, error) {
	return s.storage.Account().UpdateAccount(req)
}

func (s *AccountService) UpdateBalance(ctx context.Context, req *pb.AccountBalanceUReq) (*pb.Void, error) {
	return s.storage.Account().UpdateBalance(req)
}
