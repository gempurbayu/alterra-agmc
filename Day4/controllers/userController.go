package controllers

import (
	"day2-agmc/lib/database"
	"day2-agmc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,
		map[string]any{
			"code": 200,
			"data": users,
		},
	)
}

func CreateUserController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	userDb, err := database.CreateUser(user)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]any{
		"code": 200,
		"data": userDb,
	})
}

func GetUserByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := database.GetUserById(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,
		map[string]any{
			"code": 200,
			"data": user,
		})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := database.DeleteUserById(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK,
		map[string]any{
			"code": 200,
			"data": user,
		},
	)
}

func UpdateUserByIDController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	c.Bind(&user)

	userDb, err := database.UpdateUser(id, user)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, map[string]any{
		"code": 200,
		"data": userDb,
	})
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}