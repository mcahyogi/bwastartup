package main

import (
	"bwastartup/db"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
)

func main() {
	conn := db.SetupConnection()
	// fmt.Println("database connected!", conn)
	// migration.RunMigration()

	userRepository := user.NewRepository(conn)
	userService := user.NewService(userRepository)
	userHandler := user.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1/")

	api.POST("users", userHandler.RegisterUser)

	router.Run()
}
