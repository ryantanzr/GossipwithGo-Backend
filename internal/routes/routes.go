package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/api"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/handlers"
)

func RouteSetup(h handlers.Handler, r *gin.Engine) {

	//Registration requests
	registration := r.Group("/reg")
	{
		registration.POST("/", h.Registration)
	}

	//Login requests
	login := r.Group("/login")
	{
		login.GET("/", api.WithJWTAuthorization(h.Login))
	}

	//Update requests
	update := r.Group("/update")
	{
		update.DELETE("/", h.DeleteAccount)
		update.PUT("/", h.UpdateUserDetails)
	}
}
