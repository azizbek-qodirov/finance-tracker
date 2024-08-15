package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"gateway-service/config"
	pb "gateway-service/genprotos" // Update with your actual package path
	kfk "gateway-service/kafka"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CreateBudget godoc
// @Summary     Create Budget
// @Description Creates a new budget
// @Tags        Budget
// @Accept      json
// @Produce     json
// @Param       budget body     pb.BudgetCReqForSwagger true "Budget creation request"
// @Success     200     {object} string "Budget created successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     500     {object} string "Server error"
// @Router      /v1/budget [post]
// @Security    BearerAuth
func (h *HTTPHandler) CreateBudget(c *gin.Context) {
	var body pb.BudgetCReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	claims, _ := c.Get("claims")
	user_id := claims.(jwt.MapClaims)["user_id"].(string)

	req := pb.BudgetCReq{
		UserId:     user_id,
		CategoryId: body.CategoryId,
		Amount:     body.Amount,
		Period:     body.Period,
		StartDate:  body.StartDate,
		EndDate:    body.EndDate,
	}

	go func() {
		kafkaTopic := "budget-create"
		err := kfk.ProduceKafkaMessage(kafkaTopic, kafkaTopic+":"+user_id, &req, config.Load().KAFKA_BROKER)
		if err != nil {
			fmt.Println("Error producing Kafka message:", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Budget created successfully"})
}

// GetBudget godoc
// @Summary     Get Budget
// @Description Get budget by ID
// @Tags        Budget
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Budget ID"
// @Success     200     {object} pb.BudgetGRes "Budget details"
// @Failure     400     {object} string "Invalid budget ID"
// @Failure     404     {object} string "Budget not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/budget/{id} [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetBudget(c *gin.Context) {
	budgetId := c.Param("id")

	res, err := h.Budget.GetByID(c, &pb.ByID{Id: budgetId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateBudget godoc
// @Summary     Update Budget
// @Description Update budget details
// @Tags        Budget
// @Accept      json
// @Produce     json
// @Param       id     path     string         true "Budget ID"
// @Param       budget body     pb.BudgetUReqForSwagger true "Budget update request"
// @Success     200     {object} string "Budget updated successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     404     {object} string "Budget not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/budget/{id} [put]
// @Security    BearerAuth
func (h *HTTPHandler) UpdateBudget(c *gin.Context) {
	budgetId := c.Param("id")
	var body pb.BudgetUReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req := pb.BudgetUReq{
		Id:        budgetId,
		Amount:    body.Amount,
		Period:    body.Period,
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
	}

	go func() {
		kafkaTopic := "budget-update"
		err := kfk.ProduceKafkaMessage(kafkaTopic, kafkaTopic+":"+budgetId, &req, config.Load().KAFKA_BROKER)
		if err != nil {
			fmt.Println("Error producing Kafka message:", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Budget updated successfully"})
}

// DeleteBudget godoc
// @Summary     Delete Budget
// @Description Delete budget by ID
// @Tags        Budget
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Budget ID"
// @Success     204     {object} string "Budget deleted successfully"
// @Failure     400     {object} string "Invalid budget ID"
// @Failure     404     {object} string "Budget not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/budget/{id} [delete]
// @Security    BearerAuth
func (h *HTTPHandler) DeleteBudget(c *gin.Context) {
	budgetId := c.Param("id")

	go func() {
		kafkaTopic := "budget-delete"
		err := kfk.ProduceKafkaMessage(kafkaTopic, kafkaTopic+":"+budgetId, &pb.ByID{Id: budgetId}, config.Load().KAFKA_BROKER)
		if err != nil {
			fmt.Println("Error producing Kafka message:", err)
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "Budget deleted successfully"})
}

// GetAllBudgets godoc
// @Summary     List Budgets
// @Description List all budgets with filtering and pagination
// @Tags        Budget
// @Accept      json
// @Produce     json
// @Param       user_id      query     string         false "User ID"
// @Param       category_id  query     string         false "Category ID"
// @Param       amount_from  query     float32        false "Minimum amount"
// @Param       amount_to    query     float32        false "Maximum amount"
// @Param       period       query     string         false "Budget period"
// @Param       limit        query     int            false "Pagination limit"
// @Param       offset       query     int            false "Pagination offset"
// @Success     200     {object} pb.BudgetGARes "List of budgets"
// @Failure     400     {object} string "Invalid query parameters"
// @Failure     500     {object} string "Server error"
// @Router      /v1/budgets [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetAllBudgets(c *gin.Context) {
	var req pb.BudgetGAreq

	req.UserId = c.Query("user_id")
	req.CategoryId = c.Query("category_id")
	req.Period = c.Query("period")

	if amountFromStr := c.Query("amount_from"); amountFromStr != "" {
		amountFrom, err := strconv.ParseFloat(amountFromStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount_from parameter"})
			return
		}
		req.AmountFrom = float32(amountFrom)
	}

	if amountToStr := c.Query("amount_to"); amountToStr != "" {
		amountTo, err := strconv.ParseFloat(amountToStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount_to parameter"})
			return
		}
		req.AmountTo = float32(amountTo)
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

	res, err := h.Budget.GetAll(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
