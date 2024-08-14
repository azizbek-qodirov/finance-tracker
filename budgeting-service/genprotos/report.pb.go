// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: finance-protos/report.proto

package genprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SpendingGReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DateFrom   string      `protobuf:"bytes,2,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo     string      `protobuf:"bytes,3,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
	CategoryId string      `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Pagination *Pagination `protobuf:"bytes,5,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *SpendingGReq) Reset() {
	*x = SpendingGReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpendingGReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpendingGReq) ProtoMessage() {}

func (x *SpendingGReq) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpendingGReq.ProtoReflect.Descriptor instead.
func (*SpendingGReq) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{0}
}

func (x *SpendingGReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SpendingGReq) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *SpendingGReq) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *SpendingGReq) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *SpendingGReq) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type SpendingGRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request      *SpendingGReq       `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	TotalAmount  float32             `protobuf:"fixed32,2,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Transactions []*TransactionGARes `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *SpendingGRes) Reset() {
	*x = SpendingGRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpendingGRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpendingGRes) ProtoMessage() {}

func (x *SpendingGRes) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpendingGRes.ProtoReflect.Descriptor instead.
func (*SpendingGRes) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{1}
}

func (x *SpendingGRes) GetRequest() *SpendingGReq {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *SpendingGRes) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *SpendingGRes) GetTransactions() []*TransactionGARes {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type IncomeGReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DateFrom   string      `protobuf:"bytes,2,opt,name=date_from,json=dateFrom,proto3" json:"date_from,omitempty"`
	DateTo     string      `protobuf:"bytes,3,opt,name=date_to,json=dateTo,proto3" json:"date_to,omitempty"`
	CategoryId string      `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Pagination *Pagination `protobuf:"bytes,5,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *IncomeGReq) Reset() {
	*x = IncomeGReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeGReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeGReq) ProtoMessage() {}

func (x *IncomeGReq) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeGReq.ProtoReflect.Descriptor instead.
func (*IncomeGReq) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{2}
}

func (x *IncomeGReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IncomeGReq) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *IncomeGReq) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *IncomeGReq) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *IncomeGReq) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type IncomeGRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request      *IncomeGReq         `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	TotalAmount  float32             `protobuf:"fixed32,2,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Transactions []*TransactionGARes `protobuf:"bytes,4,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *IncomeGRes) Reset() {
	*x = IncomeGRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeGRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeGRes) ProtoMessage() {}

func (x *IncomeGRes) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeGRes.ProtoReflect.Descriptor instead.
func (*IncomeGRes) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{3}
}

func (x *IncomeGRes) GetRequest() *IncomeGReq {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *IncomeGRes) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *IncomeGRes) GetTransactions() []*TransactionGARes {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type BudgetPerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CategoryId string `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Period     string `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
	StartDate  string `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate    string `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *BudgetPerReq) Reset() {
	*x = BudgetPerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetPerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetPerReq) ProtoMessage() {}

func (x *BudgetPerReq) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetPerReq.ProtoReflect.Descriptor instead.
func (*BudgetPerReq) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{4}
}

func (x *BudgetPerReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BudgetPerReq) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *BudgetPerReq) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *BudgetPerReq) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *BudgetPerReq) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type PeriodBudgetPer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate      string  `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate        string  `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	TotalSpendings float32 `protobuf:"fixed32,3,opt,name=total_spendings,json=totalSpendings,proto3" json:"total_spendings,omitempty"`
	TargetAmount   float32 `protobuf:"fixed32,4,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	Progress       float32 `protobuf:"fixed32,5,opt,name=progress,proto3" json:"progress,omitempty"`
	Period         string  `protobuf:"bytes,6,opt,name=period,proto3" json:"period,omitempty"`
	CategoryId     string  `protobuf:"bytes,7,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *PeriodBudgetPer) Reset() {
	*x = PeriodBudgetPer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeriodBudgetPer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeriodBudgetPer) ProtoMessage() {}

func (x *PeriodBudgetPer) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeriodBudgetPer.ProtoReflect.Descriptor instead.
func (*PeriodBudgetPer) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{5}
}

func (x *PeriodBudgetPer) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *PeriodBudgetPer) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *PeriodBudgetPer) GetTotalSpendings() float32 {
	if x != nil {
		return x.TotalSpendings
	}
	return 0
}

