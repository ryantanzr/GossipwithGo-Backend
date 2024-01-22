package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new Subscription
func (h *Handler) CreateSubscription(ctx *gin.Context) {

	//Bind the json to the request object
	subscription, err := h.handleUserActionRequest(ctx)
	if err != nil {
		return
	}

	//Create a new subscription
	err = h.databaseStore.CreateSubscription(subscription)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusCreated, gin.H{"subscription": subscription})
}

// Delete a subscription
func (h *Handler) DeleteSubscription(ctx *gin.Context) {

	//Bind the request to a model
	subscription, err := h.handleUserActionRequest(ctx)
	if err != nil {
		return
	}

	err = h.databaseStore.DeleteSubscription(subscription)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"subscription": subscription})
}
