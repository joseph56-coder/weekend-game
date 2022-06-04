package api

import "fmt"

type User struct {
	Username string `json:"username"`
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

type Users struct {
	list []*User
}

func (users Users) String() string {
	return fmt.Sprint(users.list)
}

func (users *Users) AddUser(u *User) {
	users.list = append(users.list, u)
}

func (users Users) FindByUsername(username string) (*User, bool) {
	for _, u := range users.list {
		if u.Username == username {
			return u, true
		}
	}
	return nil, false
}
