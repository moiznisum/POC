package models

import (
	"context"
	"errors"
	"log"
	"time"

	"ai-saas-schedular-server/server/common"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	FirstName        string             `json:"first_name" validate:"required,min=2,max=100"`
	LastName         string             `json:"last_name" validate:"required,min=2,max=100"`
	Username         string             `json:"username" validate:"required,min=3,max=100,unique"`
	Email            string             `json:"email" validate:"required,email,unique"`
	Password         string             `json:"password" validate:"required,min=6"`
	Phone            string             `json:"phone" validate:"required,unique"`
	Address          *string            `json:"address,omitempty"`
	City             *string            `json:"city,omitempty"`
	Country          *string            `json:"country,omitempty"`
	PostalCode       *string            `json:"postal_code,omitempty"`
	ProfilePicture   *string            `json:"profile_picture,omitempty"`
	ResetPasswordToken *string           `json:"reset_password_token,omitempty"`

	UserType         string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken     *string            `json:"refresh_token,omitempty"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

var UserCollection *mongo.Collection = common.OpenCollection(common.Client, "users")

func CreateUser(userData map[string]interface{}) (*User, error) {
	firstName, ok := userData["first_name"].(string)
	if !ok || firstName == "" {
		return nil, errors.New("first_name is required")
	}

	lastName, ok := userData["last_name"].(string)
	if !ok || lastName == "" {
		return nil, errors.New("last_name is required")
	}

	username, ok := userData["username"].(string)
	if !ok || username == "" {
		return nil, errors.New("username is required")
	}

	email, ok := userData["email"].(string)
	if !ok || email == "" {
		return nil, errors.New("email is required")
	}

	phone, ok := userData["phone"].(string)
	if !ok || phone == "" {
		return nil, errors.New("phone is required")
	}

	password, ok := userData["password"].(string)
	if !ok || password == "" {
		return nil, errors.New("password is required")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	var address, city, country, postalCode, profilePicture *string

	if val, ok := userData["address"].(string); ok && val != "" {
		address = &val
	}
	if val, ok := userData["city"].(string); ok && val != "" {
		city = &val
	}
	if val, ok := userData["country"].(string); ok && val != "" {
		country = &val
	}
	if val, ok := userData["postal_code"].(string); ok && val != "" {
		postalCode = &val
	}
	if val, ok := userData["profile_picture"].(string); ok && val != "" {
		profilePicture = &val
	}

	userType, ok := userData["user_type"].(string)
	if !ok || (userType != "ADMIN" && userType != "USER") {
		return nil, errors.New("user_type must be 'ADMIN' or 'USER'")
	}

	user := User{
		FirstName:      firstName,
		LastName:       lastName,
		Username:       username,
		Email:          email,
		Phone:          phone,
		Password:       string(hashedPassword),
		UserType:       userType,
		Address:        address,
		City:           city,
		Country:        country,
		PostalCode:     postalCode,
		ProfilePicture: profilePicture,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	result, err := UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error inserting user:", err)
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)

	log.Println("User inserted:", user)

	return &user, nil
}

func FindAllUsers() ([]User, error) {
	var users []User
	cursor, err := UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func FindUserByID(id string) (*User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var user User
	err = UserCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUserByID(id string, userData map[string]interface{}) (*User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	update := bson.M{"$set": userData}
	_, err = UserCollection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return FindUserByID(id)
}

func DeleteUserByID(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	_, err = UserCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	err := UserCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByUsername(username string) (*User, error) {
	var user User
	err := UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByPhone(phone string) (*User, error) {
	var user User
	err := UserCollection.FindOne(context.Background(), bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUserIndexes() {
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	phoneIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "phone", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := UserCollection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{emailIndex, phoneIndex, usernameIndex})
	if err != nil {
		log.Fatalf("Failed to create indexes: %v", err)
	}
}

func (u *User) AuthenticatePassword(password string) bool {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

