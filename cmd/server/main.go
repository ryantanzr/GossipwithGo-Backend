package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/database"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/handlers"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/routes"
)

func main() {

	os.Setenv("DATABASE_URL", "host=localhost port=5432 dbname=GossipWithGo user=postgres sslmode=prefer connect_timeout=10")
	listenerAddress := ":3000"

	dbStore, err := database.StoreInit(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		return
	}

	handler := handlers.InitializeHandler(*dbStore)

	engine := gin.Default()
	routes.RouteSetup(handler, engine)

	engine.Run(listenerAddress)
}
