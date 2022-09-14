package routes

import (
	"day2-agmc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	//routes
	e.GET("/v1/books", controllers.GetAllBookController)
	e.GET("/v1/books/:id", controllers.GetBookByID)
	e.POST("/v1/books", controllers.CreateBookController)
	e.PUT("/v1/books/:id", controllers.UpdateBookByID)
	e.DELETE("/v1/books/:id", controllers.DeleteBooks)

	e.GET("/v1/users", controllers.GetAllUsersController)
	e.GET("/v1/users/:id", controllers.GetUserByIDController)
	e.POST("/v1/users", controllers.CreateUserController)
	e.PUT("/v1/users/:id", controllers.UpdateUserByIDController)
	e.DELETE("/v1/users/:id", controllers.DeleteUserController)

	return e
}
