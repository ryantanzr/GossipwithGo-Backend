package database

import (
	"context"

	pgx "github.com/jackc/pgx/v5"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

// Creates a new subscription in the database
func (pgs *PostgresStore) CreateSubscription(action *models.UserAction) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		insertQuery := "INSERT INTO subscriptions VALUES ($1, $2)"
		_, err := tx.Exec(context.Background(), insertQuery, action.ActorID, action.ReceiverID)
		return err
	})
}

// Delete a subscription
func (pgs *PostgresStore) DeleteSubscription(action *models.UserAction) error {
	return pgs.WithTransaction(context.Background(), func(tx pgx.Tx) error {
		deleteQuery := "DELETE FROM subscriptions WHERE subscriptions.subscriberID = $1 AND subscriptions.subscribeeID = $2"
		_, err := tx.Exec(context.Background(), deleteQuery, action.ActorID, action.ReceiverID)
		return err
	})
}

// Get if subscription from the actor to the receiver exists
func (pgs *PostgresStore) GetSubscription(action *models.UserAction) (bool, error) {
	query := "SELECT * FROM subscriptions WHERE subscriptions.subscriberID = $1 AND subscriptions.subscribeeID = $2"
	rows, err := pgs.conn.Query(context.Background(), query, action.ActorID, action.ReceiverID)

	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
