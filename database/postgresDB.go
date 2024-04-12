package database

import (
	"database/sql"
	"fmt"
	"graph-view-project/models"
	"log"
	"os"

	"github.com/joho/godotenv"
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

func (db *PostgresDB) Disconnect() error {
	err := db.Connection.Close()
	return err
}

func (db *PostgresDB) CreateUser(user models.User) (*models.User, error) {
	query := `INSERT INTO users (login, email, password) VALUES ($1, $2, $3)`
	err := db.Connection.QueryRow(query, user.Login, user.Email, user.Password).Err()
	if err != nil {
		return nil, err
	}
	return nil, nil
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

func (db *PostgresDB) GetUser(id int) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`
	var login, email, password string
	err := db.Connection.QueryRow(query, id).Scan(login, email, password)
	if err != nil {
		return nil, err
	}
	return &models.User{Id: id, Login: login, Email: email, Password: password}, nil
}

func (db *PostgresDB) CreateGraph(user models.User, graph models.Graph) (*models.Graph, error) {
	return nil, nil
}
func (db *PostgresDB) CreateNode(graph models.Graph, node models.Node) (*models.Node, error) {
	return nil, nil
}
func (db *PostgresDB) CreateEdge(graph models.Graph, edge models.Edge) (*models.Edge, error) {
	return nil, nil
}

func (db *PostgresDB) GetGraph(id int) (*models.Graph, error) {
	return nil, error(nil)
}
func (db *PostgresDB) GetEdge(id int) (*models.Edge, error) {
	return nil, error(nil)
}
func (db *PostgresDB) GetNode(id int) (*models.Node, error) {
	return nil, error(nil)
}

func (db *PostgresDB) UpdateUser(newUser models.User) error {
	return error(nil)
}

func (db *PostgresDB) UpdateGraph(newGraph models.Graph) error {
	return error(nil)
}

func (db *PostgresDB) UpdateEdge(newEdge models.Edge) error {
	return error(nil)
}

func (db *PostgresDB) UpdateNode(newNode models.Node) error {
	return error(nil)
}
