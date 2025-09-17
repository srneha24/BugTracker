package utils

import (
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/models"
)

func LookupUserUsingID(id uint) (*models.User, error) {
	var user models.User
	if err := conf.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckIfUserIsProjectMember(userID, projectID uint) (string, error) {
	var team models.Team
	if err := conf.DB.Where("user_id = ? AND project_id = ?", userID, projectID).First(&team).Error; err != nil {
		if err.Error() == "record not found" {
			return "", nil // User is not a member of the project
		}
		return "", err // Some other error occurred
	}
	return team.Role, nil // User is a member of the project
}
