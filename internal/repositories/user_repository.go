package repositories

import "gin-todo-listAPI/internal/models"


type User interface {
	GetAllUser() ([]models.User, error)
	CreateUser(user *models.User) error
	GetUserById(id string) (*models.User, error)
	DeleteUserById(id string) error
	UpdateUserById(id string, user *models.User) error
}