package databases

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile("internal/databases/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
