package main

import (
	"gin-todo-listAPI/internal/databases"
	"gin-todo-listAPI/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := databases.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}

