package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/controllers"
	"github.com/WNBARookie/BugTracker/bug-tracker-backend/src/middlewares"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("user/signup", controllers.SignUp)
	router.POST("user/login", controllers.Login)
}

func UserRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.RequireAuth)
	router.GET("user", controllers.GetUserProfile)
	router.PATCH("user", controllers.UpdateUserProfile)
	router.DELETE("user", controllers.DeleteUserProfile)
	router.GET("user/bugs", controllers.GetUserBugs)
}
