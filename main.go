package main

import (
	"github.com/labstack/echo"
	"log"
	"os"
	"sandbox/internal/handler"
)

func main() {
	e := echo.New()
	e.GET("/", handler.Hello)
	e.POST("/mail", handler.Mail)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	e.Logger.Error(e.Start(":" + port))
}
