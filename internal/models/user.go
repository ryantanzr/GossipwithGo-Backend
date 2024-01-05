package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	username string `json:"name"`
	password string `json:"password"`
}

func (user *User) encryptData() error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.password = string(hash)
	user.username = html.EscapeString(strings.TrimSpace(user.username))
	return nil
}

func (user *User) validatePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.password), []byte(pw))
}
