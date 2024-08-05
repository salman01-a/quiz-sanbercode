package services

import (
	"book-management/utils"
	"database/sql"
	"errors"
)

func Login(username, password string) (string, error) {
	var dbUser, dbPass string
	query := "SELECT username, password FROM users WHERE username=$1"
	err := utils.DB.QueryRow(query, username).Scan(&dbUser, &dbPass)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if dbPass != password {
		return "", errors.New("invalid password")
	}

	// Here you would generate a token for the user
	token := "dummy_token" // Replace with actual token generation logic

	return token, nil
}
