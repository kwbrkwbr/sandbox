package server

import "github.com/labstack/echo"

func Router(e *echo.Echo) {
	e.GET("/", Root)
	e.GET("/fetch", FetchMember)
	e.POST("/put", PutMember)
	e.GET("/delete_all", DeleteAllMember)
	e.POST("/delete", DeleteMember)
}
