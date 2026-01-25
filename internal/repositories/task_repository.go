package repositories

import "gin-todo-listAPI/internal/models"

type TaskRepository interface {
GetAllTasks() ([]models.Task, error)
GetTaskByID(id int) (*models.Task, error)
CreateTask(task *models.Task) error
UpdateTask(task *models.Task) error
DeleteTask(id int) error
}