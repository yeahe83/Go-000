package biz

import (
	data "example.com/myapp/internal/user/data"
	model "example.com/myapp/internal/user/model"
)

type UserManager struct {
	Dao *data.UserDao
}

func (u *UserManager) GetUser(id string) (*model.User, error) {
	return u.Dao.GetUser(id)
}
