package routes

import (
	"day2-agmc/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//routes
	e.GET("/v1/books", controllers.GetAllBookController)
	e.GET("/v1/books/:id", controllers.GetBookByID)

	e.POST("/v1/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)

	//eAuth := e.Group("")
	//eAuth.Use(echoMid.BasicAuth(m.BasicAuthDB))

	//eAuth.POST("/v1/books", controllers.CreateBookController)
	//eAuth.PUT("/v1/books/:id", controllers.UpdateBookByID)
	//eAuth.DELETE("/v1/books/:id", controllers.DeleteBooks)

	//eAuth.PUT("/v1/users/:id", controllers.UpdateUserByIDController)
	//eAuth.DELETE("/v1/users/:id", controllers.DeleteUserController)
	//eAuth.GET("/v1/users", controllers.GetAllUsersController)
	//eAuth.GET("/v1/users/:id", controllers.GetUserByIDController)

	eAuth := e.Group("/jwt")
	eAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eAuth.POST("/v1/books", controllers.CreateBookController)
	eAuth.PUT("/v1/books/:id", controllers.UpdateBookByID)
	eAuth.DELETE("/v1/books/:id", controllers.DeleteBooks)

	eAuth.PUT("/v1/users/:id", controllers.UpdateUserByIDController)
	eAuth.DELETE("/v1/users/:id", controllers.DeleteUserController)
	eAuth.GET("/v1/users", controllers.GetAllUsersController)
	eAuth.GET("/v1/users/:id", controllers.GetUserByIDController)

	return e
}
