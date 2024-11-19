package auth

import (
	"github.com/gin-gonic/gin"
	"ai-saas-schedular-server/server/middlewares"
)

func RegisterAuthRoutes(router *gin.RouterGroup, controller *AuthController) {
	authRoutes := router.Group("/")
	{
		authRoutes.POST("/login", controller.Login)
		authRoutes.POST("/forgot-password", controller.ForgotPassword)
		authRoutes.POST("/reset-password", controller.ResetPassword)

		protectedRoutes := authRoutes.Group("/", middlewares.Authentication())
		{
			protectedRoutes.GET("/me", controller.GetLoggedInUser)
		}
	}
}
