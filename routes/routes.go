package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(ech *echo.Echo) {
	ech.GET("/pets", getAnimals)
	ech.POST("/pets", createAnimal)
	ech.GET("/pets/:id", getAnimal)
	ech.PUT("/pets/:id", updateAnimal)
	ech.DELETE("/pets/:id", deleteAnimal)

	ech.Logger.Info("Starting server on port 8080")
	ech.Logger.Fatal(ech.Start(":8080"))
}