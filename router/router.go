package router

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
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
	api.POST("avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized with no Bearer Token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenBearer := ""
		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) == 2 {
			tokenBearer = splitToken[1]
		}

		token, err := authService.ValidateToken(tokenBearer)
		if err != nil {
			response := helper.ApiResponse("Unauthorized with Unvalidate Token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized with Token not valid", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(string)

		user, err := userService.GetUserById(userID)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
