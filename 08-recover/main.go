package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/division", Dividir)
	e.Start(":7171")
}

func Dividir(c echo.Context) error {
	d := c.QueryParam("d")
	f, _ := strconv.Atoi(d)
	a := 25000 / f
	return c.String(http.StatusOK, strconv.Itoa(a))
}
