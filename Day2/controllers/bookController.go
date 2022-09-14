package controllers

import (
	"day2-agmc/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	books = map[int]*models.Book{}
	count = 1
)

func GetAllBookController(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func CreateBookController(c echo.Context) error {
	b := &models.Book{
		ID:        count,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.Request().Header.Get("Content-type")

	if err := c.Bind(b); err != nil {
		return err
	}

	books[b.ID] = b
	count++
	return c.JSON(http.StatusCreated, b)
}

func GetBookByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func UpdateBookByID(c echo.Context) error {
	b := new(models.Book)

	if err := c.Bind(b); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Title = b.Title
	books[id].UpdatedAt = time.Now()

	return c.JSON(http.StatusOK, books[id])
}

func DeleteBooks(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	delete(books, id)

	return c.NoContent(http.StatusNoContent)
}
