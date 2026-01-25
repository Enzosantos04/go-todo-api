package services

import (
	"errors"
	"gin-todo-listAPI/internal/models"
	"gin-todo-listAPI/internal/repositories"
)


type UserService struct{
	repo repositories.User
}

func NewUserService(repo repositories.User) *UserService{
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(user *models.User) error{
if(user.Name) ==""{
	return errors.New("Name field should not be empty")
}
return u.repo.CreateUser(user);
}

