package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

// POST a new user into the database
func (h *Handler) Registration(ctx *gin.Context) {

	//Bind the json to the request object
	user, err := h.handleAccountRequest(ctx)
	if err != nil {
		return
	}

	//Create a new user
	err = h.databaseStore.CreateUser(user)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

// GET a user from postgres by their username
func (h *Handler) Login(ctx *gin.Context) {

	//Bind the json to the request object
	user, err := h.handleAccountRequest(ctx)
	if err != nil {
		return
	}

	//Query the database for the account and scan the data into the row
	row := h.databaseStore.GetUserByUsername(user)
	user, err = models.ScanIntoUser(&row)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// DELETE a user
func (h *Handler) DeleteAccount(ctx *gin.Context) {

	//Handle the account request
	user, err := h.handleAccountRequest(ctx)
	if err != nil {
		return
	}

	//Delete the user
	err = h.databaseStore.DeleteUser(user)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// PUT a new version of the user details into the database
func (h *Handler) UpdateUserDetails(ctx *gin.Context) {

	var input models.UpdateRequest

	//Bind the json to the request object
	if bindJSON(ctx, &input) != nil {
		return
	}

	newUser := models.User{
		ID:       input.ID,
		Username: input.Newusername,
		Password: input.Newpassword,
	}

	err := h.databaseStore.UpdateUser(&newUser)
	handleError(ctx, err, http.StatusBadRequest)
	if err != nil {
		return
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"user": newUser})
}
