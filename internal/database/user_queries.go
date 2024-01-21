package database

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

// Creates a user
func (pgs *PostgresStore) CreateUser(user *models.User) error {

	if err := user.EncryptData(); err != nil {
		return err
	}

	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		insertQuery := "INSERT INTO users VALUES (DEFAULT, $1, $2)"
		_, err := tx.Exec(context.Background(), insertQuery, user.Username, user.Password)
		return err
	})
}

// Get Account By Username
func (pgs *PostgresStore) GetUserByUsername(user *models.User) pgx.Rows {

	query := "SELECT * FROM users WHERE users.username = $1"
	rows, err := pgs.conn.Query(context.Background(), query, user.Username)

	if err != nil {
		fmt.Println("Failed to get account", err)
		return nil
	}

	return rows
}

// Get Account By Username
func (pgs *PostgresStore) GetUserByID(id int) pgx.Rows {

	query := `SELECT * FROM users WHERE "userID" = $1`
	rows, err := pgs.conn.Query(context.Background(), query, id)

	if err != nil {
		fmt.Println("Failed to get account", err)
		return nil
	}

	return rows
}

// Delete a user
func (pgs *PostgresStore) DeleteUser(user *models.User) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		deleteQuery := `DELETE FROM users WHERE users.username = $1`
		_, err := tx.Exec(context.Background(), deleteQuery, user.Username)
		return err
	})
}

// Updates a user's details
func (pgs *PostgresStore) UpdateUser(newUser *models.User) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		updateQuery := `UPDATE users SET username = $1, password = $2 WHERE "userID" = $3`
		_, err := tx.Exec(context.Background(), updateQuery, newUser.Username, newUser.Password, newUser.ID)
		return err
	})
}
