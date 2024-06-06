package main

import (
	"go-crowdfunding/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepoitory := user.NewRepository(db)
	userService := user.NewService(userRepoitory)

	userInput := user.RegisterUserInput{
		Name:       "Test simpan",
		Occupation: "Designer",
		Email:      "test@simpan.com",
		Password:   "password",
	}
	userService.RegisterUser(userInput)
}
