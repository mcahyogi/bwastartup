package main

import (
	"bwastartup/router"
	"log"

	"bwastartup/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// "bwastartup/user"

// "github.com/gin-gonic/gin"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := db.SetupConnection()
	r := gin.Default()

	router.SetupRouter(r, db)
	// conn := db.SetupConnection()
	// fmt.Println("database connected!", conn)
	// migration.RunMigration()

	// userRepository := user.NewRepository(conn)
	// userService := user.NewService(userRepository)
	// userHandler := user.NewUserHandler(userService)

	// router := gin.Default()

	// api := router.Group("/api/v1/")

	// api.POST("users", userHandler.RegisterUser)
	// api.POST("session", userHandler.Login)
	// api.POST("email_checkers", userHandler.CheckEmailAvailability)
	// api.POST("avatars", userHandler.UploadAvatar)

	// router.Run()
}
