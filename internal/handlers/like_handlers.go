package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new Like
func (h *Handler) CreateLike(ctx *gin.Context) {

	//Bind the json to the request object
	like, err := h.handleUserActionRequest(ctx)
	if err != nil {
		return
	}

	//Create a new like
	err = h.databaseStore.CreateLike(like)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusCreated, gin.H{"like": like})
}

// Delete a like
func (h *Handler) DeleteLike(ctx *gin.Context) {

	//Bind the request to a model
	like, err := h.handleUserActionRequest(ctx)
	if err != nil {
		return
	}

	err = h.databaseStore.DeleteLike(like)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"like": like})
}
