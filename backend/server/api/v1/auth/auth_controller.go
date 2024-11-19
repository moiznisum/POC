package auth

import (
	"net/http"
	"ai-saas-schedular-server/server/common"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	common.BaseController
	AuthService *AuthService
}

type LoginRequest struct {
    Email    *string `json:"email"`
    Username *string `json:"username"`
    Phone    *string `json:"phone"`
    Password string  `json:"password"`
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthService: NewAuthService(),
	}
}

func (ac *AuthController) Login(c *gin.Context) {
    var credentials LoginRequest

    if err := c.ShouldBindJSON(&credentials); err != nil {
        common.Logger.Error("Invalid request body for login: ", err)
        ac.Response(c.Writer, nil, http.StatusBadRequest, "Invalid request body")
        return
    }

    common.Logger.Info("Login attempt with credentials: ", credentials)

    response, err := ac.AuthService.Login(credentials)
    if err != nil {
        common.Logger.Error("Login error: ", err)
        ac.Response(c.Writer, nil, http.StatusUnauthorized, err.Error())
        return
    }

    c.SetCookie("token", response["accessToken"].(string), 604800, "/", "", false, true)
    common.Logger.Info("User logged in successfully: ", response["id"])
    ac.Response(c.Writer, response, http.StatusOK, "Login successful")
}

func (ac *AuthController) GetLoggedInUser(c *gin.Context) {
	email, _ := c.Get("email")
	firstName, _ := c.Get("first_name")
	lastName, _ := c.Get("last_name")
	uid, _ := c.Get("uid")
	userType, _ := c.Get("user_type")

	userData := map[string]interface{}{
		"email":      email,
		"first_name": firstName,
		"last_name":  lastName,
		"uid":        uid,
		"user_type":  userType,
	}

	common.Logger.Info("Fetched logged-in user data for email: ", email)
	ac.Response(c.Writer, userData, http.StatusOK, "User details fetched successfully")
}

func (ac *AuthController) ForgotPassword(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
		Host  string `json:"host"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		common.Logger.Error("Invalid request body for forgot password: ", err)
		ac.Response(c.Writer, nil, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := ac.AuthService.ForgotPassword(body.Email, body.Host)
	if err != nil {
		common.Logger.Error("Forgot password error: ", err)
		ac.Response(c.Writer, nil, http.StatusInternalServerError, err.Error())
		return
	}

	common.Logger.Info("Forgot password email sent for: ", body.Email)
	ac.Response(c.Writer, response, http.StatusOK, "Email sent successfully!")
}

func (ac *AuthController) ResetPassword(c *gin.Context) {
	var body struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		common.Logger.Error("Invalid request body for reset password: ", err)
		ac.Response(c.Writer, nil, http.StatusBadRequest, "Invalid request body")
		return
	}

	response, err := ac.AuthService.ResetPassword(body.Token, body.Password)
	if err != nil {
		common.Logger.Error("Reset password error: ", err)
		ac.Response(c.Writer, nil, http.StatusInternalServerError, err.Error())
		return
	}

	common.Logger.Info("Password reset successfully for token: ", body.Token)
	ac.Response(c.Writer, response, http.StatusOK, "Password reset successfully!")
}
