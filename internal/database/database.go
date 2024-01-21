package database

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx/v5"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

type Storage interface {
	CreateUser(*models.User) error
	DeleteUser(*models.User) error
	UpdateUser(*models.User) error
	GetUserByName(string) (*models.User, error)

	CreatePost(*models.User, *models.Post) error
	DeletePost(*models.User, *models.Post) error
	UpdatePost(*models.Post) error
}

// PostgresStore is a struct that holds the connection to the database
type PostgresStore struct {
	conn *pgxpool.Pool
}

// TransactionFunc is a function that takes a transaction and returns an error.
type TransactionFunc func(tx pgx.Tx) error

// Getter function, sorry...
func GetDatabaseStore() (*PostgresStore, error) {
	return &PostgresStore{}, nil
}

// Initializes the database
func StoreInit(dbConnectionString string) (*PostgresStore, error) {

	emptyContext := context.Background()
	newConn, err := pgxpool.New(emptyContext, os.Getenv("DATABASE_URL"))

	fmt.Println(dbConnectionString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
		return nil, err
	}

	if err := newConn.Ping(emptyContext); err == nil {
		fmt.Println("successful ping")
	}

	return &PostgresStore{
		conn: newConn,
	}, nil
}

// A helper function that handles the boilerplate code for starting and committing a transaction.
func (pgs *PostgresStore) WithTransaction(ctx context.Context, tf TransactionFunc) (err error) {
	tx, err := pgs.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = tf(tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

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
func (pgs *PostgresStore) GetAccountByUsername(user *models.User) pgx.Rows {

	query := "SELECT * FROM users WHERE users.username = $1"
	rows, err := pgs.conn.Query(context.Background(), query, user.Username)

	if err != nil {
		fmt.Println("Failed to get account", err)
		return nil
	}

	return rows
}

// Get Account By Username
func (pgs *PostgresStore) GetAccountByID(id int) pgx.Rows {

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

// Creates a new post
func (pgs *PostgresStore) CreatePost(post *models.Post) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		insertQuery := "INSERT INTO posts VALUES (DEFAULT, $1, $2, $3)"
		_, err := tx.Exec(context.Background(), insertQuery, post.Author, post.Title, post.Content)
		return err
	})
}

// Deletes a post
func (pgs *PostgresStore) DeletePost(post *models.Post) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		deleteQuery := "DELETE FROM posts WHERE posts.author = $1 AND posts.title = $2"
		_, err := tx.Exec(context.Background(), deleteQuery, post.Author, post.Title)
		return err
	})
}

// Updates a post
func (pgs *PostgresStore) UpdatePost(post *models.Post) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		updateQuery := "UPDATE post SET title = $s1, content = $s2 WHERE author = $s3, title = $4"
		_, err := tx.Exec(context.Background(), updateQuery, post.Title, post.Content, post.Author, post.Title)
		return err
	})

}
