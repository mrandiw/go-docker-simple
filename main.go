package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", Hello)
	e.GET("/get-student", GetStudent)
	e.GET("/get-grade", GetGrade)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetStudent(c echo.Context) error {
	return c.String(http.StatusOK, "Test API Get Student.")
}

func GetGrade(c echo.Context) error {
	return c.String(http.StatusOK, "Test API Get Grade.")
}
