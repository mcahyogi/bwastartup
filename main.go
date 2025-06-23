package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func printNumber(n int) {
	fmt.Println(n)
}
func main() {
	dsn := "host=localhost user=postgres password=gandul288 dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to Database")

	printNumber(1)
	printNumber(2)
	printNumber(3)
	time.Sleep(1 * time.Second)

	// Payment
	// paymentService := payment.NewService()

	// User
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// Campaign
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)

	// Transaction
	transactionRepository := transaction.NewRepository(db)
	transactionService := transaction.NewService(transactionRepository, campaignRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	// param 1 = path routingnya, param 2 = folder di project kita
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	// USER
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/check_email", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// CAMPAIGN
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-image", authMiddleware(authService, userService), campaignHandler.UploadImage)

	// Transaction
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransaction)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)

	router.Run(":8080")

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", "No Bearer Token Provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized Validate Token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized Secret Key", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(string)
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}

// ambil nilai header Authorization: Bearer "tokentoken"
// amnbil nilai tokennya saja
// validasi token
// ambil user_id
// ambil user dari db berdasarkan id user_id lewat service
// set user dengan 'Context'
