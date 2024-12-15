package main

import (
	"chat_api/configs"
	"chat_api/controllers"
	"chat_api/repositories"
	"chat_api/routes"
	"chat_api/services"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config, err := configs.GetConfig()
	if err != nil {
		e.Logger.Fatal("err")
	}

	db, err := config.ConnectDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	fmt.Println(db)
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userControllers := controllers.NewUsersController(userService)
	routes := routes.NewRoute(userControllers)
	routes.UserRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.ServerIP, config.ServerPort)))
}
