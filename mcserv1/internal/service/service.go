package service

import "mcserv1/internal/models"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(name, email string) (string, error) {
	return "user create", nil

}

func (u *UserService) GetUser(name string) (models.User, error) {
	return models.User{
		Name:  name,
		Email: "@example",
	}, nil
}
