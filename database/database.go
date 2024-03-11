package database

import "graph-view-project/models"

type DataBase interface {
	Connect(databaseName string) error
	GetUserByLogin(login string) (models.User, error)
	CreateUser(login, email, password string) error
}
