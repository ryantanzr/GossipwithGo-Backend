package models

import (
	"fmt"
	"html"
	"strings"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ScanIntoUser(rows *pgx.Rows) (User, error) {

	user := User{0, "", ""}
	if err := pgxscan.ScanRow(&user, *rows); err != nil {
		fmt.Println("Scan row error", err)
		return user, err
	}

	return user, nil

}

// Encrypts a password with bcrypt algorithm to produce a hash
func (user *User) EncryptData() error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

// Checks against the passwords
func (user *User) ValidatePassword(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pw))
}
