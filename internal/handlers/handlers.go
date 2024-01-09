package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/database"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

// POST a new user into the database
func Registeration(c *gin.Context) {

	var input models.AccountRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	pgs, err := database.GetDatabaseStore()
	if err != nil {
		return
	}

	err = pgs.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})

}

// GET a user from postgres by their username
func Login(c *gin.Context) {

	var input models.AccountRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	pgs, err := database.GetDatabaseStore()
	if err != nil {
		return
	}

	row := pgs.GetAccountByUsername(&user)
	user, err = models.ScanIntoUser(&row)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})

}

// DELETE a user
func DeleteAccount(c *gin.Context) {

	var input models.AccountRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	pgs, err := database.GetDatabaseStore()
	if err != nil {
		return
	}

	err = pgs.DeleteUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
