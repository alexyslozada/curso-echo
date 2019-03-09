package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.Static("/", "public")
	e.GET("/gopher/:name", func(c echo.Context) error {
		p := c.Param("name")
		if p == "svg" {
			return c.Inline("imgs/gopher.svg", "gopher.svg")
		}
		if p == "jpg" {
			return c.Inline("imgs/to.jpg", "alexys")
		}
		if p == "att" {
			return c.Attachment("imgs/gopher.svg", "gopher.svg")
		}
		return c.HTML(http.StatusNotFound, "<p>No encontrado</p>")
	})
	e.Start(":8080")
}
