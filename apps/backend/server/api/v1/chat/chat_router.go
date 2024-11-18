package chat

import "github.com/gin-gonic/gin"

func RegisterChatRoutes(router *gin.RouterGroup, controller *ChatController) {
	chatRoutes := router.Group("/chats") 
	{
		chatRoutes.POST("/start", controller.StartChat)
		chatRoutes.GET("/:chat_id", controller.GetChat) 
		chatRoutes.PUT("/title/:chat_id", controller.UpdateTitle) 
		chatRoutes.DELETE("/:chat_id", controller.DeleteChat) 
	}

	messageRoutes := router.Group("/messages")
	{
		messageRoutes.POST("/", controller.SendMessage)
	}
}
