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
	e.POST("/crypt/cfb", handler.CryptCFB)
	e.POST("/decrypt/cfb", handler.DecryptCFB)
	e.POST("/crypt/cbc", handler.CryptCBC)
	e.POST("/decrypt/cbc", handler.DecryptCBC)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	e.Logger.Error(e.Start(":" + port))
}
