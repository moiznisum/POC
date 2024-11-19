package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, controller *UserController) {
	userRoutes := router.Group("/") // Proper declaration of userRoutes
	{
		userRoutes.GET("/", controller.Get)
		userRoutes.POST("/", controller.Create)
		userRoutes.GET("/:id", controller.GetByID)
		userRoutes.PUT("/:id", controller.Update)
		userRoutes.DELETE("/:id", controller.Delete)
	}
}
