package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

type Storage interface {
	CreateUser(*models.User) error
	DeleteUser(*models.User) error
	UpdateUser(*models.User) error
	GetUsersByName(string) ([]*models.User, error)
}

type PostgresStore struct {
	conn *pgxpool.Pool
}

func (pgs *PostgresStore) storeInit() (*PostgresStore, error) {

	emptyContext := context.Background()
	newConn, err := pgxpool.New(emptyContext, os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
		return nil, err
	}

	if err := newConn.Ping(emptyContext); err != nil {
		return nil, err
	}

	return &PostgresStore{
		conn: newConn,
	}, nil
}

func (pgs *PostgresStore) CreateUser(user *models.User) error {

	insertQuery := "INSERT INTO users VALUES (DEFAULT, $1, $2)"
	return pgs.NewTransaction(insertQuery, user.username, user.password)

}

func (pgs *PostgresStore) DeleteUser(user *models.User) error {

	deleteQuery := "DELETE FROM users WHERE users.username = $1"
	return pgs.NewTransaction(deleteQuery, user.username)
}

func (pgs *PostgresStore) UpdateUser(user *models.User, newName string, newPw string) error {

	updateQuery := "UPDATE users SET username = $s1, password = $s2 WHERE username = $s3, password = $4"
	return pgs.NewTransaction(updateQuery, newName, newPw, user.username, user.password)
}

func (pgs *PostgresStore) NewTransaction(query string, args ...string) error {

	//Begin the transaction (all-or-nothing)
	tx, err := pgs.conn.Begin(context.Background())
	if err != nil {
		fmt.Println("Beginning transaction", err)
		return err
	}
	//Is a no-op if the transaction completes
	defer tx.Rollback(context.Background())

	//Execute the transaction
	_, err = tx.Exec(context.Background(), query, args)
	if err != nil {
		fmt.Println("Executing transaction", err)
		return err
	}

	//Commit the transaction (mark it as completed)
	err = tx.Commit(context.Background())
	if err != nil {
		fmt.Println("Commit transaction", err)
		return err
	}
}
