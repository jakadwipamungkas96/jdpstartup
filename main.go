package main

import (
	"jdpstartup/handler"
	"jdpstartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_cf?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepo(db)

	userService := user.NewService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// ENDPOINT
	api.POST("/registerusers", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}
