package database

import "graph-view-project/server/models"

type DataBase interface {
	GetUserByLogin(login string) (models.User, error)
	CreateUser(login, email, password string) error
	Connect(databaseName string) error
}
