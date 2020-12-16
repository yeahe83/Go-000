// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"example.com/myapp/internal/biz"
	"example.com/myapp/internal/data"
	"fmt"
	"log"
)

// Injectors from main.go:

func initUserManager() (*biz.UserManager, error) {
	userDao, err := NewUserDao()
	if err != nil {
		return nil, err
	}
	userManager, err := NewUserManager(userDao)
	if err != nil {
		return nil, err
	}
	return userManager, nil
}

// main.go:

func main() {
	mgr, err := initUserManager()
	if err != nil {
		log.Fatalln(err.Error())
	}
	user, err := mgr.GetUser("1")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("user.Name = ", user.Name)
}

// provider
func NewUserManager(dao *data.UserDao) (*biz.UserManager, error) {
	mgr := biz.UserManager{Dao: dao}
	return &mgr, nil
}

func NewUserDao() (*data.UserDao, error) {
	dao := data.UserDao{}
	return &dao, nil
}