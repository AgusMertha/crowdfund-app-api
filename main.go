package main

import (
	"kitabantu-api/handler"
	"kitabantu-api/user"
	"log"

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
	userService := user.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/login", userHandler.LoginUser)
	api.POST("/users", userHandler.RegisterUser)
	api.POST("email-check", userHandler.CheckEmailAvailability)
	router.Run(":3000")
}
