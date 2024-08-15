package handlers

import (
	"net/http"
	"strconv"

	pb "gateway-service/genprotos" // Update with your actual package path

	"github.com/gin-gonic/gin"
)

// GetSpendings godoc
// @Summary     Get Spendings Report
// @Description Get a report of spendings with filtering and pagination
// @Tags        Report
// @Accept      json
// @Produce     json
// @Param       user_id      query     string         false "User ID"
// @Param       date_from    query     string         false "Start date (YYYY-MM-DD)"
// @Param       date_to      query     string         false "End date (YYYY-MM-DD)"
// @Param       category_id  query     string         false "Category ID"
// @Param       limit        query     int            false "Pagination limit"
// @Param       offset       query     int            false "Pagination offset"
// @Success     200     {object} pb.SpendingGRes "Spending report"
// @Failure     400     {object} string "Invalid query parameters"
// @Failure     500     {object} string "Server error"
// @Router      /v1/report/spendings [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetSpendings(c *gin.Context) {
	var req pb.SpendingGReq

	req.UserId = c.Query("user_id")
	req.DateFrom = c.Query("date_from")
	req.DateTo = c.Query("date_to")
	req.CategoryId = c.Query("category_id")

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

	res, err := h.Report.GetSpendings(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetIncomes godoc
// @Summary     Get Incomes Report
// @Description Get a report of incomes with filtering and pagination
// @Tags        Report
// @Accept      json
// @Produce     json
// @Param       user_id      query     string         false "User ID"
// @Param       date_from    query     string         false "Start date (YYYY-MM-DD)"
// @Param       date_to      query     string         false "End date (YYYY-MM-DD)"
// @Param       category_id  query     string         false "Category ID"
// @Param       limit        query     int            false "Pagination limit"
// @Param       offset       query     int            false "Pagination offset"
// @Success     200     {object} pb.IncomeGRes "Income report"
// @Failure     400     {object} string "Invalid query parameters"
// @Failure     500     {object} string "Server error"
// @Router      /v1/report/incomes [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetIncomes(c *gin.Context) {
	var req pb.IncomeGReq

	req.UserId = c.Query("user_id")
	req.DateFrom = c.Query("date_from")
	req.DateTo = c.Query("date_to")
	req.CategoryId = c.Query("category_id")

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

	res, err := h.Report.GetIncomes(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetBudgetPerformance godoc
// @Summary     Get Budget Performance Report
// @Description Get a report of budget performance with filtering
// @Tags        Report
// @Accept      json
// @Produce     json
// @Param       user_id      query     string         false "User ID"
// @Param       category_id  query     string         false "Category ID"
// @Param       period       query     string         false "Budget period"
// @Param       start_date   query     string         false "Start date (YYYY-MM-DD)"
// @Param       end_date     query     string         false "End date (YYYY-MM-DD)"
// @Success     200     {object} pb.BudgetPerGet "Budget performance report"
// @Failure     400     {object} string "Invalid query parameters"
// @Failure     500     {object} string "Server error"
// @Router      /v1/report/budget-performance [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetBudgetPerformance(c *gin.Context) {
	var req pb.BudgetPerReq

	req.UserId = c.Query("user_id")
	req.CategoryId = c.Query("category_id")
	req.Period = c.Query("period")
	req.StartDate = c.Query("start_date")
	req.EndDate = c.Query("end_date")

	res, err := h.Report.BudgetPerformance(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetGoalProgress godoc
// @Summary     Get Goal Progress Report
// @Description Get a report of goal progress with filtering
// @Tags        Report
// @Accept      json
// @Produce     json
// @Param       user_id       query     string         false "User ID"
// @Param       status        query     string         false "Goal status"
// @Param       deadline_from query     string         false "Deadline from (YYYY-MM-DD)"
// @Param       deadline_to   query     string         false "Deadline to (YYYY-MM-DD)"
// @Success     200     {object} pb.GoalProgresGet "Goal progress report"
// @Failure     400     {object} string "Invalid query parameters"
// @Failure     500     {object} string "Server error"
// @Router      /v1/report/goal-progress [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetGoalProgress(c *gin.Context) {
	var req pb.GoalProgresReq

	req.UserId = c.Query("user_id")
	req.Status = c.Query("status")
	req.DeadlineFrom = c.Query("deadline_from")
	req.DeadlineTo = c.Query("deadline_to")

	res, err := h.Report.GoalProgress(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
