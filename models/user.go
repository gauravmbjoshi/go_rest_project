package models

import (
	"errors"

	"example.com/go_rest_api_backend_project/db"
	"example.com/go_rest_api_backend_project/utils"
)

type User struct {
	ID           int64
	Email        string `Binding:"required"`
	Password     string `Binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword,err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err == nil {
		u.ID = id
	}
	return err
}
func (u *User) ValidateCredentials() error {

	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}

	isPasswordValid := utils.CheckPassword(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("credentials invalid")
	}
	return nil
}