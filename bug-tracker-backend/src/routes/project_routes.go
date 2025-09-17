package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/controllers"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/middlewares"
)

func ProjectRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.RequireAuth)
	router.POST("project", controllers.CreateProject)
	router.GET("project", controllers.GetAllProjects)
	router.GET("project/:projectID", middlewares.ProjectCheckMiddleware, controllers.GetProjectByID)
	router.PATCH("project/:projectID", middlewares.ProjectCheckMiddleware, controllers.UpdateProject)
	router.DELETE("project/:projectID", middlewares.ProjectCheckMiddleware, controllers.DeleteProject)
}

func TeamRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.RequireAuth)
	projectGroup := router.Group("project/:projectID/")
	projectGroup.Use(middlewares.ProjectCheckMiddleware)

	projectGroup.POST("team/add", controllers.AddToTeam)
	projectGroup.GET("team", controllers.GetTeamMembers)
	projectGroup.POST("team/action", controllers.TeamAction)
}

func BugRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.RequireAuth)
	projectGroup := router.Group("project/:projectID/")
	projectGroup.Use(middlewares.ProjectCheckMiddleware)

	projectGroup.POST("bug", controllers.CreateBug)
	projectGroup.GET("bug", controllers.GetAllBugs)
	projectGroup.GET("bug/:bugID", middlewares.BugCheckMiddleware, controllers.GetBugByID)
	projectGroup.PATCH("bug/:bugID", middlewares.BugCheckMiddleware, controllers.UpdateBug)
	projectGroup.DELETE("bug/:bugID", middlewares.BugCheckMiddleware, controllers.DeleteBug)
}
