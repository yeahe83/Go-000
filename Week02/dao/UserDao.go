package dao

import (
	"database/sql"
	"fmt"

	"../model"
)

// GetOne get user by id
func GetOne(id int) (user model.User, err error) {

	var db *sql.DB
	user = model.User{}

	db, err = sql.Open("sqlserver", "...") // TODO
	if err != nil {
		return user, fmt.Errorf("Connect Fail: %w", err) // wrap
	}

	err = db.QueryRow("select ID,Name from Users Where ID=@ID", sql.Named("ID", id)).Scan(&user.ID, &user.Name)
	if err != nil {
		return user, fmt.Errorf("GetOne Fail: %w", err) // wrap
	}
	return
}
