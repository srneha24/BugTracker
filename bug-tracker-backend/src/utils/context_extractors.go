package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/models"
)

func ExtractUserFromContext(c *gin.Context) models.User {
	contextUser, _ := c.Get("user")
	user, _ := contextUser.(models.User)

	return user
}

func ExtractProjectFromContext(c *gin.Context) models.Project {
	contextProject, _ := c.Get("project")
	project, _ := contextProject.(models.Project)

	return project
}

func ExtractBugFromContext(c *gin.Context) models.Bug {
	contextBug, _ := c.Get("bug")
	bug, _ := contextBug.(models.Bug)

	return bug
}
