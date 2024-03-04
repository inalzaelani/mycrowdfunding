package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"latihanGo/handler"
	"latihanGo/user"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/latihan_golang_nuxt?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
}
