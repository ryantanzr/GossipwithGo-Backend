package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

type Storage interface {
	CreateUser(*models.User) error
	DeleteUser(*models.User) error
	UpdateUser(*models.User) error
	GetUsersByName(string) ([]*models.User, error)

	CreatePost(*models.User, *models.Post) error
	DeletePost(*models.User, *models.Post) error
	UpdatePost(*models.Post) error
}

type PostgresStore struct {
	conn *pgxpool.Pool
}

func GetDatabaseStore() (*PostgresStore, error) {
	return &PostgresStore{}, nil
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

// Creates a user
func (pgs *PostgresStore) CreateUser(user *models.User) error {

	if err := user.EncryptData(); err != nil {
		return err
	}

	insertQuery := "INSERT INTO users VALUES (DEFAULT, $1, $2)"
	return pgs.NewTransaction(insertQuery, user.Username, user.Password)

}

// Get Account By Username
func (pgs *PostgresStore) GetAccountByUsername(user *models.User) pgx.Rows {

	query := "SELECT FROM users WHERE users.username = $1"
	rows, err := pgs.conn.Query(context.Background(), query, user.Username)

	if err != nil {
		fmt.Println("Failed to get account", err)
		return nil
	}

	return rows
}

// Delete a user
func (pgs *PostgresStore) DeleteUser(user *models.User) error {

	deleteQuery := "DELETE FROM users WHERE users.username = $1"
	return pgs.NewTransaction(deleteQuery, user.Username)

}

// Updates a user's details
func (pgs *PostgresStore) UpdateUser(user *models.User, newName string, newPw string) error {

	updateQuery := "UPDATE users SET username = $s1, password = $s2 WHERE username = $s3, password = $4"
	return pgs.NewTransaction(updateQuery, newName, newPw, user.Username, user.Password)
}

// Creates a new post
func (pgs *PostgresStore) CreatePost(user *models.User, post *models.Post) error {

	insertQuery := "INSERT INTO posts VALUES (DEFAULT, $1, $2, $3)"
	return pgs.NewTransaction(insertQuery, user.Username, post.Title, post.Content)

}

// Deletes a post
func (pgs *PostgresStore) DeletePost(user *models.User, post *models.Post) error {

	deleteQuery := "DELETE FROM posts WHERE posts.author = $1 AND posts.title = $2"
	return pgs.NewTransaction(deleteQuery, user.Username, post.Title)

}

// Updates a post
func (pgs *PostgresStore) UpdatePost(user *models.User, post *models.Post) error {

	updateQuery := "UPDATE post SET title = $s1, content = $s2 WHERE author = $s3, title = $4"
	return pgs.NewTransaction(updateQuery, post.Title, post.Content, user.Username, post.Title)
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

	return nil
}
