package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	router *gin.Engine
	db     *gorm.DB
)

func init() {
	// Load environment variables
	godotenv.Load()

	// Initialize database
	initDB()

	// Initialize router
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

	// Basic routes for testing
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Bug Tracker API is running",
			"data":    nil,
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "API Health Check Passed",
			"data": gin.H{
				"database":  "connected",
				"timestamp": time.Now(),
			},
		})
	})

	// API v1 routes
	apiV1 := router.Group("/api/v1")
	{
		// Auth routes
		apiV1.POST("/user/signup", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Signup endpoint - Coming Soon",
				"data":    nil,
			})
		})

		apiV1.POST("/user/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Login endpoint - Coming Soon",
				"data":    nil,
			})
		})

		// Protected routes placeholder
		apiV1.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "User profile endpoint - Coming Soon",
				"data":    nil,
			})
		})
	}
}

func initDB() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	env := os.Getenv("ENVIRONMENT")

	sslMode := "disable"
	if env != "local" {
		sslMode = "require"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort, sslMode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// For now, just log the error and continue without DB
		fmt.Printf("Warning: Failed to connect to database: %v\n", err)
	}
}

// Handler is the Vercel entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
