package v1

import (
	"ai-saas-schedular-server/server/api/v1/auth"
	"ai-saas-schedular-server/server/api/v1/user"
	"ai-saas-schedular-server/server/api/v1/chat"

	"github.com/gin-gonic/gin"
)

func SetupV1Routes(router *gin.RouterGroup) {
	authController := auth.NewAuthController()
	userController := user.NewUserController()
	chatController := chat.NewChatController()

	auth.RegisterAuthRoutes(router.Group("/auth"), authController)
	user.RegisterUserRoutes(router.Group("/users"), userController)
	chat.RegisterChatRoutes(router.Group("/chat-ai"), chatController)

}
