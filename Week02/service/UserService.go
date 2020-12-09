package service

import (
	"fmt"

	"../dao"
	"../model"
)

// GetOne get user by id
func GetOne(id int) (user model.User, err error) {

	user, err = dao.GetOne(id) // 直接调dao了
	if err != nil {
		return user, fmt.Errorf("GetOne in service Fail: %w", err) // wrap
	}
	return user, err
}
