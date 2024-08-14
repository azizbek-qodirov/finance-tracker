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

type GoalManager struct {
	Collection *mongo.Collection
}

func NewGoalManager(client *mongo.Client, dbName, collectionName string) *GoalManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &GoalManager{Collection: collection}
}

func (m *GoalManager) Create(req *pb.GoalCReq) (*pb.Void, error) {
	deadline, err := time.Parse("2006-01-02", req.Deadline)
	if err != nil {
		return nil, fmt.Errorf("invalid deadline format: %v", err)
	}

	newGoal := bson.M{
		"user_id":        req.UserId,
		"name":           req.Name,
		"target_amount":  req.TargetAmount,
		"current_amount": req.CurrentAmount,
		"deadline":       deadline,
		"status":         req.Status,
		"created_at":     time.Now(),
		"updated_at":     time.Now(),
	}

	result, err := m.Collection.InsertOne(context.Background(), newGoal)
	if err != nil {
		return nil, fmt.Errorf("failed to create goal: %v", err)
	}

	fmt.Println("Created goal with ID:", result.InsertedID)
	return &pb.Void{}, nil
}

func (m *GoalManager) GetByID(req *pb.ByID) (*pb.GoalGRes, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid goal ID: %v", err)
	}

	var goal pb.GoalGRes
	err = m.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&goal)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("goal not found")
		}
		return nil, fmt.Errorf("failed to get goal: %v", err)
	}

	return &goal, nil
}

func (m *GoalManager) Update(req *pb.GoalUReq) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid goal ID: %v", err)
	}

	deadline, err := time.Parse("2006-01-02", req.Deadline) // Adjust date format if needed
	if err != nil {
		return nil, fmt.Errorf("invalid deadline format: %v", err)
	}

	update := bson.M{"$set": bson.M{
		"name":          req.Name,
		"target_amount": req.TargetAmount,
		"deadline":      deadline,
		"status":        req.Status,
		"updated_at":    time.Now(),
	}}

	_, err = m.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update goal: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *GoalManager) UpdateCurrentAmount(req *pb.GoalCurrentAmountUReq) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid goal ID: %v", err)
	}

	update := bson.M{"$set": bson.M{
		"current_amount": req.CurrentAmount,
		"updated_at":     time.Now(),
	}}

	_, err = m.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update goal current amount: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *GoalManager) Delete(req *pb.ByID) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid goal ID: %v", err)
	}

	_, err = m.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return nil, fmt.Errorf("failed to delete goal: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *GoalManager) GetAll(req *pb.GoalGAReq) (*pb.GoalGARes, error) {
	filter := bson.M{}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.Status != "" {
		filter["status"] = req.Status
	}
	if req.TargetFrom != 0 {
		filter["target_amount"] = bson.M{"$gte": req.TargetFrom}
	}
	if req.TargetTo != 0 {
		filter["target_amount"] = bson.M{"$lte": req.TargetTo}
	}

	if req.DeadlineFrom != "" {
		deadlineFrom, err := time.Parse("2006-01-02", req.DeadlineFrom)
		if err != nil {
			return nil, fmt.Errorf("invalid deadline from format: %v", err)
		}
		if filter["deadline"] == nil {
			filter["deadline"] = bson.M{}
		}
		filter["deadline"].(bson.M)["$gte"] = deadlineFrom
	}
	if req.DeadlineTo != "" {
		deadlineTo, err := time.Parse("2006-01-02", req.DeadlineTo)
		if err != nil {
			return nil, fmt.Errorf("invalid deadline to format: %v", err)
		}
		if filter["deadline"] == nil {
			filter["deadline"] = bson.M{}
		}
		filter["deadline"].(bson.M)["$lte"] = deadlineTo
	}

	opts := options.Find()
	if req.Pagination != nil {
		if req.Pagination.Limit > 0 {
			opts.SetLimit(req.Pagination.Limit)
		}
		if req.Pagination.Offset > 0 {
			opts.SetSkip(req.Pagination.Offset)
		}
	}

	cursor, err := m.Collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get goals: %v", err)
	}
	defer cursor.Close(context.Background())

	var goals []*pb.GoalGRes
	for cursor.Next(context.Background()) {
		var goal pb.GoalGRes
		err := cursor.Decode(&goal)
		if err != nil {
			return nil, fmt.Errorf("failed to decode goal: %v", err)
		}
		goals = append(goals, &goal)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &pb.GoalGARes{Goals: goals}, nil
}
