package main

import (
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/middlewares"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/routes"
)

func init() {
	conf.LoadEnvVars()
	conf.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.StandardResponseMiddleware)
	router.Use(middlewares.EnhancedContextMiddleware)

	// Routers
	apiV1 := router.Group("/api/v1")
	{
		routes.AuthRoutes(apiV1)
		routes.UserRoutes(apiV1)
		routes.ProjectRoutes(apiV1)
		routes.TeamRoutes(apiV1)
		routes.BugRoutes(apiV1)
	}

	router.Run(":8080")
}
