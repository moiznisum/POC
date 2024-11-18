package user

import (
	"net/http"
	"ai-saas-schedular-server/server/common"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	common.BaseController
	UserService *UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: NewUserService(),
	}
}

// Get retrieves all users
func (uc *UserController) Get(c *gin.Context) {
	users, err := uc.UserService.Get()
	if err != nil {
		common.Logger.Error("Error fetching users: ", err)
		uc.Response(c.Writer, nil, http.StatusInternalServerError, "Error fetching users")
		return
	}
	uc.Response(c.Writer, users, http.StatusOK, "Users retrieved successfully")
}

// Create creates a new user
func (uc *UserController) Create(c *gin.Context) {
	var userData map[string]interface{}
	if err := c.ShouldBindJSON(&userData); err != nil {
		common.Logger.Error("Invalid user data: ", err)
		uc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := uc.UserService.Create(userData)
	if err != nil {
		common.Logger.Error("Error creating user: ", err)
		uc.Response(c.Writer, nil, http.StatusInternalServerError, "Error creating user")
		return
	}
	uc.Response(c.Writer, user, http.StatusOK, "User created successfully")
}

// GetByID retrieves a user by ID
func (uc *UserController) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.UserService.GetByID(id)
	if err != nil {
		common.Logger.Error("Error fetching user by ID: ", err)
		uc.Response(c.Writer, nil, http.StatusInternalServerError, err.Error())
		return
	}
	uc.Response(c.Writer, user, http.StatusOK, "User retrieved successfully")
}

// Update updates a user's details
func (uc *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var userData map[string]interface{}
	if err := c.ShouldBindJSON(&userData); err != nil {
		common.Logger.Error("Invalid update data: ", err)
		uc.Response(c.Writer, nil, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := uc.UserService.Update(id, userData)
	if err != nil {
		common.Logger.Error("Error updating user: ", err)
		uc.Response(c.Writer, nil, http.StatusInternalServerError, err.Error())
		return
	}
	uc.Response(c.Writer, user, http.StatusOK, "User updated successfully")
}

// Delete deletes a user by ID
func (uc *UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := uc.UserService.Delete(id)
	if err != nil {
		common.Logger.Error("Error deleting user: ", err)
		uc.Response(c.Writer, nil, http.StatusInternalServerError, err.Error())
		return
	}
	uc.Response(c.Writer, nil, http.StatusOK, "User deleted successfully")
}
