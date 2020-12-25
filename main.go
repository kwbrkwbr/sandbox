package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
	"sandbox/internal/handler"
	"sandbox/internal/server"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// qiitaのやつ、あとでリファクタ
	e.Renderer = server.Init()
	server.Router(e)

	// route
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
