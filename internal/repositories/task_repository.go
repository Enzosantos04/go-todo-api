package repositories

import "gin-todo-listAPI/internal/models"

type Task interface {
GetAllTasks() ([]models.Task, error)
GetTaskByID(id string) (*models.Task, error)
CreateTask(task *models.Task) error
UpdateTaskById(id string, task *models.Task)  error 
DeleteTaskById(id string) error;
}