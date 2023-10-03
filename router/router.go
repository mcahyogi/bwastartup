package router

import (
	"bwastartup/auth"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouterUser(r *gin.Engine, db *gorm.DB) {
	// conn := database.SetupConnection()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := user.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1/")

	api.POST("users", userHandler.RegisterUser)
	api.POST("session", userHandler.Login)
	api.POST("email_checkers", userHandler.CheckEmailAvailability)
	api.POST("avatars", userHandler.UploadAvatar)

	router.Run()
}
