package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/string", func(c echo.Context) error {
		return c.String(http.StatusOK, "1,2018-03-31,ALEXYS,LOZADA,38")
	})
	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
				<h1>Hola mundo</h1>
				<script>alert("hola mundo")</script>
			`)
	})
	e.GET("/no-content", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.Start(":8080")
}
