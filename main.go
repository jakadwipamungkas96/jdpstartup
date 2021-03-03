package main

import (
	"jdpstartup/auth"
	"jdpstartup/handler"
	"jdpstartup/helper"
	"jdpstartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

	authService := auth.NewServiceToken()

	userService.SaveAva(1, "images/1-profile.png")

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// ENDPOINT
	api.POST("/registerusers", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/emailcheckers", userHandler.CheckEmailAvailability)
	api.POST("/uploadavatar", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unautorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // AbortWithStatusJSON = agar tidak lanjut ke proses selanjutnya
			return
		}

		// Bearer tokentoken
		tokenString := ""
		arrToken := strings.Split(authHeader, " ")

		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unautorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // AbortWithStatusJSON = agar tidak lanjut ke proses selanjutnya
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unautorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // AbortWithStatusJSON = agar tidak lanjut ke proses selanjutnya
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unautorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // AbortWithStatusJSON = agar tidak lanjut ke proses selanjutnya
			return
		}

		c.Set("currentUser", user) //Context
	}

}

// Middleware
// 1. Mengambil nilai header authorization: Bearer generateToken
// 2. header authorization -> get token value
// 3. Validasi Token, menggunakan service validateToken
// 4. jika valid, get user_id -> get user di db by user_id
