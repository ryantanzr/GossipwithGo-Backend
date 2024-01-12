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

type PostgresStore struct {
	conn *pgxpool.Pool
}

func GetDatabaseStore() (*PostgresStore, error) {
	return &PostgresStore{}, nil
}

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

// Creates a user
func (pgs *PostgresStore) CreateUser(user *models.User) error {

	if err := user.EncryptData(); err != nil {
		return err
	}

	//Begin the transaction (all-or-nothing)
	tx, err := pgs.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	//Is a no-op if the transaction completes
	defer tx.Rollback(context.Background())

	//Execute the transaction
	insertQuery := "INSERT INTO users VALUES (DEFAULT, $1, $2)"
	_, err = tx.Exec(context.Background(), insertQuery, user.Username, user.Password)
	if err != nil {
		return err
	}

	//Commit the transaction (mark it as completed)
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
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

	//Begin the transaction (all-or-nothing)
	tx, err := pgs.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	//Is a no-op if the transaction completes
	defer tx.Rollback(context.Background())

	//Execute the transaction
	deleteQuery := `DELETE FROM users WHERE users.username = $1`
	_, err = tx.Exec(context.Background(), deleteQuery, user.Username)
	if err != nil {
		return err
	}

	//Commit the transaction (mark it as completed)
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// Updates a user's details
func (pgs *PostgresStore) UpdateUser(newUser *models.User) error {

	//Begin the transaction (all-or-nothing)
	tx, err := pgs.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	//Is a no-op if the transaction completes
	defer tx.Rollback(context.Background())

	//Execute the transaction
	updateQuery := `UPDATE users SET username = $1, password = $2 WHERE "userID" = $3`
	_, err = tx.Exec(context.Background(), updateQuery, newUser.Username, newUser.Password, newUser.ID)
	if err != nil {
		return err
	}

	//Commit the transaction (mark it as completed)
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// Creates a new post
func (pgs *PostgresStore) CreatePost(user *models.User, post *models.Post) error {

	//Begin the transaction (all-or-nothing)
	tx, err := pgs.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	//Is a no-op if the transaction completes
	defer tx.Rollback(context.Background())

	//Execute the transaction
	insertQuery := "INSERT INTO posts VALUES (DEFAULT, $1, $2, $3)"
	_, err = tx.Exec(context.Background(), insertQuery, user.Username, post.Title, post.Content)
	if err != nil {
		return err
	}

	//Commit the transaction (mark it as completed)
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil

}

// Deletes a post
func (pgs *PostgresStore) DeletePost(user *models.User, post *models.Post) error {

	//Begin the transaction (all-or-nothing)
	tx, err := pgs.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	//Is a no-op if the transaction completes
	defer tx.Rollback(context.Background())

	//Execute the transaction
	deleteQuery := "DELETE FROM posts WHERE posts.author = $1 AND posts.title = $2"
	_, err = tx.Exec(context.Background(), deleteQuery, user.Username, post.Title)
	if err != nil {
		return err
	}

	//Commit the transaction (mark it as completed)
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil

}

// Updates a post
func (pgs *PostgresStore) UpdatePost(user *models.User, post *models.Post) error {

	//Begin the transaction (all-or-nothing)
	tx, err := pgs.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	//Is a no-op if the transaction completes
	defer tx.Rollback(context.Background())

	//Execute the transaction
	updateQuery := "UPDATE post SET title = $s1, content = $s2 WHERE author = $s3, title = $4"
	_, err = tx.Exec(context.Background(), updateQuery, post.Title, post.Content, user.Username, post.Title)
	if err != nil {
		return err
	}

	//Commit the transaction (mark it as completed)
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
