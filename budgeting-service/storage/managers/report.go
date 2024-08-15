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

type ReportManager struct {
	transactionCollection *mongo.Collection
	budgetCollection      *mongo.Collection
	goalCollection        *mongo.Collection
}

func NewReportManager(client *mongo.Client, dbName, transactionCollectionName, budgetCollectionName, goalCollectionName string) *ReportManager {
	transactionCollection := client.Database(dbName).Collection(transactionCollectionName)
	budgetCollection := client.Database(dbName).Collection(budgetCollectionName)
	goalCollection := client.Database(dbName).Collection(goalCollectionName)
	return &ReportManager{
		transactionCollection: transactionCollection,
		budgetCollection:      budgetCollection,
		goalCollection:        goalCollection,
	}
}

func (m *ReportManager) GetSpendings(req *pb.SpendingGReq) (*pb.SpendingGRes, error) {
	filter := m.buildTransactionFilter(req.UserId, req.CategoryId, "payment", req.DateFrom, req.DateTo) // Use "payment" type

	totalAmount, err := m.calculateTotalAmount(filter)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate total spending amount: %v", err)
	}

	transactions, err := m.getTransactions(filter, req.Pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to get spending transactions: %v", err)
	}

	return &pb.SpendingGRes{
		Request:      req,
		TotalAmount:  float32(totalAmount),
		Transactions: transactions,
	}, nil
}

func (m *ReportManager) GetIncomes(req *pb.IncomeGReq) (*pb.IncomeGRes, error) {
	filter := m.buildTransactionFilter(req.UserId, req.CategoryId, "income", req.DateFrom, req.DateTo) // Use "income" type

	// Calculate total income amount
	totalAmount, err := m.calculateTotalAmount(filter)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate total income amount: %v", err)
	}

	// Get transactions (with pagination if needed)
	transactions, err := m.getTransactions(filter, req.Pagination)
	if err != nil {
		return nil, fmt.Errorf("failed to get income transactions: %v", err)
	}

	return &pb.IncomeGRes{
		Request:      req,
		TotalAmount:  float32(totalAmount),
		Transactions: transactions,
	}, nil
}

func (m *ReportManager) BudgetPerformance(req *pb.BudgetPerReq) (*pb.BudgetPerGet, error) {
	budgetFilter := bson.M{"user_id": req.UserId}
	if req.CategoryId != "" {
		budgetFilter["category_id"] = req.CategoryId
	}
	if req.Period != "" {
		budgetFilter["period"] = req.Period
	}

	budgetCursor, err := m.budgetCollection.Find(context.Background(), budgetFilter)
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

		transactionFilter := m.buildTransactionFilter(
			req.UserId,
			budget.CategoryID,
			"payment",
			budget.StartDate.Format("2006-01-02"),
			budget.EndDate.Format("2006-01-02"),
		)

		totalSpendings, err := m.calculateTotalAmount(transactionFilter)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate total spending amount: %v", err)
		}

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
	filter := bson.M{"user_id": req.UserId}
	if req.Status != "" {
		filter["status"] = req.Status
	}

	// Date Range Filtering (adjust date format if needed)
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
			GoalName      string    `bson:"name"` // Assuming "goal_name" maps to "name" in MongoDB
		}
		err := cursor.Decode(&goal)
		if err != nil {
			return nil, fmt.Errorf("failed to decode goal: %v", err)
		}

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

func (m *ReportManager) buildTransactionFilter(userID, categoryID, transactionType, dateFrom, dateTo string) bson.M {
	filter := bson.M{"user_id": userID}
	if categoryID != "" {
		filter["category_id"] = categoryID
	}
	if transactionType != "" {
		filter["type"] = transactionType
	}

	if dateFrom != "" {
		dateFromTime, _ := time.Parse("2006-01-02", dateFrom)
		if filter["created_datetime"] == nil {
			filter["created_datetime"] = bson.M{}
		}
		filter["created_datetime"].(bson.M)["$gte"] = dateFromTime
	}
	if dateTo != "" {
		dateToTime, _ := time.Parse("2006-01-02", dateTo)
		if filter["created_datetime"] == nil {
			filter["created_datetime"] = bson.M{}
		}
		filter["created_datetime"].(bson.M)["$lte"] = dateToTime
	}

	return filter
}

func (m *ReportManager) calculateTotalAmount(filter bson.M) (float64, error) {
	pipeline := []bson.M{
		{"$match": filter},
		{"$group": bson.M{
			"_id":         nil,
			"totalAmount": bson.M{"$sum": "$amount"},
		}},
	}

	cursor, err := m.transactionCollection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return 0, fmt.Errorf("failed to aggregate transactions: %v", err)
	}
	defer cursor.Close(context.Background())

	var result []bson.M
	if err := cursor.All(context.Background(), &result); err != nil {
		return 0, fmt.Errorf("failed to decode aggregation result: %v", err)
	}

	if len(result) > 0 {
		if totalAmount, ok := result[0]["totalAmount"].(float64); ok {
			return totalAmount, nil
		}
	}

	return 0, nil
}

func (m *ReportManager) getTransactions(filter bson.M, pagination *pb.Pagination) ([]*pb.TransactionGARes, error) {
	opts := options.Find()
	if pagination != nil {
		if pagination.Limit > 0 {
			opts.SetLimit(pagination.Limit)
		}
		if pagination.Offset > 0 {
			opts.SetSkip(pagination.Offset)
		}
	}

	cursor, err := m.transactionCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %v", err)
	}
	defer cursor.Close(context.Background())

	var transactions []*pb.TransactionGARes
	for cursor.Next(context.Background()) {
		var transaction pb.TransactionGARes
		err := cursor.Decode(&transaction)
		if err != nil {
			return nil, fmt.Errorf("failed to decode transaction: %v", err)
		}
		transactions = append(transactions, &transaction)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("transaction cursor error: %v", err)
	}

	return transactions, nil
}
