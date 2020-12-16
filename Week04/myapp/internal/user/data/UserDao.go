package data

import (
	"errors"

	model "example.com/myapp/internal/user/model"
)

// UserDao dao layout for user
type UserDao struct {
}

// GetUser return user by id
func (u *UserDao) GetUser(id string) (*model.User, error) {
	if id == "1" {
		user := model.User{
			Id:   "1",
			Name: "yeahe83",
			Age:  10,
		}
		return &user, nil
	}

	return nil, errors.New("NotFound")
}
