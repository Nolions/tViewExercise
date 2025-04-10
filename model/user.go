package model

type User struct {
	Id       int
	Name     string
	Username string
	Password string
	Email    string
	Enable   bool
	Note     string
}

func NewUser() User {
	return User{
		Id:       1,
		Name:     "Sam",
		Username: "sam78",
		Password: "secret",
		Email:    "sam78@cc.com",
		Enable:   true,
		Note:     "test tset test\ntest",
	}
}
