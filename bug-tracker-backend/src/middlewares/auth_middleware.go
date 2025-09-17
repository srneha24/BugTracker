package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/utils"
)

func RequireAuth(c *gin.Context) {
	// Extract the token from the Authorization header
	tokenString := strings.Split(c.GetHeader("Authorization"), " ")[1]

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
		c.Abort()
		return
	}

	// Validate the token using the JWT secret
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to authorize token"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token has expired"})
			c.Abort()
			return
		}

		subFloat, ok := claims["sub"].(float64)
		if !ok {
			log.Println("Invalid subject in token")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to authorize token"})
			c.Abort()
			return
		}

		user, _ := utils.LookupUserUsingID(uint(subFloat))
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			c.Abort()
			return
		}

		// Set the user in the context for further use
		c.Set("user", *user)

		// Proceed with the request
		c.Next()

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}
}
