package service

import (
	pb "budget-service/genprotos"
	"budget-service/storage"
	"context"
)

type CategoryService struct {
	storage storage.StorageI
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryService(storage storage.StorageI) *CategoryService {
	return &CategoryService{storage: storage}
}

func (s *CategoryService) Create(ctx context.Context, req *pb.CategoryCReq) (*pb.Void, error) {
	return s.storage.Category().Create(req)
}

func (s *CategoryService) GetByID(ctx context.Context, req *pb.ByID) (*pb.CategoryGRes, error) {
	return s.storage.Category().GetByID(req)
}

func (s *CategoryService) Update(ctx context.Context, req *pb.CategoryUpdate) (*pb.Void, error) {
	return s.storage.Category().Update(req)
}

func (s *CategoryService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.storage.Category().Delete(req)
}

func (s *CategoryService) GetAll(ctx context.Context, req *pb.CategoryGAReq) (*pb.CategoryGARes, error) {
	return s.storage.Category().GetAll(req)
}
