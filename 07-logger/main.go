package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	myLog, err := os.OpenFile(
		"logs.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatalf("no se pudo abrir o crear el archivo de logs: %v", err)
	}
	defer myLog.Close()

	logConfig := middleware.LoggerConfig{
		Output: myLog,
	}

	e.Use(middleware.LoggerWithConfig(logConfig))
	e.GET("/", Index)

	log.Println(e.Start(":7878"))
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo")
}
