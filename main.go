package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.GET("", helloWorld)

	e.Start(":4000")
}

func helloWorld(c echo.Context) error {
	c.JSON(200, "hello")
	return nil
}
