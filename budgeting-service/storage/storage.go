package storage

import (
	pb "budget-service/genprotos"
)

type StorageI interface {
	Account() AccountI
	Budget() BudgetI
	Category() CategoryI
	Goal() GoalI
	Report() ReportI
	Transaction() TransactionI
}

type AccountI interface {
	GetAccount(*pb.ByUserID) (*pb.AccountGRes, error)
	GetBalance(*pb.ByUserID) (*pb.AccountBalanceGRes, error)
	UpdateAccount(*pb.AccountUReq) (*pb.Void, error)
	UpdateBalance(*pb.AccountBalanceUReq) (*pb.Void, error)
}

type BudgetI interface {
	Create(*pb.BudgetCReq) (*pb.Void, error)
	GetByID(*pb.ByID) (*pb.BudgetGRes, error)
	Update(*pb.BudgetUReq) (*pb.Void, error)
	Delete(*pb.ByID) (*pb.Void, error)
	GetAll(*pb.BudgetGAreq) (*pb.BudgetGARes, error)
}

type CategoryI interface {
	Create(*pb.CategoryCReq) (*pb.Void, error)
	GetByID(*pb.ByID) (*pb.CategoryGRes, error)
	Update(*pb.CategoryUReq) (*pb.Void, error)
	Delete(*pb.ByID) (*pb.Void, error)
	GetAll(*pb.CategoryGAReq) (*pb.CategoryGARes, error)
}

type GoalI interface {
	Create(*pb.GoalCReq) (*pb.Void, error)
	GetByID(*pb.ByID) (*pb.GoalGRes, error)
	Update(*pb.GoalUReq) (*pb.Void, error)
	UpdateCurrentAmount(*pb.GoalCurrentAmountUReq) (*pb.Void, error)
	Delete(*pb.ByID) (*pb.Void, error)
	GetAll(*pb.GoalGAReq) (*pb.GoalGARes, error)
}

type ReportI interface {
	GetSpendings(*pb.SpendingGReq) (*pb.SpendingGRes, error)
	GetIncomes(*pb.IncomeGReq) (*pb.IncomeGRes, error)
	BudgetPerformance(*pb.BudgetPerReq) (*pb.BudgetPerGet, error)
	GoalProgress(*pb.GoalProgresReq) (*pb.GoalProgresGet, error)
}

type TransactionI interface {
	Create(*pb.TransactionCReq) (*pb.Void, error)
	GetByID(*pb.ByID) (*pb.TransactionGRes, error)
	Delete(*pb.ByID) (*pb.Void, error)
	GetAll(*pb.TransactionGAReq) (*pb.TransactionGARes, error)
}
