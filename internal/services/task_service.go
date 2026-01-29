package services

import (
	"gin-todo-listAPI/internal/repositories"
)


type TaskService struct{
	repo repositories.Task
}

func NewTaskHandler(repo repositories.Task) *TaskService{
	return &TaskService{repo: repo}
}