func (x *PeriodBudgetPer) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *PeriodBudgetPer) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *PeriodBudgetPer) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *PeriodBudgetPer) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type BudgetPerGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Performances []*PeriodBudgetPer `protobuf:"bytes,1,rep,name=performances,proto3" json:"performances,omitempty"`
}

func (x *BudgetPerGet) Reset() {
	*x = BudgetPerGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetPerGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetPerGet) ProtoMessage() {}

func (x *BudgetPerGet) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetPerGet.ProtoReflect.Descriptor instead.
func (*BudgetPerGet) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{6}
}

func (x *BudgetPerGet) GetPerformances() []*PeriodBudgetPer {
	if x != nil {
		return x.Performances
	}
	return nil
}

type GoalProgresReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status       string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	DeadlineFrom string `protobuf:"bytes,3,opt,name=deadline_from,json=deadlineFrom,proto3" json:"deadline_from,omitempty"`
	DeadlineTo   string `protobuf:"bytes,4,opt,name=deadline_to,json=deadlineTo,proto3" json:"deadline_to,omitempty"`
}

func (x *GoalProgresReq) Reset() {
	*x = GoalProgresReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalProgresReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalProgresReq) ProtoMessage() {}

func (x *GoalProgresReq) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalProgresReq.ProtoReflect.Descriptor instead.
func (*GoalProgresReq) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{7}
}

func (x *GoalProgresReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GoalProgresReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GoalProgresReq) GetDeadlineFrom() string {
	if x != nil {
		return x.DeadlineFrom
	}
	return ""
}

func (x *GoalProgresReq) GetDeadlineTo() string {
	if x != nil {
		return x.DeadlineTo
	}
	return ""
}

type GoalProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deadline      string  `protobuf:"bytes,1,opt,name=deadline,proto3" json:"deadline,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,2,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,3,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount,omitempty"`
	Progress      float32 `protobuf:"fixed32,4,opt,name=progress,proto3" json:"progress,omitempty"`
	GoalName      string  `protobuf:"bytes,5,opt,name=goal_name,json=goalName,proto3" json:"goal_name,omitempty"`
}

func (x *GoalProgress) Reset() {
	*x = GoalProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalProgress) ProtoMessage() {}

func (x *GoalProgress) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalProgress.ProtoReflect.Descriptor instead.
func (*GoalProgress) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{8}
}

func (x *GoalProgress) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *GoalProgress) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *GoalProgress) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *GoalProgress) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *GoalProgress) GetGoalName() string {
	if x != nil {
		return x.GoalName
	}
	return ""
}

type GoalProgresGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goals []*GoalProgress `protobuf:"bytes,1,rep,name=goals,proto3" json:"goals,omitempty"`
}

func (x *GoalProgresGet) Reset() {
	*x = GoalProgresGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_finance_protos_report_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalProgresGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalProgresGet) ProtoMessage() {}

func (x *GoalProgresGet) ProtoReflect() protoreflect.Message {
	mi := &file_finance_protos_report_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalProgresGet.ProtoReflect.Descriptor instead.
func (*GoalProgresGet) Descriptor() ([]byte, []int) {
	return file_finance_protos_report_proto_rawDescGZIP(), []int{9}
}

func (x *GoalProgresGet) GetGoals() []*GoalProgress {
	if x != nil {
		return x.Goals
	}
	return nil
}

var File_finance_protos_report_proto protoreflect.FileDescriptor

var file_finance_protos_report_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66,
	0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x20, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0c, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x47, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x01, 0x0a, 0x0c, 0x53, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x47, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x47, 0x52,
	0x65, 0x71, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x41, 0x52, 0x65, 0x73, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb1, 0x01,
	0x0a, 0x0a, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x47, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x9d, 0x01, 0x0a, 0x0a, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x47, 0x52, 0x65, 0x73,
	0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x47, 0x52, 0x65, 0x71, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x41,
	0x52, 0x65, 0x73, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xee,
	0x01, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22,
	0x4c, 0x0a, 0x0c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x47, 0x65, 0x74, 0x12,
	0x3c, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x52,
	0x0c, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x87, 0x01,
	0x0a, 0x0e, 0x47, 0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x47, 0x6f, 0x61, 0x6c,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x67, 0x6f, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x0e, 0x47, 0x6f, 0x61,
	0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x67,
	0x6f, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x32, 0x8a, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x15, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x47, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x47, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x47, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x66, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x47, 0x52, 0x65, 0x73,
	0x12, 0x41, 0x0a, 0x11, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x66,
	0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x17, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x6f,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x66,
	0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x47, 0x65, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_finance_protos_report_proto_rawDescOnce sync.Once
	file_finance_protos_report_proto_rawDescData = file_finance_protos_report_proto_rawDesc
)

