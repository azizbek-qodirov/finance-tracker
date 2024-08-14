package service

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"context"
)

type GoalService struct {
	storage storage.StorageI
	pb.UnimplementedGoalServiceServer
}

func NewGoalService(storage storage.StorageI) *GoalService {
	return &GoalService{storage: storage}
}

func (s *GoalService) Create(ctx context.Context, req *pb.GoalCReq) (*pb.Void, error) {
	return s.storage.Goal().Create(req)
}

func (s *GoalService) GetByID(ctx context.Context, req *pb.ByID) (*pb.GoalGRes, error) {
	return s.storage.Goal().GetByID(req)
}

func (s *GoalService) Update(ctx context.Context, req *pb.GoalUReq) (*pb.Void, error) {
	return s.storage.Goal().Update(req)
}

func (s *GoalService) UpdateCurrentAmount(ctx context.Context, req *pb.GoalCurrentAmountUReq) (*pb.Void, error) {
	return s.storage.Goal().UpdateCurrentAmount(req)
}

func (s *GoalService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.storage.Goal().Delete(req)
}

func (s *GoalService) GetAll(ctx context.Context, req *pb.GoalGAReq) (*pb.GoalGARes, error) {
	return s.storage.Goal().GetAll(req)
}
