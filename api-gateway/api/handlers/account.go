package handlers

import (
	"net/http"

	pb "gateway-service/genprotos"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GetAccount godoc
// @Summary     Get Account
// @Description Get user account by User ID
// @Tags        Account
// @Accept      json
// @Produce     json
// @Success     200     {object} pb.AccountGRes "Account details"
// @Failure     400     {object} string "Invalid user ID"
// @Failure     404     {object} string "Account not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/account [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetAccount(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user_id := claims.(jwt.MapClaims)["user_id"].(string)

	res, err := h.Account.GetAccount(c, &pb.ByUserID{UserId: user_id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetBalance godoc
// @Summary     Get Account Balance
// @Description Get user account balance by User ID
// @Tags        Account
// @Accept      json
// @Produce     json
// @Success     200     {object} pb.AccountBalanceGRes "Account balance"
// @Failure     400     {object} string "Invalid user ID"
// @Failure     404     {object} string "Account not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/account/balance [get]
// @Security    BearerAuth
func (h *HTTPHandler) GetBalance(c *gin.Context) {
	claims, _ := c.Get("claims")
	userId := claims.(jwt.MapClaims)["user_id"].(string)

	res, err := h.Account.GetBalance(c, &pb.ByUserID{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateAccount godoc
// @Summary     Update Account
// @Description Update user account details
// @Tags        Account
// @Accept      json
// @Produce     json
// @Param       account body     pb.AccountUReq true "Account update request"
// @Success     200     {object} string "Account updated successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     404     {object} string "Account not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/account [put]
// @Security    BearerAuth
func (h *HTTPHandler) UpdateAccount(c *gin.Context) {
	var req pb.AccountUReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.Account.UpdateAccount(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account updated successfully"})
}

// UpdateBalance godoc
// @Summary     Update Account Balance
// @Description Update user account balance
// @Tags        Account
// @Accept      json
// @Produce     json
// @Param       balance body     pb.AccountBalanceUReqForSwagger true "Account balance update request"
// @Success     200     {object} string "Account balance updated successfully"
// @Failure     400     {object} string "Invalid request payload"
// @Failure     404     {object} string "Account not found"
// @Failure     500     {object} string "Server error"
// @Router      /v1/account/balance [put]
// @Security    BearerAuth
func (h *HTTPHandler) UpdateBalance(c *gin.Context) {
	var body pb.AccountBalanceUReqForSwagger
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	claims, _ := c.Get("claims")
	user_id := claims.(jwt.MapClaims)["user_id"].(string)

	req := pb.AccountBalanceUReq{
		UserId:  user_id,
		Balance: body.Balance,
	}

	_, err := h.Account.UpdateBalance(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account balance updated successfully"})
}
