package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/database"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

type Handler struct {
	databaseStore database.PostgresStore
}

func InitializeHandler(dbs database.PostgresStore) Handler {
	return Handler{databaseStore: dbs}
}

// Utility function to process a json object for general account related activities
func (h *Handler) handleAccountRequest(ctx *gin.Context) (*models.User, error) {

	var input models.AccountRequest

	//Bind the json to the request object
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"accError": err.Error()})
		return nil, err
	}

	//Map it to the user model and return it
	return &models.User{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
	}, nil
}

// POST a new user into the database
func (h *Handler) Registration(ctx *gin.Context) {

	//Bind the json to the request object
	user, err := h.handleAccountRequest(ctx)
	if err != nil {
		return
	}

	//Create a new user
	err = h.databaseStore.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	row := h.databaseStore.GetAccountByUsername(user)
	user, err = models.ScanIntoUser(&row)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// PUT a new version of the user details into the database
func (h *Handler) UpdateUserDetails(ctx *gin.Context) {

	var input models.UpdateRequest

	//Bind the json to the request object
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{
		ID:       input.ID,
		Username: input.Newusername,
		Password: input.Newpassword,
	}

	err := h.databaseStore.UpdateUser(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Return a json to indicate the operation was successful
	ctx.JSON(http.StatusOK, gin.H{"user": newUser})
}
