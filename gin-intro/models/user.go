package models

import (
	"errors"
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO users(email, password) 
	VALUES ($1, $2)
	RETURNING id
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var id int64
	err = stmt.QueryRow(user.Email, hashPassword).Scan(&id)
	defer stmt.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	user.ID = id
	return nil
}

func (user *User) ValidateCredentials() error {
	query := `
	SELECT password FROM users WHERE email = $1
	`

	var retrievedPassword string
	err := db.DB.QueryRow(query, user.Email).Scan(&retrievedPassword)

	if err != nil {
		fmt.Println(err)
		return err
	}

	passwordIsValid := utils.CheckPassword(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credential Invalid")
	}

	return nil
}