func file_finance_protos_report_proto_rawDescGZIP() []byte {
	file_finance_protos_report_proto_rawDescOnce.Do(func() {
		file_finance_protos_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_finance_protos_report_proto_rawDescData)
	})
	return file_finance_protos_report_proto_rawDescData
}

var file_finance_protos_report_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_finance_protos_report_proto_goTypes = []any{
	(*SpendingGReq)(nil),     // 0: finance.SpendingGReq
	(*SpendingGRes)(nil),     // 1: finance.SpendingGRes
	(*IncomeGReq)(nil),       // 2: finance.IncomeGReq
	(*IncomeGRes)(nil),       // 3: finance.IncomeGRes
	(*BudgetPerReq)(nil),     // 4: finance.BudgetPerReq
	(*PeriodBudgetPer)(nil),  // 5: finance.PeriodBudgetPer
	(*BudgetPerGet)(nil),     // 6: finance.BudgetPerGet
	(*GoalProgresReq)(nil),   // 7: finance.GoalProgresReq
	(*GoalProgress)(nil),     // 8: finance.GoalProgress
	(*GoalProgresGet)(nil),   // 9: finance.GoalProgresGet
	(*Pagination)(nil),       // 10: finance.Pagination
	(*TransactionGARes)(nil), // 11: finance.TransactionGARes
}
var file_finance_protos_report_proto_depIdxs = []int32{
	10, // 0: finance.SpendingGReq.pagination:type_name -> finance.Pagination
	0,  // 1: finance.SpendingGRes.request:type_name -> finance.SpendingGReq
	11, // 2: finance.SpendingGRes.transactions:type_name -> finance.TransactionGARes
	10, // 3: finance.IncomeGReq.pagination:type_name -> finance.Pagination
	2,  // 4: finance.IncomeGRes.request:type_name -> finance.IncomeGReq
	11, // 5: finance.IncomeGRes.transactions:type_name -> finance.TransactionGARes
	5,  // 6: finance.BudgetPerGet.performances:type_name -> finance.PeriodBudgetPer
	8,  // 7: finance.GoalProgresGet.goals:type_name -> finance.GoalProgress
	0,  // 8: finance.ReportService.GetSpendings:input_type -> finance.SpendingGReq
	2,  // 9: finance.ReportService.GetIncomes:input_type -> finance.IncomeGReq
	4,  // 10: finance.ReportService.BudgetPerformance:input_type -> finance.BudgetPerReq
	7,  // 11: finance.ReportService.GoalProgress:input_type -> finance.GoalProgresReq
	1,  // 12: finance.ReportService.GetSpendings:output_type -> finance.SpendingGRes
	3,  // 13: finance.ReportService.GetIncomes:output_type -> finance.IncomeGRes
	6,  // 14: finance.ReportService.BudgetPerformance:output_type -> finance.BudgetPerGet
	9,  // 15: finance.ReportService.GoalProgress:output_type -> finance.GoalProgresGet
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_finance_protos_report_proto_init() }
func file_finance_protos_report_proto_init() {
	if File_finance_protos_report_proto != nil {
		return
	}
	file_finance_protos_transaction_proto_init()
	file_finance_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_finance_protos_report_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SpendingGReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SpendingGRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IncomeGReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IncomeGRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BudgetPerReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PeriodBudgetPer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*BudgetPerGet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GoalProgresReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GoalProgress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_finance_protos_report_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GoalProgresGet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_finance_protos_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_finance_protos_report_proto_goTypes,
		DependencyIndexes: file_finance_protos_report_proto_depIdxs,
		MessageInfos:      file_finance_protos_report_proto_msgTypes,
	}.Build()
	File_finance_protos_report_proto = out.File
	file_finance_protos_report_proto_rawDesc = nil
	file_finance_protos_report_proto_goTypes = nil
	file_finance_protos_report_proto_depIdxs = nil
}