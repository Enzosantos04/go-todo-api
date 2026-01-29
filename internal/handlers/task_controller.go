package handlers

import (
	"gin-todo-listAPI/internal/models"
	"gin-todo-listAPI/internal/services"

	"github.com/gin-gonic/gin"
)



type TaskHandler struct{
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler{
	return &TaskHandler{service:  service}
}


func (h *TaskHandler) CreateTask(c *gin.Context){
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil{
			c.JSON(400, gin.H{"erorr": "Invalid body"})
		return
	}

	err := h.service.CreateTask(&task)
	if err != nil{
	c.JSON(400, gin.H{"error": err.Error()})
	return
	}

	c.JSON(200, task)
}

