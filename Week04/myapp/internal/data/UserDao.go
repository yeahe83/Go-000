package data

import (
	"errors"

	model "example.com/myapp/internal/model"
)

type UserDao struct {
}

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
