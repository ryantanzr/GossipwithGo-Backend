package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/api"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/handlers"
)

func RouteSetup(h handlers.Handler, r *gin.Engine) {

	// Account-related requests
	accounts := r.Group("/accounts")
	{
		accounts.POST("/", h.Registration)
		accounts.PUT("/", h.UpdateUserDetails)
		accounts.DELETE("/", h.DeleteAccount)
		accounts.POST("/login", api.WithJWTAuthorization(h.Login))
	}

	// Post-related requests
	posts := r.Group("/posts")
	{
		posts.POST("/", h.CreatePost)
		posts.PUT("/", h.UpdatePost)
		posts.DELETE("/", h.DeletePost)
	}

	// Like-related requests
	likes := r.Group("/likes")
	{
		likes.POST("/", h.CreateLike)
		likes.DELETE("/", h.DeleteLike)
	}

	// Subscription-related requests
	subscriptions := r.Group("/subscriptions")
	{
		subscriptions.POST("/", h.CreateSubscription)
		subscriptions.DELETE("/", h.DeleteSubscription)
	}
}
