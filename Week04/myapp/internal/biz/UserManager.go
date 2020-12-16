package biz

import (
	data "example.com/myapp/internal/data"
	model "example.com/myapp/internal/model"
)

type UserManager struct {
	Dao *data.UserDao
}

func (u *UserManager) GetUser(id string) (*model.User, error) {
	return u.Dao.GetUser(id)
}
