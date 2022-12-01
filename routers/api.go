package routers

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/controllers"
)

func StartServer() *gin.Engine{
	router := gin.Default()

	orderRouter := router.Group("/order")
	{
		orderRouter.GET("/:orderId", controllers.OrderGet)
		orderRouter.GET("/", controllers.OrderGetList)
		orderRouter.POST("/", controllers.OrderCreate)
		orderRouter.PUT("/:orderId", controllers.OrderUpdate)
		orderRouter.DELETE("/:orderId", controllers.OrderDelete)
	}

	return router
}