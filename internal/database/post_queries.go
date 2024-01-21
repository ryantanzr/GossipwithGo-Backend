package database

import (
	"context"

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
