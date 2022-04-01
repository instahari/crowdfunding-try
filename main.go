package main

import (
	"crowdfunding/handler"
	"crowdfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// inputLogin := user.LoginUserInput{
	// 	Email:    "Hanifridwan40@gmail.com",
	// 	Password: "Hanif1234.,!",
	// }
	// user, err := userService.LoginUserInput(inputLogin)
	// if err != nil {
	// 	fmt.Println("Terjadi Kesalahan!")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("Selamat Datang", user.Name)
	// fmt.Println("Anda Berhasil Login")

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	router.Run()
}
