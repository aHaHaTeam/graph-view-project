package database

import (
	"errors"
	"graph-view-project/server/models"
	"graph-view-project/server/utils"
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
	hash, _ := utils.GenerateHashPassword("password")
	_ = DB.CreateUser("user", "user@user", hash)
	hash, _ = utils.GenerateHashPassword("12345678password")
	_ = DB.CreateUser("12345", "123456", hash)
	return nil
}
