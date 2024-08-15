package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gateway-service/api/docs"
	"gateway-service/api/handlers"
	"gateway-service/api/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connB *grpc.ClientConn) *gin.Engine {
	h := handlers.NewHandler(connB)
	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.Use(middleware.IsUserMiddleware())

	v1 := protected.Group("/v1")
	{
		account := v1.Group("/account")
		{
			account.GET("/", h.GetAccount)
			account.GET("/balance", h.GetBalance)
			account.PUT("/", h.UpdateAccount)
			account.PUT("/balance", h.UpdateBalance)
		}

		budget := v1.Group("/budget")
		{
			budget.POST("/", h.CreateBudget)
			budget.GET("/:id", h.GetBudget)
			budget.PUT("/:id", h.UpdateBudget)
			budget.DELETE("/:id", h.DeleteBudget)
			v1.GET("/budgets", h.GetAllBudgets)
		}

		category := v1.Group("/category")
		{
			category.POST("/", h.CreateCategory)
			category.GET("/:id", h.GetCategory)
			category.PUT("/:id", h.UpdateCategory)
			category.DELETE("/:id", h.DeleteCategory)
			v1.GET("/categories", h.GetAllCategories)
		}

		goal := v1.Group("/goal")
		{
			goal.POST("/", h.CreateGoal)
			goal.GET("/:id", h.GetGoal)
			goal.PUT("/:id", h.UpdateGoal)
			goal.PUT("/:id/current_amount", h.UpdateGoalCurrentAmount)
			goal.DELETE("/:id", h.DeleteGoal)
			v1.GET("/goals", h.GetAllGoals)
		}

		transaction := v1.Group("/transaction")
		{
			transaction.POST("/", h.CreateTransaction)
			transaction.GET("/:id", h.GetTransaction)
			transaction.DELETE("/:id", h.DeleteTransaction)
			v1.GET("/transactions", h.GatAllTransactions)
		}

		report := v1.Group("/report")
		{
			report.GET("/spendings", h.GetSpendings)
			report.GET("/incomes", h.GetIncomes)
			report.GET("/budget-performance", h.GetBudgetPerformance)
			report.GET("/goal-progress", h.GetGoalProgress)
		}

	}
	return router
}
