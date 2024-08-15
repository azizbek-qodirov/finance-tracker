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

type TransactionManager struct {
	Collection *mongo.Collection
}

func NewTransactionManager(client *mongo.Client, dbName, collectionName string) *TransactionManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &TransactionManager{Collection: collection}
}

func (m *TransactionManager) Create(req *pb.TransactionCReq) (*pb.Void, error) {
	newTransaction := bson.M{
		"user_id":          req.UserId,
		"account_id":       req.AccountId,
		"category_id":      req.CategoryId,
		"amount":           req.Amount,
		"type":             req.Type,
		"description":      req.Description,
		"created_datetime": time.Now().Format(time.RFC3339),
	}

	result, err := m.Collection.InsertOne(context.Background(), newTransaction)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	fmt.Println("Created transaction with ID:", result.InsertedID)
	return &pb.Void{}, nil
}

func (m *TransactionManager) GetByID(req *pb.ByID) (*pb.TransactionGRes, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid transaction ID: %v", err)
	}

	var transactionData struct {
		ID              primitive.ObjectID `bson:"_id"`
		UserID          string             `bson:"user_id"`
		AccountID       string             `bson:"account_id"`
		CategoryID      string             `bson:"category_id"`
		Amount          float32            `bson:"amount"`
		Type            string             `bson:"type"`
		Description     string             `bson:"description"`
		CreatedDatetime time.Time          `bson:"created_datetime"`
	}

	err = m.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&transactionData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, fmt.Errorf("failed to get transaction: %v", err)
	}

	transaction := &pb.TransactionGRes{
		Id:              transactionData.ID.Hex(),
		UserId:          transactionData.UserID,
		AccountId:       transactionData.AccountID,
		CategoryId:      transactionData.CategoryID,
		Amount:          transactionData.Amount,
		Type:            transactionData.Type,
		Description:     transactionData.Description,
		CreatedDatetime: transactionData.CreatedDatetime.Format(time.RFC3339),
	}

	return transaction, nil
}

func (m *TransactionManager) Delete(req *pb.ByID) (*pb.Void, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid transaction ID: %v", err)
	}

	_, err = m.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return nil, fmt.Errorf("failed to delete transaction: %v", err)
	}

	return &pb.Void{}, nil
}

func (m *TransactionManager) GetAll(req *pb.TransactionGAReq) (*pb.TransactionGARes, error) {
	filter := bson.M{}

	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.AccountId != "" {
		filter["account_id"] = req.AccountId
	}
	if req.CategoryId != "" {
		filter["category_id"] = req.CategoryId
	}
	if req.Amount != 0 {
		filter["amount"] = req.Amount
	}
	if req.Type != "" {
		filter["type"] = req.Type
	}

	if req.DateFrom != "" {
		dateFrom, err := time.Parse("2006-01-02", req.DateFrom)
		if err != nil {
			return nil, fmt.Errorf("invalid date from format: %v", err)
		}
		if filter["created_datetime"] == nil {
			filter["created_datetime"] = bson.M{}
		}
		filter["created_datetime"].(bson.M)["$gte"] = dateFrom
	}
	if req.DateTo != "" {
		dateTo, err := time.Parse("2006-01-02", req.DateTo)
		if err != nil {
			return nil, fmt.Errorf("invalid date to format: %v", err)
		}
		if filter["created_datetime"] == nil {
			filter["created_datetime"] = bson.M{}
		}
		filter["created_datetime"].(bson.M)["$lte"] = dateTo
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
		return nil, fmt.Errorf("failed to get transactions: %v", err)
	}
	defer cursor.Close(context.Background())

	var transactions []*pb.TransactionGRes
	for cursor.Next(context.Background()) {
		var transactionData struct {
			ID              primitive.ObjectID `bson:"_id"`
			UserID          string             `bson:"user_id"`
			AccountID       string             `bson:"account_id"`
			CategoryID      string             `bson:"category_id"`
			Amount          float32            `bson:"amount"`
			Type            string             `bson:"type"`
			Description     string             `bson:"description"`
			CreatedDatetime time.Time          `bson:"created_datetime"`
		}
		err := cursor.Decode(&transactionData)
		if err != nil {
			return nil, fmt.Errorf("failed to decode transaction: %v", err)
		}

		transaction := &pb.TransactionGRes{
			Id:              transactionData.ID.Hex(),
			UserId:          transactionData.UserID,
			AccountId:       transactionData.AccountID,
			CategoryId:      transactionData.CategoryID,
			Amount:          transactionData.Amount,
			Type:            transactionData.Type,
			Description:     transactionData.Description,
			CreatedDatetime: transactionData.CreatedDatetime.Format(time.RFC3339),
		}
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return &pb.TransactionGARes{Transactions: transactions}, nil
}
