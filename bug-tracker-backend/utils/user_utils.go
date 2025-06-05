package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/models"
)

func LookupUserUsingID(id int) (*models.User, error) {
	var user models.User
	if err := conf.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func ExtractUserFromContext(c *gin.Context) models.User {
	contextUser, _ := c.Get("user")
	user, _ := contextUser.(models.User)

	return user
}
