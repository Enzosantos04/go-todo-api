package handlers

import (
	"gin-todo-listAPI/internal/services"
)



type taskHandler struct{
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *taskHandler{
	return &taskHandler{service:  service}
}

