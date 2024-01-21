package database

import (
	"context"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

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

// Gets first ten posts by recent published dates
func (pgs *PostgresStore) GetSomePostsByRecent(num int) pgx.Rows {
	query := "SELECT * FROM posts ORDER BY posts.published DESC LIMIT $1"
	rows, err := pgs.conn.Query(context.Background(), query, num)

	if err != nil {
		fmt.Println("Failed to get some posts", err)
		return nil
	}

	return rows
}

// Gets a post by author, sorted by recent published dates
func (pgs *PostgresStore) GetPostByAuthor(author string) pgx.Rows {
	query := "SELECT * FROM posts WHERE posts.author = $1 ORDER BY posts.published DESC"
	rows, err := pgs.conn.Query(context.Background(), query, author)

	if err != nil {
		fmt.Println("Failed to get post by author", err)
		return nil
	}

	return rows
}
