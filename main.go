package main

import (
	"brifast-service-login/auth"
	"brifast-service-login/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	load := godotenv.Load()
	if load != nil {
		fmt.Println("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	authRepository := auth.NewRepository(db)

	authService := auth.NewService(authRepository)

	authHandler := handler.NewAuthHandler(authService)

	router := gin.Default()
	router.Use(cors.Default())
	
	api := router.Group("api/v1")
	api.POST("/login", authHandler.Login)
	api.GET("/testing", func(c *gin.Context) {
		responseData := gin.H{"message": "Hello, world"}

		c.JSON(http.StatusOK, responseData)
	})

	router.Run()

}