package repositories

import (
	"database/sql"
	"gin-todo-listAPI/internal/models"
)

type SqliteTaskRepository struct {
	DB *sql.DB
}

func NewSqliteTaskRepository(db *sql.DB) *SqliteTaskRepository{
	return &SqliteTaskRepository{DB: db}
}



func (r *SqliteTaskRepository) GetAllTasks() ([]models.Task, error) {
	rows, err := r.DB.Query("SELECT id, title, completed FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}



func (r *SqliteTaskRepository) CreateTask(task *models.Task) error {
	result, err := r.DB.Exec("INSERT INTO task (title, completed, user_id) VALUES (?, ?, ?)", 
	task.Title, task.Completed, task.User_id)

	if err != nil{
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = int(id)
	return nil


}