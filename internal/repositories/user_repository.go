package repositories

import "gin-todo-listAPI/internal/models"


type User interface{
	GetAllUser() ([]models.User, error)
	CreateUser( user *models.User) error
    GetUserById(string) (*models.User, error)
	DeleteUserById(string) error
	updateUserById(string, string) error
}