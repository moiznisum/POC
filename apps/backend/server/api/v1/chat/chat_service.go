package chat

import (
	"context"
	"errors"
	"os"
	"time"
	"fmt"

	"ai-saas-schedular-server/server/models"
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"ai-saas-schedular-server/server/common"
)

type ChatService struct {
	openaiClient *openai.Client
}

func NewChatService() *ChatService {
	openaiClient := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return &ChatService{
		openaiClient: openaiClient,
	}
}

func (cs *ChatService) StartChat(userID primitive.ObjectID) (*models.Chat, error) {
	chat := models.Chat{
		UserID:    userID,
		Messages:  []models.Message{},
		Title:	   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := models.ChatCollection.InsertOne(context.Background(), chat)
	if err != nil {
		return nil, err
	}

	chat.ID = result.InsertedID.(primitive.ObjectID)
	return &chat, nil
}

func (cs *ChatService) AddMessageToChat(chatID primitive.ObjectID, sender, content string) (*models.Chat, error) {
	var chat models.Chat
	err := models.ChatCollection.FindOne(context.Background(), bson.M{"_id": chatID}).Decode(&chat)
	if err != nil {
		return nil, errors.New("chat not found")
	}

	message := models.Message{
		ID:        primitive.NewObjectID(),
		Sender:    sender,
		Content:   content,
		CreatedAt: time.Now(),
	}

	chat.Messages = append(chat.Messages, message)

	update := bson.M{
		"$set": bson.M{
			"messages":   chat.Messages,
			"updated_at": time.Now(),
		},
	}

	if chat.Title == "" {
		update["$set"].(bson.M)["title"] = content;
	}

	_, err = models.ChatCollection.UpdateOne(context.Background(), bson.M{"_id": chatID}, update)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (cs *ChatService) GetChat(chatID primitive.ObjectID) (*models.Chat, error) {
	var chat models.Chat
	err := models.ChatCollection.FindOne(context.Background(), bson.M{"_id": chatID}).Decode(&chat)
	if err != nil {
		return nil, errors.New("chat not found")
	}
	return &chat, nil
}

func (cs *ChatService) UpdateTitle(chatID primitive.ObjectID, newTitle string) (*models.Chat, error) {
	var chat models.Chat
	err := models.ChatCollection.FindOne(context.Background(), bson.M{"_id": chatID}).Decode(&chat)
	if err != nil {
		return nil, errors.New("chat not found")
	}

	update := bson.M{
		"$set": bson.M{
			"title":     newTitle,
			"updated_at": time.Now(),
		},
	}

	_, err = models.ChatCollection.UpdateOne(context.Background(), bson.M{"_id": chatID}, update)
	if err != nil {
		return nil, err
	}

	err = models.ChatCollection.FindOne(context.Background(), bson.M{"_id": chatID}).Decode(&chat)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (cs *ChatService) SendMessageToAI(message string) (string, error) {
	resp, err := cs.openaiClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
		},
		Temperature: 0.7,
	})

	if err != nil {
		common.Logger.Error("Error calling OpenAI API: ", err)
		return "", fmt.Errorf("error calling OpenAI API: %w", err)
	}

	common.Logger.Info("OpenAI API Response: ", resp)

	if len(resp.Choices) == 0 {
		common.Logger.Error("No AI response received.")
		return "", fmt.Errorf("no AI response received")
	}

	return resp.Choices[0].Message.Content, nil
}

func (cs *ChatService) DeleteChat(chatID primitive.ObjectID) error {
	_, err := models.ChatCollection.DeleteOne(context.Background(), bson.M{"_id": chatID})
	if err != nil {
		return err
	}

	return nil
}
