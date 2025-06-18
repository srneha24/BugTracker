package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/api"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/utils"
)

func GetUserProfile(c *gin.Context) {

	user := utils.ExtractUserFromContext(c)

	response := gin.H{
		"success": true,
		"message": "User found.",
		"data": gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"username":   user.Username,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	}
	c.JSON(http.StatusOK, response)

}

func UpdateUserProfile(c *gin.Context) {
}

func DeleteUserProfile(c *gin.Context) {
}

func GetUserBugs(c *gin.Context) {
	type bugResult struct {
		ID           uint      `json:"id"`
		Title        string    `json:"title"`
		Status       string    `json:"status"`
		Priority     uint      `json:"priority"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		ProjectID    uint      `json:"project_id"`
		ProjectTitle string    `json:"project_title"`
	}

	var userBugs []api.UserBugsResponse
	var rawResults []bugResult
	ec := conf.EnhancedContext{Context: c}

	user := utils.ExtractUserFromContext(c)

	query := conf.DB.Table("bugs").
		Select(`
            bugs.id,
            bugs.title,
            bugs.status,
            bugs.priority,
            bugs.created_at,
            bugs.updated_at,
            bugs.project_id as "project_id",
            projects.title as "project_title"
        `).
		Joins("JOIN projects ON bugs.project_id = projects.id").
		Where("bugs.assigned_to = ?", user.ID).
		Order("bugs.priority ASC").
		Limit(5)

	if err := query.Scan(&rawResults).Error; err != nil {
		log.Println("Error while retrieving bugs:", err)
		ec.BadRequestWithNoMessageAndNoData()
		return
	}

	for _, result := range rawResults {
		userBugs = append(userBugs, api.UserBugsResponse{
			ID:        result.ID,
			Title:     result.Title,
			Status:    api.BugStatus(result.Status),
			Priority:  api.Priority(result.Priority),
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
			Project: api.UserBugsProject{
				ID:           result.ProjectID,
				ProjectTitle: result.ProjectTitle,
			},
		})
	}

	ec.Success(userBugs)
}
