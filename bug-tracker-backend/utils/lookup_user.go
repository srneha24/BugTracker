package utils

import (
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/models"
)

func LookupUser(email string) (*models.User, error) {
	var user models.User
	if err := conf.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
