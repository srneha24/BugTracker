package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/controllers"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/middlewares"
)

func ProjectRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.RequireAuth)
	router.POST("project", controllers.CreateProject)
	router.GET("project", controllers.GetAllProjects)
	router.GET("project/:projectID", controllers.GetProjectByID)
	router.PATCH("project/:projectID", controllers.UpdateProject)
	router.DELETE("project/:projectID", controllers.DeleteProject)
}

func TeamRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.RequireAuth)
	projectGroup := router.Group("project/:projectID/")

	projectGroup.POST("team/add", controllers.AddToTeam)
	projectGroup.GET("team", controllers.GetTeamMembers)
	projectGroup.POST("team/action", controllers.TeamAction)
}

func BugRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.RequireAuth)
	projectGroup := router.Group("project/:projectID/")

	projectGroup.POST("bug", controllers.CreateBug)
	projectGroup.GET("bug", controllers.GetAllBugs)
	projectGroup.GET("bug/:bugID", controllers.GetBugByID)
	projectGroup.PATCH("bug/:bugID", controllers.UpdateBug)
	projectGroup.DELETE("bug/:bugID", controllers.DeleteBug)
}
