package routes

import (
	"github.com/gin-gonic/gin"
)

func routeSetup(r *gin.Engine) {

	//Registration requests
	registration := r.Group("/successfulRegistration")
	{
		registration.PUT("/successfulRegistration", handlers.successfulRegistration)
	}

	//Login requests
	login := r.Group("/login")
	{
		login.GET("/SuccessfulLogin", handlers.getUserByUsername)
	}
}
