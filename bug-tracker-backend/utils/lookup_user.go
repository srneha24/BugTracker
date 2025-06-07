package utils

import (
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/models"
)

func LookupUserUsingID(id uint) (*models.User, error) {
	var user models.User
	if err := conf.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
