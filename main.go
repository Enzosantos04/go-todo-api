package main

import (
	"gin-todo-listAPI/internal/databases"
	"gin-todo-listAPI/internal/handlers"
	"gin-todo-listAPI/internal/repositories"
	"gin-todo-listAPI/internal/routes"
	"gin-todo-listAPI/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := databases.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repositories.NewSqliteUserRepository(db)
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)


	router := gin.Default()
	routes.RegisterRoutes(router, userHandler)
	router.Run(":8080")
}

