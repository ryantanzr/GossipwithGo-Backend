package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/database"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

// Utility function to process a json object for general account related activities
func handleAccountRequest(ctx *gin.Context) (*models.User, error) {

	var input models.AccountRequest

	//Bind the json to the request object
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	//Map it to the user model and return it
	return &models.User{
		Username: input.Username,
		Password: input.Password,
	}, nil
}

// POST a new user into the database
func Registration(ctx *gin.Context) {

	//Bind the json to the request object
	user, err := handleAccountRequest(ctx)
	if err != nil {
		return
	}

	//Get the database store
	pgs, err := database.GetDatabaseStore()
	if err != nil {
		return
	}

	//Create a new user
	err = pgs.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

// GET a user from postgres by their username
func Login(ctx *gin.Context) {

	//Bind the json to the request object
	user, err := handleAccountRequest(ctx)
	if err != nil {
		return
	}

	//Get the database
	pgs, err := database.GetDatabaseStore()
	if err != nil {
		return
	}

	//Query the database for the account and scan the data into the row
	row := pgs.GetAccountByUsername(user)
	user, err = models.ScanIntoUser(&row)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// DELETE a user
func DeleteAccount(ctx *gin.Context) {

	//Handle the account request
	user, err := handleAccountRequest(ctx)
	if err != nil {
		return
	}

	//Get the store
	pgs, err := database.GetDatabaseStore()
	if err != nil {
		return
	}

	//Delete the user
	err = pgs.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// PUT a new version of the user details into the database
func UpdateUserDetails(ctx *gin.Context) {

	var input models.UpdateRequest

	//Bind the json to the request object
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Map the details in the request accordingly
	oldUser := models.User{
		Username: input.Oldusername,
		Password: input.Oldpassword,
	}
	newUser := models.User{
		Username: input.Newusername,
		Password: input.Newpassword,
	}

	//Get the store
	pgs, err := database.GetDatabaseStore()
	if err != nil {
		return
	}

	//Update the user
	err = pgs.UpdateUser(&oldUser, input.Newusername, input.Newpassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"user": newUser})
}
