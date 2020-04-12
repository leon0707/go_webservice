package models

import "errors"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("error")
	}
	u.ID = nextID
	nextID++

	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, errors.New("not found")
}

func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1])
			return nil
		}
	}
	return errors.New("not found")
}
