package main

import (
	"example.com/golang-assignment/db"
	"example.com/golang-assignment/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	ech:=echo.New()
	routes.RegisterRoutes(ech)
	
}
