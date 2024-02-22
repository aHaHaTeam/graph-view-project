package database

import (
	"errors"
	"graph-view-project/server/models"
)

var DB DataBase

type DataBase interface {
	GetUserByLogin(login string) (models.User, error)
	CreateUser(login, email, password string) error
	Connect() error
}

type MockDB struct {
	users map[string]models.User
}

func (db *MockDB) GetUserByLogin(login string) (models.User, error) {
	user, ok := db.users[login]
	if !ok {
		return models.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (db *MockDB) CreateUser(login, email, password string) error {
	_, ok := db.users[login]
	if ok {
		return errors.New("user already exists")
	}

	db.users[login] = models.User{Id: 42, Login: login, Email: email, Password: password}

	return nil
}

func (db *MockDB) Connect() error {
	DB = db
	db.users = make(map[string]models.User)
	_ = DB.CreateUser("user", "user@user", "password")
	_ = DB.CreateUser("123", "123", "123")
	return nil
}
