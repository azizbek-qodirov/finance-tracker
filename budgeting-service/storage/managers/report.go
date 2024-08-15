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

type ReportManager struct {
	transactionCollection *mongo.Collection
	transactionManager    *TransactionManager
	budgetCollection      *mongo.Collection
	goalCollection        *mongo.Collection
}

func NewReportManager(client *mongo.Client, dbName, transactionCollectionName, budgetCollectionName, goalCollectionName string) *ReportManager {
	transactionCollection := client.Database(dbName).Collection(transactionCollectionName)
	budgetCollection := client.Database(dbName).Collection(budgetCollectionName)
	goalCollection := client.Database(dbName).Collection(goalCollectionName)
	return &ReportManager{
		transactionManager:    NewTransactionManager(client, dbName, transactionCollectionName),
		budgetCollection:      budgetCollection,
		goalCollection:        goalCollection,
		transactionCollection: transactionCollection,
	}
}

func (m *ReportManager) GetSpendings(req *pb.SpendingGReq) (*pb.SpendingGRes, error) {
	// Build the filter for transactions
	transactionReq := &pb.TransactionGAReq{
		UserId:     req.UserId,
		CategoryId: req.CategoryId,
		Type:       "payment", // Filter by "payment" type
		DateFrom:   req.DateFrom,
		DateTo:     req.DateTo,
		Pagination: req.Pagination,
	}

	// Get transactions using TransactionManager
	transactionRes, err := m.transactionManager.GetAll(transactionReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get spending transactions: %v", err)
	}

	// Calculate total spending amount
	totalAmount := float64(0)
	for _, transaction := range transactionRes.Transactions {
		totalAmount += float64(transaction.Amount)
	}

	return &pb.SpendingGRes{
		Request:      req,
		TotalAmount:  float32(totalAmount),
		Transactions: transactionRes.Transactions,
	}, nil
}

func (m *ReportManager) GetIncomes(req *pb.IncomeGReq) (*pb.IncomeGRes, error) {
	// Build the filter for transactions
	transactionReq := &pb.TransactionGAReq{
		UserId:     req.UserId,
		CategoryId: req.CategoryId,
		Type:       "income", // Filter by "income" type
		DateFrom:   req.DateFrom,
		DateTo:     req.DateTo,
		Pagination: req.Pagination,
	}

	// Get transactions using TransactionManager
	transactionRes, err := m.transactionManager.GetAll(transactionReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get income transactions: %v", err)
	}

	// Calculate total income amount
	totalAmount := float64(0)
	for _, transaction := range transactionRes.Transactions {
		totalAmount += float64(transaction.Amount)
	}

	return &pb.IncomeGRes{
		Request:      req,
		TotalAmount:  float32(totalAmount),
		Transactions: transactionRes.Transactions,
	}, nil
}
func (m *ReportManager) BudgetPerformance(req *pb.BudgetPerReq) (*pb.BudgetPerGet, error) {
	// Build the filter for budgets
	filter := bson.M{"user_id": req.UserId}
	if req.CategoryId != "" {
		filter["category_id"] = req.CategoryId
	}
	if req.Period != "" {
		filter["period"] = req.Period
	}

	// Find matching budgets
	budgetCursor, err := m.budgetCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get budgets: %v", err)
	}
	defer budgetCursor.Close(context.Background())

	var performances []*pb.PeriodBudgetPer
	for budgetCursor.Next(context.Background()) {
		var budget struct {
			ID         primitive.ObjectID `bson:"_id"`
			UserID     string             `bson:"user_id"`
			CategoryID string             `bson:"category_id"`
			Amount     float64            `bson:"amount"`
			Period     string             `bson:"period"`
			StartDate  time.Time          `bson:"start_date"`
			EndDate    time.Time          `bson:"end_date"`
		}
		err := budgetCursor.Decode(&budget)
		if err != nil {
			return nil, fmt.Errorf("failed to decode budget: %v", err)
		}

		// Build the filter for transactions within the budget period
		transactionFilter := bson.M{
			"user_id":          req.UserId,
			"category_id":      budget.CategoryID,
			"type":             "payment",
			"created_datetime": bson.M{"$gte": budget.StartDate, "$lte": budget.EndDate},
		}

		// Calculate total spendings for the budget period using aggregation
		pipeline := []bson.M{
			{"$match": transactionFilter},
			{"$group": bson.M{
				"_id":            nil,
				"totalSpendings": bson.M{"$sum": "$amount"},
			}},
		}
		cursor, err := m.transactionCollection.Aggregate(context.Background(), pipeline)
		if err != nil {
			return nil, fmt.Errorf("failed to aggregate transactions: %v", err)
		}
		defer cursor.Close(context.Background())
		var result []bson.M
		if err := cursor.All(context.Background(), &result); err != nil {
			return nil, fmt.Errorf("failed to decode aggregation result: %v", err)
		}
		var totalSpendings float64
		if len(result) > 0 {
			if val, ok := result[0]["totalSpendings"].(float64); ok {
				totalSpendings = val
			}
		}

		// Calculate progress
		progress := float64(0)
		if budget.Amount > 0 {
			progress = totalSpendings / budget.Amount
		}

		performances = append(performances, &pb.PeriodBudgetPer{
			StartDate:      budget.StartDate.Format("2006-01-02"),
			EndDate:        budget.EndDate.Format("2006-01-02"),
			TotalSpendings: float32(totalSpendings),
			TargetAmount:   float32(budget.Amount),
			Progress:       float32(progress),
			Period:         budget.Period,
			CategoryId:     budget.CategoryID,
		})
	}

	if err := budgetCursor.Err(); err != nil {
		return nil, fmt.Errorf("budget cursor error: %v", err)
	}

	return &pb.BudgetPerGet{Performances: performances}, nil
}

func (m *ReportManager) GoalProgress(req *pb.GoalProgresReq) (*pb.GoalProgresGet, error) {
	// Build the filter for goals
	filter := bson.M{"user_id": req.UserId}
	if req.Status != "" {
		filter["status"] = req.Status
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

	// Find matching goals
	cursor, err := m.goalCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get goals: %v", err)
	}
	defer cursor.Close(context.Background())

	var goalProgresses []*pb.GoalProgress
	for cursor.Next(context.Background()) {
		var goal struct {
			Deadline      time.Time `bson:"deadline"`
			TargetAmount  float64   `bson:"target_amount"`
			CurrentAmount float64   `bson:"current_amount"`
			GoalName      string    `bson:"name"`
		}
		err := cursor.Decode(&goal)
		if err != nil {
			return nil, fmt.Errorf("failed to decode goal: %v", err)
		}

		// Calculate progress
		progress := float64(0)
		if goal.TargetAmount > 0 {
			progress = goal.CurrentAmount / goal.TargetAmount
		}

		goalProgresses = append(goalProgresses, &pb.GoalProgress{
			Deadline:      goal.Deadline.Format("2006-01-02"),
			TargetAmount:  float32(goal.TargetAmount),
			CurrentAmount: float32(goal.CurrentAmount),
			Progress:      float32(progress),
			GoalName:      goal.GoalName,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("goal cursor error: %v", err)
	}

	return &pb.GoalProgresGet{Goals: goalProgresses}, nil
}
