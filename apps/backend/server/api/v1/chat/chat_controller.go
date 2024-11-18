package chat

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"ai-saas-schedular-server/server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatController struct {
	common.BaseController
	ChatService *ChatService
}

func NewChatController() *ChatController {
	return &ChatController{
		ChatService: NewChatService(),
	}
}

func (cc *ChatController) StartChat(c *gin.Context) {
	var requestBody struct {
		UserID string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		cc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid request body")
		return
	}

	userObjectID, err := primitive.ObjectIDFromHex(requestBody.UserID)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid User ID")
		return
	}

	chat, err := cc.ChatService.StartChat(userObjectID)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusInternalServerError, "Failed to start chat")
		return
	}

	cc.Response(c.Writer, chat, http.StatusOK, "Chat started successfully")
}

func (cc *ChatController) SendMessage(c *gin.Context) {
	var requestBody struct {
		ChatID  string `json:"chat_id" binding:"required"`
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		cc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid request body")
		return
	}

	chatObjectID, err := primitive.ObjectIDFromHex(requestBody.ChatID)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid Chat ID")
		return
	}

	chat, err := cc.ChatService.AddMessageToChat(chatObjectID, "user", requestBody.Message)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusInternalServerError, "Failed to send message")
		return
	}

	aiResponse, err := cc.ChatService.SendMessageToAI(requestBody.Message)
	if err != nil {
		common.Logger.Error("Failed to get AI response: ", err)
		cc.Response(c.Writer, nil, http.StatusInternalServerError, "Failed to get AI response")
		return
	}

	_, err = cc.ChatService.AddMessageToChat(chatObjectID, "ai", aiResponse)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusInternalServerError, "Failed to save AI response")
		return
	}

	cc.Response(c.Writer, chat, http.StatusOK, "Message sent successfully")
}

func (cc *ChatController) GetChat(c *gin.Context) {
	chatID := c.Param("chat_id")

	chatObjectID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid Chat ID")
		return
	}

	chat, err := cc.ChatService.GetChat(chatObjectID)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusNotFound, "Chat not found")
		return
	}

	cc.Response(c.Writer, chat, http.StatusOK, "Chat fetched successfully")
}

func (cc *ChatController) UpdateTitle(c *gin.Context) {
	chatID := c.Param("chat_id")
	
	var requestBody struct {
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		cc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid request body")
		return
	}

	chatObjectID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid Chat ID")
		return
	}

	chat, err := cc.ChatService.UpdateTitle(chatObjectID, requestBody.Title)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusInternalServerError, "Failed to update title")
		return
	}

	cc.Response(c.Writer, chat, http.StatusOK, "Title updated successfully")
}

func (cc *ChatController) DeleteChat(c *gin.Context) {
	chatID := c.Param("chat_id")

	chatObjectID, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid Chat ID")
		return
	}

	err = cc.ChatService.DeleteChat(chatObjectID)
	if err != nil {
		cc.Response(c.Writer, nil, http.StatusInternalServerError, "Failed to delete chat")
		return
	}

	cc.Response(c.Writer, nil, http.StatusOK, "Chat deleted successfully")
}

