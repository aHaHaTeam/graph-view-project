package database

import (
	"errors"
	"graph-view-project/server/models"
	"graph-view-project/server/utils"
)

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

func (db *MockDB) Connect(databaseName string) error {
	db.users = make(map[string]models.User)
	hash, _ := utils.GenerateHashPassword("password")
	_ = db.CreateUser("user", "user@user", hash)
	hash, _ = utils.GenerateHashPassword("12345678password")
	_ = db.CreateUser("12345", "123456", hash)
	return nil
}
