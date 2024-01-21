package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/ryantanzr/GossipwithGo-Backend/internal/models"
)

const VALID_DURATION = 24

// JWT Middleware to be used together with a login function
func WithJWTAuthorization(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fmt.Println("JWT auth middleware")

		//Get the tokenString
		tokenString := ctx.Request.Header.Get("Auth")
		//Validate the token
		token, err := validateJWT(tokenString)
		if !token.Valid || err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"Access Denied": err})
			return
		}

		//Assert the claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid token claims"})
			return
		}

		// Add claims to context
		ctx.Set("claims", claims)

		//Call the handler function
		handlerFunc(ctx)

	}
}

func createJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * VALID_DURATION)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        string(user.ID),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(os.Getenv("JWT_SECRET"))
}

func validateJWT(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		//Checks the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		//Return the secret (change this when deploying)
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
