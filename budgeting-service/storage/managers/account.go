package managers

import (
	"context"
	"fmt"
	"time"

	pb "budget-service/genprotos" // Update with your actual package path

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountManager struct {
	Collection *mongo.Collection
}

func NewAccountManager(client *mongo.Client, dbName, collectionName string) *AccountManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &AccountManager{Collection: collection}
}

func (m *AccountManager) GetAccount(req *pb.ByUserID) (*pb.AccountGRes, error) {
	filter := bson.M{"user_id": req.UserId}

	var account struct {
		ID        primitive.ObjectID `bson:"_id"`
		UserID    string             `bson:"user_id"`
		Name      string             `bson:"name"`
		Type      string             `bson:"type"`
		Currency  string             `bson:"currency"`
		CreatedAt time.Time          `bson:"created_at"`
		UpdatedAt time.Time          `bson:"updated_at"`
	}

	err := m.Collection.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("account not found")
		}
		return nil, fmt.Errorf("failed to get account: %v", err)
	}

	return &pb.AccountGRes{
		Id:     account.ID.Hex(),
		UserId: account.UserID,
		Name:   account.Name,
		Type:   account.Type,
	}, nil
}

func (m *AccountManager) GetBalance(req *pb.ByUserID) (*pb.AccountBalanceGRes, error) {
	filter := bson.M{"user_id": req.UserId}

	var result struct {
		Balance  float32 `bson:"balance"`
		Currency string  `bson:"currency"`
	}

	err := m.Collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("account not found")
		}
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}

	return &pb.AccountBalanceGRes{
		Balance:  result.Balance,
		Currency: result.Currency,
	}, nil
}

func (m *AccountManager) UpdateAccount(req *pb.AccountUReq) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid account ID: %v", err)
	}

	update := bson.M{"$set": bson.M{
		"name":       req.Name,
		"updated_at": time.Now(),
	}}

	_, err = m.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update account: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *AccountManager) UpdateBalance(req *pb.AccountBalanceUReq) (*pb.Void, error) {
	filter := bson.M{"user_id": req.UserId}
	update := bson.M{"$set": bson.M{
		"balance":    req.Balance,
		"updated_at": time.Now(),
	}}

	_, err := m.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update balance: %v", err)
	}

	return &pb.Void{}, nil
}
