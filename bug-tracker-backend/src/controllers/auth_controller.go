package controllers

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	s "strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/models"
	types "github.com/WNBARookie/BugTracker/bug-tracker-backend/src/types"
)

func generateRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func createNewUser(newUser models.User) (*models.User, error) {
	result := conf.DB.Create(&newUser)

	if result.Error != nil {
		if s.Contains(result.Error.Error(), "uni_users_username") {
			newUser.Username = newUser.Username + "-" + generateRandomString()
			return createNewUser(newUser)
		}
		return nil, result.Error
	}

	return &newUser, nil
}

func SignUp(c *gin.Context) {
	var user types.SignUpUser

	if err := c.ShouldBindJSON(&user); err != nil {
		ec := conf.EnhancedContext{Context: c}
		ec.ValidationError(err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password"})
		return
	}

	username := user.Username
	if username == "" {
		username = s.Split(user.Email, "@")[0]
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Username: username,
		Password: string(hashedPassword),
	}
	createdUser, dbErr := createNewUser(newUser)

	if dbErr != nil {
		if s.Contains(dbErr.Error(), "uni_users_email") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		} else {
			log.Println("Error while creating user", dbErr.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create user"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": types.UserResponse{
			ID:        createdUser.ID,
			Name:      createdUser.Name,
			Username:  createdUser.Username,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		},
	})

}

func Login(c *gin.Context) {
	var user types.LoginUser
	if err := c.ShouldBindJSON(&user); err != nil {
		ec := conf.EnhancedContext{Context: c}
		ec.ValidationError(err.Error())
		return
	}

	var existingUser *models.User
	conf.DB.First(&existingUser, "email = ?", user.Email)
	if existingUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid password"})
		return
	}

	expiresInStr := os.Getenv("JWT_EXPIRES_IN")
	expiresIn, _ := strconv.Atoi(expiresInStr)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Duration(expiresIn) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("Failed to sign token:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type":  "Bearer",
		"token": tokenString,
	})
}
