package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/api"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/models"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/utils"
)

func CreateBug(c *gin.Context) {
	var bug api.CreateBug
	ec := conf.EnhancedContext{Context: c}

	if err := c.ShouldBindJSON(&bug); err != nil {
		ec.ValidationError(err.Error())
		return
	}

	assignedTo, _ := utils.LookupUserUsingID(bug.AssignedTo)
	if assignedTo == nil {
		ec.BadRequestWithMessageAndNoData("Assigned user not found")
		return
	}

	if time.Now().After(bug.Deadline) {
		ec.BadRequestWithMessageAndNoData("Deadline cannot be in the past")
		return
	}

	newBug := models.Bug{
		Title:       bug.Title,
		Description: bug.Description,
		Tags:        bug.Tags,
		Deadline:    bug.Deadline,
		Status:      bug.Status.Value(),
		Priority:    bug.Priority.Value(),
		AssignedTo:  bug.AssignedTo,
		ProjectID:   utils.ExtractProjectFromContext(c).ID,
	}

	result := conf.DB.Create(&newBug)
	if result.Error != nil {
		log.Println("Error while creating bug:", result.Error)
		ec.BadRequestWithMessageAndNoData("Failed to create bug")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bug created successfully",
		"data": api.BugResponse{
			ID:          newBug.ID,
			Title:       newBug.Title,
			Description: newBug.Description,
			Tags:        newBug.Tags,
			Deadline:    newBug.Deadline,
			Status:      api.BugStatus(newBug.Status),
			Priority:    api.Priority(newBug.Priority),
			AssignedTo: api.AssignedTo{
				ID:    assignedTo.ID,
				Name:  assignedTo.Name,
				Email: assignedTo.Email,
			},
			ProjectID: newBug.ProjectID,
			CreatedAt: newBug.CreatedAt,
			UpdatedAt: newBug.UpdatedAt,
		},
	})
}

func GetAllBugs(c *gin.Context) {
}

func GetBugByID(c *gin.Context) {
}

func UpdateBug(c *gin.Context) {
}

func DeleteBug(c *gin.Context) {
}
