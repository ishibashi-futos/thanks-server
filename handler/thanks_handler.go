package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	database "thanks-server/db"
)

type thanksRequest struct {
	Message     string `json:"message"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

func PostThank(c echo.Context) error {
	request := new(thanksRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	db, err := database.NewThanksRepository()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer db.Close()
	db.Save(request.Message, request.Source, request.Destination)
	return c.JSON(http.StatusCreated, "")
}

func GetThanksCount(c echo.Context) error {
	var s database.Summaries
	var diff int
	diff, _ = strconv.Atoi(c.QueryParam("diff"))
	db, err := database.NewThanksRepository()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer db.Close()
	if s, err = db.Summary(diff); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, s)
}
