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

	lesson := protected.Group("/lesson")
	{
		lesson.GET("/:id", h.LessonGet)
		protected.GET("/lessons", h.LessonGetAll)
	}

	return router
}
