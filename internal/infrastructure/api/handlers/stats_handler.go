package handlers

import (
	"fizzbuzz/api/internal/domain/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetStatsHandler(c echo.Context, repo repositories.IStat) error {
	mostFrequentRequest := repo.GetMostFrequentRequest()
	return c.JSON(http.StatusOK, mostFrequentRequest)
}
