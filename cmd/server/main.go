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

	listenerAddress := ":3000"

	dbStore, err := database.StoreInit(os.Getenv("DB_URL"))
	if err != nil {
		fmt.Println(err)
		return
	}

	handler := handlers.InitializeHandler(*dbStore)

	engine := gin.Default()
	routes.RouteSetup(*handler, engine)

	engine.Run(listenerAddress)
}
