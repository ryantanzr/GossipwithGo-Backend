package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST a new post into the database
func (h *Handler) CreatePost(ctx *gin.Context) {

	//Bind the json to the request object
	post, err := h.handlePostRequest(ctx)
	if err != nil {
		return
	}

	//Create a new post
	err = h.databaseStore.CreatePost(post)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusCreated, gin.H{"user": post})
}

// PUT a new version of an updated post
func (h *Handler) UpdatePost(ctx *gin.Context) {

	//Bind the request to a model
	post, err := h.handlePostRequest(ctx)
	if err != nil {
		return
	}

	err = h.databaseStore.UpdatePost(post)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"post": post})
}

// PUT a new version of an updated post
func (h *Handler) DeletePost(ctx *gin.Context) {

	//Handle the post request
	post, err := h.handlePostRequest(ctx)
	if err != nil {
		return
	}

	//Delete the post
	err = h.databaseStore.DeletePost(post)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"post": post})
}
