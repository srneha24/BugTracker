package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/models"
	api "github.com/WNBARookie/BugTracker/bug-tracker-backend/src/types"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/utils"
)

func ProjectCheckMiddleware(c *gin.Context) {
	var projectURI api.ProjectURI
	if err := c.ShouldBindUri(&projectURI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid project ID"})
		c.Abort()
		return
	}

	var project models.Project
	if err := conf.DB.First(&project, projectURI.ProjectID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Project not found"})
		c.Abort()
		return
	}

	// Check if the user is part of the project team
	user := utils.ExtractUserFromContext(c)
	role, err := utils.CheckIfUserIsProjectMember(user.ID, project.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error checking project membership"})
		c.Abort()
		return
	} else if role == "" {
		c.JSON(http.StatusForbidden, gin.H{"message": "User is not a member of this project"})
		c.Abort()
		return
	}

	c.Set("project", project)
	c.Set("userRole", role)
	c.Next()
}

func BugCheckMiddleware(c *gin.Context) {
	var bugURI api.BugURI
	if err := c.ShouldBindUri(&bugURI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid bug ID"})
		c.Abort()
		return
	}

	var bug models.Bug
	if err := conf.DB.First(&bug, bugURI.BugID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Bug not found"})
		c.Abort()
		return
	}

	c.Set("bug", bug)
	c.Next()
}
