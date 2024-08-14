package managers

import (
	"context"
	"fmt"
	"time"

	pb "budget-service/genprotos" // Update with your actual package path

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BudgetManager struct {
	Collection *mongo.Collection
}

func NewBudgetManager(client *mongo.Client, dbName, collectionName string) *BudgetManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &BudgetManager{Collection: collection}
}

func (m *BudgetManager) Create(req *pb.BudgetCReq) (*pb.Void, error) {
	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %v", err)
	}

	newBudget := bson.M{
		"user_id":     req.UserId,
		"category_id": req.CategoryId,
		"amount":      req.Amount,
		"period":      req.Period,
		"start_date":  startDate,
		"end_date":    endDate,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	result, err := m.Collection.InsertOne(context.Background(), newBudget)
	if err != nil {
		return nil, fmt.Errorf("failed to create budget: %v", err)
	}

	fmt.Println("Created budget with ID:", result.InsertedID)
	return &pb.Void{}, nil
}

func (m *BudgetManager) GetByID(req *pb.ByID) (*pb.BudgetGRes, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid budget ID: %v", err)
	}

	var budget pb.BudgetGRes
	err = m.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&budget)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("budget not found")
		}
		return nil, fmt.Errorf("failed to get budget: %v", err)
	}

	return &budget, nil
}

func (m *BudgetManager) Update(req *pb.BudgetUReq) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid budget ID: %v", err)
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %v", err)
	}

	update := bson.M{"$set": bson.M{
		"amount":     req.Amount,
		"period":     req.Period,
		"start_date": startDate,
		"end_date":   endDate,
		"updated_at": time.Now(),
	}}

	_, err = m.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update budget: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *BudgetManager) Delete(req *pb.ByID) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid budget ID: %v", err)
	}

	_, err = m.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return nil, fmt.Errorf("failed to delete budget: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *BudgetManager) GetAll(req *pb.BudgetGAreq) (*pb.BudgetGARes, error) {
	filter := bson.M{}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.CategoryId != "" {
		filter["category_id"] = req.CategoryId
	}
	if req.AmountFrom != 0 {
		filter["amount"] = bson.M{"$gte": req.AmountFrom}
	}
	if req.AmountTo != 0 {
		filter["amount"] = bson.M{"$lte": req.AmountTo}
	}
	if req.Period != "" {
		filter["period"] = req.Period
	}

	opts := options.Find()

	if req.Pagination.Limit != 0 {
		opts.SetLimit(req.Pagination.Limit)
	}
	if req.Pagination.Offset != 0 {
		opts.SetSkip(req.Pagination.Offset)
	}

	cursor, err := m.Collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get budgets: %v", err)
	}
	defer cursor.Close(context.Background())

	var budgets []*pb.BudgetGRes
	for cursor.Next(context.Background()) {
		var budget pb.BudgetGRes
		err := cursor.Decode(&budget)
		if err != nil {
			return nil, fmt.Errorf("failed to decode budget: %v", err)
		}
		budgets = append(budgets, &budget)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &pb.BudgetGARes{Budgets: budgets}, nil
}
