package routes

import (
	"net/http"
	"strconv"

	"example.com/golang-assignment/models"
	"github.com/labstack/echo/v4"
)

func getAnimals(c echo.Context) error {
	animals, err := models.GetAnimals()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return c.JSON(http.StatusOK, animals)
}

func createAnimal(c echo.Context) error {
	var a models.Animal
	if err := c.Bind(&a); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := a.Create(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return c.JSON(http.StatusCreated, a)
}

func getAnimal(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	a, err := models.GetAnimal(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Animal not found"})
	}
	return c.JSON(http.StatusOK, a)
}

func updateAnimal(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	var a models.Animal
	if err := c.Bind(&a); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	a.ID = id
	if err := a.Update(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return c.JSON(http.StatusOK, a)
}

func deleteAnimal(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	var a models.Animal
	a.ID = id
	if err := a.Delete(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return c.NoContent(http.StatusNoContent)
}