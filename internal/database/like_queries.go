package database

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

// Creates a new like in the database
func (pgs *PostgresStore) CreateLike(action *models.UserAction) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		insertQuery := "INSERT INTO likes VALUES ($1, $2)"
		_, err := tx.Exec(context.Background(), insertQuery, action.ActorID, action.ReceiverID)
		return err
	})
}

// Deletes a like from the database
func (pgs *PostgresStore) DeleteLike(action *models.UserAction) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		deleteQuery := "DELETE FROM likes WHERE likes.userID = $1"
		_, err := tx.Exec(context.Background(), deleteQuery, action.ActorID)
		return err
	})
}

// Gets the number of likes for a post
func (pgs *PostgresStore) GetNumLikes(postID int) (int, error) {
	query := "SELECT COUNT(*) FROM likes WHERE likes.postID = $1"
	rows, err := pgs.conn.Query(context.Background(), query, postID)

	if err != nil {
		return 0, err
	}

	var numLikes int
	for rows.Next() {
		err = rows.Scan(&numLikes)
		if err != nil {
			return 0, err
		}
	}

	return numLikes, nil
}

// Gets if like from the actor to the receiver exists
func (pgs *PostgresStore) GetLike(action *models.UserAction) (bool, error) {
	query := "SELECT * FROM likes WHERE likes.userID = $1 AND likes.postID = $2"
	rows, err := pgs.conn.Query(context.Background(), query, action.ActorID, action.ReceiverID)

	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
