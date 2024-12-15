package controllers

import (
	"chat_api/models"
	"chat_api/services"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	UserService services.UserService
}

func NewUsersController(userService services.UserService) UsersController {
	return UsersController{UserService: userService}
}

func (u UsersController) GetAll(c echo.Context) error {
	return c.String(200, "Get All Users")
}

func (u UsersController) GetById(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, "Get User By Id: "+id)
}

func (u UsersController) Create(c echo.Context) error {
	user := new(models.User)
	user.Name = "John Doe"
	email := "r6xYI@example.com"
	user.Email = &email
	user.Age = 20
	user.Birthday = nil

	createdUser, err := u.UserService.CreateUser(*user)
	if err != nil {
		return err
	}
	return c.JSON(200, createdUser)

}

func (u UsersController) Update(c echo.Context) error {
	return c.String(200, "Update User")
}

func (u UsersController) Delete(c echo.Context) error {
	return c.String(200, "Delete User")
}
