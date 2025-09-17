package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/conf"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/middlewares"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/routes"
)

// Create a single shared Gin router instance
var router *gin.Engine

func init() {
	// Load env vars & DB connection (runs once on cold start)
	conf.LoadEnvVars()
	conf.ConnectToDatabase()

	router = gin.New()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // TODO: update with your frontend domain
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
}

// Vercel entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
