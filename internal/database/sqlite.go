package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)



func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./todo.db")
content, err := os.ReadFile("internal/schema.sql")
if err != nil {
	log.Fatal(err)
	
}

_, err = db.Exec(string(content))
if err != nil {
	log.Fatal(err)
	
}
	return db, err
}


