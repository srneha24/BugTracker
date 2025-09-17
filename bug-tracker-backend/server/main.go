package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/middlewares"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/routes"
)

func init() {
	conf.LoadEnvVars()
	conf.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
