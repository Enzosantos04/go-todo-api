package repositories

import (
	"database/sql"
	"gin-todo-listAPI/internal/models"
)

type SqliteUserRepository struct{
	DB *sql.DB
}

func NewSqliteUserRepository(db *sql.DB) *SqliteUserRepository{
	return &SqliteUserRepository{DB: db}

}


func (r SqliteUserRepository) CreateUser(user *models.User) error{
	query := "INSERT INTO users (name) VALUES(?) "
	_, err := r.DB.Exec(query, user.Name)
	return err
}

