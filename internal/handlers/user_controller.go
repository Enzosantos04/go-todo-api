package handlers

import (
	"gin-todo-listAPI/internal/models"
	"gin-todo-listAPI/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(400, gin.H{"erorr": "Invalid body"})
		return
	}

	err := h.service.CreateUser(&user)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, user)
	

}


func (h *UserHandler) GetAllUser(c *gin.Context){

	users, err := h.service.GetAllUser()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
    return
	}

	c.JSON(200, users)
}


func (h *UserHandler) GetUserById(c *gin.Context) {
    id := c.Param("id")

    user, err := h.service.GetUserById(id)
    if err != nil {
        c.JSON(400, gin.H{"error": "User not found", "id": id, "details": err.Error()})
        return
    }

    c.JSON(200, user)
}


func (h *UserHandler) DeleteUserById(c *gin.Context){
	id := c.Param("id")

	 err := h.service.DeleteUserById(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found", "id": id, "details": err.Error()})
		
	}

	c.JSON(200, http.StatusOK)
}


func (h *UserHandler) UpdateUserById(c *gin.Context){
	id := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(400, gin.H{"error": "Invalid body"})
		return
	}

	err := h.service.UpdateUserById(id, &user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}