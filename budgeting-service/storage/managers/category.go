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

type CategoryManager struct {
	Collection *mongo.Collection
}

func NewCategoryManager(client *mongo.Client, dbName, collectionName string) *CategoryManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &CategoryManager{Collection: collection}
}

func (m *CategoryManager) Create(req *pb.CategoryCReq) (*pb.Void, error) {
	newCategory := bson.M{
		"user_id":    req.UserId,
		"name":       req.Name,
		"type":       req.Type,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	result, err := m.Collection.InsertOne(context.Background(), newCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %v", err)
	}

	fmt.Println("Created category with ID:", result.InsertedID)
	return &pb.Void{}, nil
}

func (m *CategoryManager) GetByID(req *pb.ByID) (*pb.CategoryGRes, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %v", err)
	}

	var category pb.CategoryGRes
	err = m.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("category not found")
		}
		return nil, fmt.Errorf("failed to get category: %v", err)
	}

	return &category, nil
}

func (m *CategoryManager) Update(req *pb.CategoryUReq) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %v", err)
	}

	update := bson.M{"$set": bson.M{
		"name":       req.Name,
		"updated_at": time.Now(),
	}}

	_, err = m.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *CategoryManager) Delete(req *pb.ByID) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid category ID: %v", err)
	}

	_, err = m.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return nil, fmt.Errorf("failed to delete category: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *CategoryManager) GetAll(req *pb.CategoryGAReq) (*pb.CategoryGARes, error) {
	filter := bson.M{}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
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
		return nil, fmt.Errorf("failed to get categories: %v", err)
	}
	defer cursor.Close(context.Background())

	var categories []*pb.CategoryGRes
	for cursor.Next(context.Background()) {
		var category pb.CategoryGRes
		err := cursor.Decode(&category)
		if err != nil {
			return nil, fmt.Errorf("failed to decode category: %v", err)
		}
		categories = append(categories, &category)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &pb.CategoryGARes{Categories: categories}, nil
}
