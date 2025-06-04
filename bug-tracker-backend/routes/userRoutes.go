package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/WNBARookie/BugTracker/bug-tracker-backend/controllers"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("user/signup", controllers.SignUp)
	router.POST("user/login", controllers.Login)
}

func UserRoutes(router *gin.RouterGroup) {
	router.GET("user", controllers.GetUserProfile)
	router.PATCH("user", controllers.UpdateUserProfile)
	router.DELETE("user", controllers.DeleteUserProfile)
	router.GET("user/bugs", controllers.GetUserBugs)
}
