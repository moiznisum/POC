package user

import (
	"errors"
	"ai-saas-schedular-server/server/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Get() ([]models.User, error) {
	users, err := models.FindAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) Create(userData map[string]interface{}) (*models.User, error) {
	user, err := models.CreateUser(userData)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetByID(id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("invalid payload")
	}

	user, err := models.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) Update(id string, userData map[string]interface{}) (*models.User, error) {
	if id == "" || userData == nil {
		return nil, errors.New("invalid payload")
	}

	user, err := models.UpdateUserByID(id, userData)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) Delete(id string) error {
	if id == "" {
		return errors.New("invalid payload")
	}

	err := models.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
