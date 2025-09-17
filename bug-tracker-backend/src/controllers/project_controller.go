package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/models"
	types "github.com/WNBARookie/BugTracker/bug-tracker-backend/src/types"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/utils"
)

func CreateProject(c *gin.Context) {
	var project types.CreateProject
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
		Role:      types.TeamRoleAdmin.Value(),
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
		"data": types.ProjectResponse{
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
	var params types.ProjectListQueryParams
	var data []types.ProjectResponse
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
	ec := conf.EnhancedContext{Context: c}
	project := utils.ExtractProjectFromContext(c)

	ec.SuccessWithMessage(
		"Project retrieved successfully",
		types.ProjectResponse{
			ID:          int(project.ID),
			Title:       project.Title,
			Description: project.Description,
			CreatedBy:   int(project.CreatedBy),
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		},
	)
}

func UpdateProject(c *gin.Context) {
	var updatedProject types.UpdateProject
	ec := conf.EnhancedContext{Context: c}
	project := utils.ExtractProjectFromContext(c)

	if err := c.ShouldBindJSON(&updatedProject); err != nil {
		ec.ValidationError(err.Error())
		return
	}

	updateData := make(map[string]any)

	if updatedProject.Title != nil {
		updateData["title"] = *updatedProject.Title
	}
	if updatedProject.Description != nil {
		updateData["description"] = *updatedProject.Description
	}

	result := conf.DB.Model(&project).Updates(updateData)

	if result.Error != nil {
		ec.BadRequestWithMessage("Failed to update project: ", result.Error.Error())
		return
	}

	ec.SuccessWithMessage(
		"Project updated successfully",
		types.ProjectResponse{
			ID:          int(project.ID),
			Title:       project.Title,
			Description: project.Description,
			CreatedBy:   int(project.CreatedBy),
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		},
	)
}

func DeleteProject(c *gin.Context) {
	ec := conf.EnhancedContext{Context: c}
	project := utils.ExtractProjectFromContext(c)

	// Check if user is admin of the project
	contextUserRole, exists := c.Get("userRole")
	if !exists || contextUserRole != types.TeamRoleAdmin.Value() {
		c.JSON(http.StatusForbidden, gin.H{"message": "Only project admins can delete the project"})
		return
	}

	if err := conf.DB.Delete(&project).Error; err != nil {
		log.Println("Error while deleting project:", err)
		ec.BadRequestWithMessageAndNoData("Failed to delete project")
		return
	}

	ec.SuccessWithMessageAndNoData("Project deleted successfully")
}
