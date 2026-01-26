package routes

import (
	"gin-todo-listAPI/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	router.GET("/ping", handlers.Ping)
	router.POST("/user", userHandler.CreateUser)
	router.GET("/user", userHandler.GetAllUser)
	router.GET("/user/:id", userHandler.GetUserById)
}