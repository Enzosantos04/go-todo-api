package repositories

import "gin-todo-listAPI/internal/models"


type User interface{
	GetAllUser() ([]models.User, error)
	CreateUser( user *models.User) error
}