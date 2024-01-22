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

func InitializeHandler(dbs database.PostgresStore) *Handler {
	return &Handler{databaseStore: dbs}
}

// Utility function to handle errors
func handleError(ctx *gin.Context, err error, status int) {
	if err != nil {
		ctx.JSON(status, gin.H{"error": err.Error()})
		ctx.Abort() // prevent calling further handlers
	}
}

// Utility function to bind a json object to a struct
func bindJSON(ctx *gin.Context, obj interface{}) error {
	if err := ctx.BindJSON(obj); err != nil {
		handleError(ctx, err, http.StatusBadRequest)
		return err
	}
	return nil
}

// Utility function to process a json object for general account related activities
func (h *Handler) handleAccountRequest(ctx *gin.Context) (*models.User, error) {

	var input models.AccountRequest

	//Bind the json to the request object
	if err := bindJSON(ctx, &input); err != nil {
		return nil, err
	}

	//Map it to the user model and return it
	return &models.User{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
	}, nil
}

// Utility function to process a json object for general account related activities
func (h *Handler) handlePostRequest(ctx *gin.Context) (*models.Post, error) {

	var input models.PostRequest

	//Bind the json to the request object
	if err := bindJSON(ctx, &input); err != nil {
		return nil, err
	}

	//Map it to the user model and return it
	return &models.Post{
		ID:      input.ID,
		Author:  input.Author,
		Title:   input.Title,
		Content: input.Content,
		Likes:   input.Likes,
	}, nil
}

// Utility function to process a json object for general user action related activities
func (h *Handler) handleUserActionRequest(ctx *gin.Context) (*models.UserAction, error) {

	var input models.UserActionRequest

	//Bind the json to the request object
	if err := bindJSON(ctx, &input); err != nil {
		return nil, err
	}

	//Map it to the user model and return it
	return &models.UserAction{
		ActorID:    input.ActorID,
		ReceiverID: input.ReceiverID,
		ActionType: input.ActionType,
	}, nil
}
