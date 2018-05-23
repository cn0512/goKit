package main

/*
import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", Handler)
	e.GET("/api/:a", HandlerA)
	e.GET("/api", HandlerB)
	e.Static("/static", "static")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, yk!")
}

func HandlerA(c echo.Context) error {
	a := c.Param("a")
	return c.String(http.StatusOK, a)
}

func HandlerB(c echo.Context) error {
	a := c.QueryParam("name")
	b := c.QueryParam("age")
	return c.String(http.StatusOK, "result="+a+","+b)
}
*/
