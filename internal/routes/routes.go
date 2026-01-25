package routes

import (
	"gin-todo-listAPI/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", handlers.Ping)
}