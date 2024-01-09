package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/handlers"
)

func routeSetup(r *gin.Engine) {

	//Registration requests
	registration := r.Group("/reg")
	{
		registration.PUT("/success", handlers.Registeration)
	}

	//Login requests
	login := r.Group("/login")
	{
		login.GET("/success", handlers.Login)
	}
}
