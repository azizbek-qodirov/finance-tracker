package handlers

import (
	"net/http"
	"strconv"

	pb "gateway-service/genprotos" // Update with your actual package path

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CreateGoal godoc
// @Summary     Create Goal
// @Description Creates a new goal
// @Tags        Goal
// @Accept      json
// @Produce     json
// @Param       goal body     pb.GoalCReqForSwagger true "Goal creation request"
// @Success     200     {object} string "Goal created successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     500     {object} string "Server error"
// @Router      /v1/goal [post]
// @Security    BearerAuth
func (h *HTTPHandler) CreateGoal(c *gin.Context) {
	var body pb.GoalCReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	claims, _ := c.Get("claims")
	user_id := claims.(jwt.MapClaims)["user_id"].(string)

	req := pb.GoalCReq{
		UserId:       user_id,
		Name:         body.Name,
		TargetAmount: body.TargetAmount,
		Status:       body.Status,
		Deadline:     body.Deadline,
	}

	_, err := h.Goal.Create(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal created successfully"})
}

// GetGoal godoc
// @Summary     Get Goal
// @Description Get goal by ID
// @Tags        Goal
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Goal ID"
// @Success     200     {object} pb.GoalGRes "Goal details"
// @Failure     400     {object} string "Invalid goal ID"
// @Failure     404     {object} string "Goal not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/goal/{id} [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetGoal(c *gin.Context) {
	goalId := c.Param("id")

	res, err := h.Goal.GetByID(c, &pb.ByID{Id: goalId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateGoal godoc
// @Summary     Update Goal
// @Description Update goal details
// @Tags        Goal
// @Accept      json
// @Produce     json
// @Param       id     path     string         true "Goal ID"
// @Param       goal body     pb.GoalUReqForSwagger true "Goal update request"
// @Success     200     {object} string "Goal updated successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     404     {object} string "Goal not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/goal/{id} [put]
// @Security    BearerAuth
func (h *HTTPHandler) UpdateGoal(c *gin.Context) {
	goalId := c.Param("id")
	var body pb.GoalUReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req := pb.GoalUReq{
		Id:           goalId,
		Name:         body.Name,
		Status:       body.Status,
		Deadline:     body.Deadline,
		TargetAmount: body.TargetAmount,
	}

	_, err := h.Goal.Update(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal updated successfully"})
}

// UpdateGoalCurrentAmount godoc
// @Summary     Update Goal Current Amount
// @Description Update the current amount of a goal
// @Tags        Goal
// @Accept      json
// @Produce     json
// @Param       id     path     string                     true "Goal ID"
// @Param       amount body     pb.GoalCurrentAmountUReqForSwagger true "Goal current amount update request"
// @Success     200     {object} string "Goal current amount updated successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     404     {object} string "Goal not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/goal/{id}/current_amount [put]
// @Security    BearerAuth
func (h *HTTPHandler) UpdateGoalCurrentAmount(c *gin.Context) {
	goalId := c.Param("id")
	var body pb.GoalCurrentAmountUReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	req := pb.GoalCurrentAmountUReq{
		Id:            goalId,
		CurrentAmount: body.CurrentAmount,
	}
	_, err := h.Goal.UpdateCurrentAmount(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal current amount updated successfully"})
}

// DeleteGoal godoc
// @Summary     Delete Goal
// @Description Delete goal by ID
// @Tags        Goal
// @Accept      json
// @Produce     json
// @Param       id   path     string true "Goal ID"
// @Success     204     {object} string "Goal deleted successfully"
// @Failure     400     {object} string "Invalid goal ID"
// @Failure     404     {object} string "Goal not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/goal/{id} [delete]
// @Security    BearerAuth
func (h *HTTPHandler) DeleteGoal(c *gin.Context) {
	goalId := c.Param("id")

	_, err := h.Goal.Delete(c, &pb.ByID{Id: goalId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goal deleted successfully"})
}

// GetAllGoals godoc
// @Summary     List Goals
// @Description List all goals with filtering and pagination
// @Tags        Goal
// @Accept      json
// @Produce     json
// @Param       user_id       query     string         false "User ID"
// @Param       status        query     string         false "Goal status"
// @Param       target_from   query     float32        false "Minimum target amount"
// @Param       target_to     query     float32        false "Maximum target amount"
// @Param       deadline_from query     string         false "Deadline from (YYYY-MM-DD)"
// @Param       deadline_to   query     string         false "Deadline to (YYYY-MM-DD)"
// @Param       limit         query     int            false "Pagination limit"
// @Param       offset        query     int            false "Pagination offset"
// @Success     200     {object} pb.GoalGARes "List of goals"
// @Failure     400     {object} string "Invalid query parameters"
// @Failure     500     {object} string "Server error"
// @Router      /v1/goals [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetAllGoals(c *gin.Context) {
	var req pb.GoalGAReq

	req.UserId = c.Query("user_id")
	req.Status = c.Query("status")
	req.DeadlineFrom = c.Query("deadline_from")
	req.DeadlineTo = c.Query("deadline_to")

	if targetFromStr := c.Query("target_from"); targetFromStr != "" {
		targetFrom, err := strconv.ParseFloat(targetFromStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target_from parameter"})
			return
		}
		req.TargetFrom = float32(targetFrom)
	}

	if targetToStr := c.Query("target_to"); targetToStr != "" {
		targetTo, err := strconv.ParseFloat(targetToStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target_to parameter"})
			return
		}
		req.TargetTo = float32(targetTo)
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

	res, err := h.Goal.GetAll(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
