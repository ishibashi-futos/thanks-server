package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"thanks-server/handler"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/api/thanks", handler.PostThank)
	e.GET("/api/thanks", handler.GetThanksCount)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
