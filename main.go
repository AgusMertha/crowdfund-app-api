package main

import (
	"kitabantu-api/auth"
	"kitabantu-api/campaign"
	"kitabantu-api/handler"
	"kitabantu-api/helper"
	"kitabantu-api/transaction"
	"kitabantu-api/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewUserRepository(db)
	campaignRepository := campaign.NewCampaignRepository(db)
	transactionRepository := transaction.NewTransactionRepository(db)

	authService := auth.NewJwtService()
	userService := user.NewUserService(userRepository)
	campaignService := campaign.NewCampaignService(campaignRepository)
	transactionService := transaction.NewTransactionService(transactionRepository, campaignRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTranscationHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./images")

	api := router.Group("/api/v1")

	api.POST("/login", userHandler.LoginUser)
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/email-check", userHandler.CheckEmailAvailability)
	api.POST("upload-avatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.POST("/campaigns/upload-image", authMiddleware(authService, userService), campaignHandler.UploadCampaignImage)
	api.GET("/campaigns/:id", campaignHandler.GetCampaignById)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetTransactionByCampaign)

	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransaction)

	router.Run(":3000")
}

func authMiddleware(authService auth.Service, userService user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
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
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userId)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
