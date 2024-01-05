package api

import "github.com/ryantanzr/GossipwithGo-Backend/internal/database"

type APIServer struct {
	listenerAddress string
	storage         database.Storage
}
