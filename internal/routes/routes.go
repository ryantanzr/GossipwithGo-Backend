package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/api"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/handlers"
)

func RouteSetup(h handlers.Handler, r *gin.Engine) {

	//Account-related requests
	registration := r.Group("/acc")
	{
		registration.POST("/", h.Registration)
		registration.PUT("/", h.UpdateUserDetails)
		registration.DELETE("/", h.DeleteAccount)

	}

	//Login requests
	login := r.Group("/login")
	{
		login.GET("/", api.WithJWTAuthorization(h.Login))
	}

	//Post-related requests
	posts := r.Group("/post")
	{
		posts.POST("/", h.Post)
		posts.PUT("/", h.UpdatePost)
		posts.DELETE("/", h.DeletePost)
	}
}
