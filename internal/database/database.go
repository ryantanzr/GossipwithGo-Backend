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
