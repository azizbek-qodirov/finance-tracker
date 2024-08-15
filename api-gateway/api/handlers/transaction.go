package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"gateway-service/config"
	pb "gateway-service/genprotos"
	kfk "gateway-service/kafka"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CreateTransaction godoc
// @Summary     Create Transaction
// @Description Creates a new transaction
// @Tags        Transaction
// @Accept      json
// @Produce     json
// @Param       transaction body     pb.TransactionCReqForSwagger true "Transaction creation request"
// @Success     200     {object} string "Transaction created successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     500     {object} string "Server error"
// @Security BearerAuth
// @Router      /v1/transaction [post]
func (h *HTTPHandler) CreateTransaction(c *gin.Context) {
	var body pb.TransactionCReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	claims, _ := c.Get("claims")
	user_id := claims.(jwt.MapClaims)["user_id"].(string)

	req := pb.TransactionCReq{
		UserId:      user_id,
		AccountId:   body.AccountId,
		CategoryId:  body.CategoryId,
		Amount:      body.Amount,
		Type:        body.Type,
		Description: body.Description,
	}

	go func() {
		kafkaTopic := "transaction-create"
		kafkaKey := "transaction-create-" + req.UserId
		err := kfk.ProduceKafkaMessage(kafkaTopic, kafkaKey, &req, config.Load().KAFKA_BROKER)
		if err != nil {
			fmt.Println("Error producing Kafka message:", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully"})
}

// GetTransaction godoc
// @Summary     Get Transaction
// @Description Get transaction by ID
// @Tags        Transaction
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Transaction ID"
// @Success     200     {object} pb.TransactionGRes "Transaction details"
// @Failure     400     {object} string "Invalid transaction ID"
// @Failure     404     {object} string "Transaction not found"
// @Failure     500     {object} string "Server error"
// @Security BearerAuth
// @Router      /v1/transaction/{id} [get]
func (h *HTTPHandler) GetTransaction(c *gin.Context) {
	transactionId := c.Param("id")

	res, err := h.Transaction.GetByID(c, &pb.ByID{Id: transactionId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteTransaction godoc
// @Summary     Delete Transaction
// @Description Delete transaction by ID
// @Tags        Transaction
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Transaction ID"
// @Success     204     {object} string "Transaction deleted successfully"
// @Failure     400     {object} string "Invalid transaction ID"
// @Failure     404     {object} string "Transaction not found"
// @Failure     500     {object} string "Server error"
// @Security BearerAuth
// @Router      /v1/transaction/{id} [delete]
func (h *HTTPHandler) DeleteTransaction(c *gin.Context) {
	transactionId := c.Param("id")

	go func() {
		kafkaTopic := "transaction-delete"
		kafkaKey := "transaction-delete-" + transactionId
		err := kfk.ProduceKafkaMessage(kafkaTopic, kafkaKey, &pb.ByID{Id: transactionId}, config.Load().KAFKA_BROKER)
		if err != nil {
			fmt.Println("Error producing Kafka message:", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}

// GatAllTransactions godoc
// @Summary     List Transactions
// @Description List all transactions with filtering and pagination
// @Tags        Transaction
// @Accept      json
// @Produce     json
// @Param       user_id      query     string         false "User ID"
// @Param       account_id   query     string         false "Account ID"
// @Param       category_id  query     string         false "Category ID"
// @Param       amount       query     float32        false "Amount"
// @Param       type         query     string         false "Transaction type"
// @Param       date_from    query     string         false "Start date (YYYY-MM-DD)"
// @Param       date_to      query     string         false "End date (YYYY-MM-DD)"
// @Param       limit        query     int            false "Pagination limit"
// @Param       offset       query     int            false "Pagination offset"
// @Success     200     {object} pb.TransactionGARes "List of transactions"
// @Failure     400     {object} string "Invalid query parameters"
// @Failure     500     {object} string "Server error"
// @Security BearerAuth
// @Router      /v1/transactions [get]
func (h *HTTPHandler) GatAllTransactions(c *gin.Context) {
	var req pb.TransactionGAReq

	req.UserId = c.GetString("user_id")
	req.AccountId = c.Query("account_id")
	req.CategoryId = c.Query("category_id")
	req.Type = c.Query("type")
	req.DateFrom = c.Query("date_from")
	req.DateTo = c.Query("date_to")

	if amountStr := c.Query("amount"); amountStr != "" {
		amount, err := strconv.ParseFloat(amountStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount parameter"})
			return
		}
		req.Amount = float32(amount)
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
		req.Pagination = &pb.Pagination{Limit: int64(limit)}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
		if req.Pagination == nil {
			req.Pagination = &pb.Pagination{}
		}
		req.Pagination.Offset = int64(offset)
	}

	res, err := h.Transaction.GetAll(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
