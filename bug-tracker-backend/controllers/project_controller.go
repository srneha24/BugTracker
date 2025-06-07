package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/api"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/models"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/utils"
)

func CreateProject(c *gin.Context) {
	var project api.CreateProject
	ec := conf.EnhancedContext{Context: c}

	if err := c.ShouldBindJSON(&project); err != nil {
		ec.ValidationError(err.Error())
		return
	}

	user := utils.ExtractUserFromContext(c)

	newProject := models.Project{
		Title:       project.Title,
		Description: project.Description,
		CreatedBy:   user.ID,
	}

	// Start transaction
	tx := conf.DB.Begin()
	if tx.Error != nil {
		log.Println("Error while starting project creation transaction", tx.Error)
		ec.BadRequestWithMessageAndNoData("Failed to create project")
		return
	}

	if err := tx.Create(&newProject).Error; err != nil {
		tx.Rollback()
		log.Println("Error while creating project:", err)
		ec.BadRequestWithMessageAndNoData("Failed to create project")
		return
	}

	// Create project team
	projectTeam := models.Team{
		ProjectID: newProject.ID,
		UserID:    user.ID,
		Role:      api.TeamRoleAdmin.Value(),
	}

	if err := tx.Create(&projectTeam).Error; err != nil {
		tx.Rollback()
		log.Println("Error while creating team:", err)
		ec.BadRequestWithMessageAndNoData("Failed to create project")
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Println("Error while committing transaction:", err)
		ec.BadRequestWithMessageAndNoData("Failed to create project")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
		"data": api.ProjectResponse{
			ID:          int(newProject.ID),
			Title:       newProject.Title,
			Description: newProject.Description,
			CreatedBy:   int(newProject.CreatedBy),
			CreatedAt:   newProject.CreatedAt,
			UpdatedAt:   newProject.UpdatedAt,
		},
	})
}

func GetAllProjects(c *gin.Context) {
	var params api.ProjectListQueryParams
	var data []api.ProjectResponse
	ec := conf.EnhancedContext{Context: c}

	if err := c.ShouldBindQuery(&params); err != nil {
		ec.ValidationError(err.Error())
		return
	}

	user := utils.ExtractUserFromContext(c)

	query := conf.DB.Model(&models.Project{}).Joins("INNER JOIN teams ON teams.project_id = projects.id").Where("teams.user_id = ?", user.ID)

	if params.Search != nil && *params.Search != "" {
		query = query.Where("title ILIKE ?", "%"+*params.Search+"%")
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		log.Println("Error while counting projects:", err)
		ec.BadRequestWithNoMessageAndNoData()
		return
	}

	offset := (params.Page - 1) * params.Limit
	query = query.Limit(params.Limit).Offset(offset)

	if err := query.Find(&data).Error; err != nil {
		log.Println("Error while retrieving projects:", err)
		ec.BadRequestWithNoMessageAndNoData()
		return
	}

	paginatedResponse := utils.Paginate(data, params.Page, params.Limit, int(totalCount))
	c.JSON(http.StatusOK, paginatedResponse)

}

func GetProjectByID(c *gin.Context) {
}

func UpdateProject(c *gin.Context) {
}

func DeleteProject(c *gin.Context) {
}
