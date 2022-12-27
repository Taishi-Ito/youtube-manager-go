package main

import(
	"youtube-manager-go/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
