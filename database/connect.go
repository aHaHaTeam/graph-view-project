package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Connect(databaseName string) (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	user := os.Getenv("DB1_USER")
	password := os.Getenv("DB1_PASSWORD")
	host := os.Getenv("DB1_HOST")
	port := os.Getenv("DB1_PORT")
	dbName := os.Getenv(databaseName)
	template := "postgres://%s:%s@%s:%s/%s"

	connStr := fmt.Sprintf(template, user, password, host, port, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
