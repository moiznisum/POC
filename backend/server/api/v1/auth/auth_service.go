package auth

import (
	"context"
	"errors"
	"time"

	"ai-saas-schedular-server/server/common"
	"ai-saas-schedular-server/server/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
)


type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Login(credentials LoginRequest) (map[string]interface{}, error) {
    var user *models.User
    var err error

    if credentials.Email != nil && *credentials.Email != "" {
        user, err = models.FindUserByEmail(*credentials.Email)
    }

    if user == nil && credentials.Username != nil && *credentials.Username != "" {
        user, err = models.FindUserByUsername(*credentials.Username)
    }

    if user == nil && credentials.Phone != nil && *credentials.Phone != "" {
        user, err = models.FindUserByPhone(*credentials.Phone)
    }

    if user == nil || err != nil {
        common.Logger.Warn("Login failed: user not found for credentials")
        return nil, errors.New("user not exists")
    }

    if !user.AuthenticatePassword(credentials.Password) {
        common.Logger.Warn("Invalid password attempt")
        return nil, errors.New("email or password invalid")
    }

    userID := user.ID.Hex()
    token := as.generateToken(userID)

    return map[string]interface{}{
        "id":             user.ID,
        "email":          user.Email,
        "firstName":      user.FirstName,
        "lastName":       user.LastName,
        "username":       user.Username,
        "phone":          user.Phone,
        "accessToken":    token,
        "profilePicture": user.ProfilePicture,
    }, nil
}


func (as *AuthService) ForgotPassword(email, host string) (string, error) {
	user, err := models.FindUserByEmail(email)
	if err != nil {
		common.Logger.Warn("Forgot password attempt for non-existent email: ", email)
		return "", errors.New("no user found with that email address")
	}

	token := as.generateToken(user.ID.Hex())
	resetLink := host + "/reset-password/" + token

	return "Password reset email sent! Use the link: " + resetLink, nil
}

func (as *AuthService) ResetPassword(token, password string) (string, error) {
	claims, err := as.validateToken(token)
	if err != nil {
		return "", errors.New("invalid token")
	}

	user, err := models.FindUserByID(claims.Id)
	if err != nil {
		return "", errors.New("user not found")
	}

	update := bson.M{"$set": bson.M{"password": password}}
	_, err = models.UserCollection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, update)
	if err != nil {
		return "", errors.New("failed to reset password")
	}

	return "Password reset successfully!", nil
}

// private methods
func (as *AuthService) findUserByEmailOrUsernameOrPhone(identifier string) (*models.User, error) {
	user, err := models.FindUserByEmail(identifier)
	if err == nil {
		return user, nil
	}

	user, err = models.FindUserByUsername(identifier)
	if err == nil {
		return user, nil
	}

	user, err = models.FindUserByPhone(identifier)
	if err == nil {
		return user, nil
	}

	return nil, errors.New("user not exists")
}

func (as *AuthService) generateToken(userID string) string {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(common.SECRET_KEY))
	return tokenString
}

func (as *AuthService) validateToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
