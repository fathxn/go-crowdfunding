package main

import (
	"go-crowdfunding/auth"
	"go-crowdfunding/handler"
	"go-crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	authService := auth.NewService()

	userRepoitory := user.NewRepository(db)
	userService := user.NewService(userRepoitory)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_check", userHandler.CheckEmailAvailable)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run("127.0.0.1:8080")
}
