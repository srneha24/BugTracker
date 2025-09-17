package controllers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/models"
	types "github.com/WNBARookie/BugTracker/bug-tracker-backend/src/types"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/utils"
)

func CreateBug(c *gin.Context) {
	var bug types.CreateBug
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
		"data": types.BugResponse{
			ID:          newBug.ID,
			Title:       newBug.Title,
			Description: newBug.Description,
			Tags:        newBug.Tags,
			Deadline:    newBug.Deadline,
			Status:      types.BugStatus(newBug.Status),
			Priority:    types.Priority(newBug.Priority),
			AssignedTo: types.AssignedTo{
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
	var params types.BugListQueryParams
	ec := conf.EnhancedContext{Context: c}

	if err := c.ShouldBindQuery(&params); err != nil {
		ec.ValidationError(err.Error())
		return
	}

	project := utils.ExtractProjectFromContext(c)
	query := conf.DB.Model(&models.Bug{}).Where("project_id = ?", project.ID)

	if params.Search != nil {
		searchTerm := "%" + *params.Search + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ?", searchTerm, searchTerm)
	}

	if params.Tags != nil {
		// Split the comma-separated tags into a slice
		parts := strings.Split(*params.Tags, ",")
		tags := make([]string, 0, len(parts))
		for _, p := range parts {
			t := strings.TrimSpace(p)
			if t != "" {
				tags = append(tags, t)
			}
		}
		if len(tags) == 0 {
			tags = []string{*params.Tags}
		}

		query = query.Where("tags && ?", pq.Array(tags))
	}

	if params.Deadline != nil {
		// Parse the deadline string to time.Time
		deadline, err := time.Parse("2006-01-02", *params.Deadline)
		if err != nil {
			ec.ValidationError("Invalid deadline format (expected YYYY-MM-DD)")
			return
		}
		query = query.Where("deadline <= ?", deadline)
	}

	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	}

	if params.Priority != nil {
		query = query.Where("priority = ?", *params.Priority)
	}

	if params.AssignedTo != nil {
		query = query.Where("assigned_to = ?", *params.AssignedTo)
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		log.Println("Error while counting bugs:", err)
		ec.BadRequestWithNoMessageAndNoData()
		return
	}

	offset := (params.Page - 1) * params.Limit
	query = query.Limit(params.Limit).Offset(offset)

	var bugs []models.Bug
	if err := query.Find(&bugs).Error; err != nil {
		log.Println("Error while retrieving bugs:", err)
		ec.BadRequestWithNoMessageAndNoData()
		return
	}

	// Convert models.Bug to types.BugResponse
	data := make([]types.BugResponse, 0)
	for _, bug := range bugs {
		assignedTo, _ := utils.LookupUserUsingID(bug.AssignedTo)
		assignedToResponse := types.AssignedTo{
			ID:    bug.AssignedTo,
			Name:  "",
			Email: "",
		}
		if assignedTo != nil {
			assignedToResponse.Name = assignedTo.Name
			assignedToResponse.Email = assignedTo.Email
		}

		data = append(data, types.BugResponse{
			ID:          bug.ID,
			Title:       bug.Title,
			Description: bug.Description,
			Tags:        bug.Tags,
			Deadline:    bug.Deadline,
			Status:      types.BugStatus(bug.Status),
			Priority:    types.Priority(bug.Priority),
			AssignedTo:  assignedToResponse,
			ProjectID:   bug.ProjectID,
			CreatedAt:   bug.CreatedAt,
			UpdatedAt:   bug.UpdatedAt,
		})
	}

	paginatedResponse := utils.Paginate(data, params.Page, params.Limit, int(totalCount))
	c.JSON(http.StatusOK, paginatedResponse)
}

func GetBugByID(c *gin.Context) {
	ec := conf.EnhancedContext{Context: c}
	bug := utils.ExtractBugFromContext(c)

	ec.SuccessWithMessage(
		"Bug retrieved successfully",
		types.BugResponse{
			ID:          bug.ID,
			Title:       bug.Title,
			Description: bug.Description,
			Tags:        bug.Tags,
			Deadline:    bug.Deadline,
			Status:      types.BugStatus(bug.Status),
			Priority:    types.Priority(bug.Priority),
			AssignedTo: types.AssignedTo{
				ID:    bug.AssignedTo,
				Name:  "", // To be filled if needed
				Email: "", // To be filled if needed
			},
			ProjectID: bug.ProjectID,
			CreatedAt: bug.CreatedAt,
			UpdatedAt: bug.UpdatedAt,
		},
	)
}

func UpdateBug(c *gin.Context) {
}

func DeleteBug(c *gin.Context) {
}
