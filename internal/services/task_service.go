package services

import (
	"errors"
	"gin-todo-listAPI/internal/models"
	"gin-todo-listAPI/internal/repositories"
)


type TaskService struct{
	repo repositories.Task
}

func NewTaskService(repo repositories.Task) *TaskService{
	return &TaskService{repo: repo}
}



func (t *TaskService) CreateTask(task *models.Task) error{
	if task.Title == ""{
		return errors.New("task name should no be empty")
	}
	return t.repo.CreateTask(task);
}
