package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"ai-saas-schedular-server/server/common"
	"go.mongodb.org/mongo-driver/mongo"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Sender    string             `bson:"sender" json:"sender"`
	Content   string             `bson:"content" json:"content"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type Chat struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Title		string 			   `json:"title,omitempty"`
	Messages    []Message          `bson:"messages" json:"messages"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

var ChatCollection *mongo.Collection = common.OpenCollection(common.Client, "chats")
