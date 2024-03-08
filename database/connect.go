package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"graph-view-project/server/models"
	"log"
	"os"
)

type PostgresDB struct {
	Connection *sql.DB
}

func (db *PostgresDB) Connect(databaseName string) error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	user := os.Getenv("DB1_USER")
	password := os.Getenv("DB1_PASSWORD")
	host := os.Getenv("DB1_HOST")
	port := os.Getenv("DB1_PORT")
	dbName := databaseName
	template := "postgres://%s:%s@%s:%s/%s"

	connStr := fmt.Sprintf(template, user, password, host, port, dbName)

	connection_, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	db.Connection = connection_
	return nil
}

func (db *PostgresDB) GetUserByLogin(login string) (*models.User, error) {
	query := `SELECT * FROM users WHERE login = $1`
	var email, password string
	err := db.Connection.QueryRow(query, login).Scan(email, password)
	if err != nil {
		return nil, err
	}
	return &models.User{Id: 0, Login: login, Email: email, Password: password}, nil
}

func (db *PostgresDB) CreateUser(login, email, password string) error {
	query := `INSERT INTO users (login, email, password) VALUES ($1, $2, $3)`
	err := db.Connection.QueryRow(query, login, email, password).Err()
	if err != nil {
		return err
	}
	return nil
}
