package repositories

import (
	"database/sql"
	"errors"
	"gin-todo-listAPI/internal/models"
)

type SqliteUserRepository struct{
	DB *sql.DB
}

func NewSqliteUserRepository(db *sql.DB) *SqliteUserRepository{
	return &SqliteUserRepository{DB: db}

}


func (r SqliteUserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name) VALUES (?)"

	result, err := r.DB.Exec(query, user.Name)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}


func (r *SqliteUserRepository) GetAllUser() ([]models.User, error){
	query := "SELECT id, name FROM users"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
		
	}
	defer rows.Close()

	var users []models.User

	for rows.Next(){
		var user models.User 
		if err := rows.Scan(&user.ID, &user.Name); err != nil{
		return nil, err
	}
	users = append(users, user)
	}

	if err := rows.Err(); err != nil{
		return nil, err

	}



	return users, nil

	

}


func (r *SqliteUserRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name FROM users WHERE id = ?"

	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name)

	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil{
		return nil, err
	}

	return &user, nil
}

func (r *SqliteUserRepository) DeleteUserById(id string)  error{
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, id)

	return err
}


func(r *SqliteUserRepository) updateUserById(user *models.User) error{
	query := "UPDATE users SET name = ? WHERE id = ?"

	res, err := r.DB.Exec(query, user.Name, user.ID)

	if err != nil{
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil{
		return err
	}
	if rows == 0{
		return errors.New("User not found")
	}


	return nil 
}

