package models

type User struct {
	Id       int      `json:"id"`
	Login    string   `json:"login"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Graphs   []*Graph `json:"graphs"`
}

func NewUser(
	id int,
	login string,
	email string,
	password string,
	graphs []*Graph,
) *User {
	return &User{
		Id:       id,
		Login:    login,
		Email:    email,
		Password: password,
		Graphs:   graphs,
	}
}
