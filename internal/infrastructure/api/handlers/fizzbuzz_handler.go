package handlers

import (
	"fizzbuzz/api/internal/application/command"
	"fizzbuzz/api/internal/domain/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateFizzBuzzRequest struct {
	Limit int64  `json:"limit,omitempty"`
	Int1  int64  `json:"int1,omitempty"`
	Str1  string `json:"str1,omitempty"`
	Int2  int64  `json:"int2,omitempty"`
	Str2  string `json:"str2,omitempty"`
}

func CreateFizzBuzzHandler(c echo.Context, repo repositories.IStat) error {

	var request CreateFizzBuzzRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	createFizzBuzzCommand := command.CreateFizzBuzzCommand{
		Limit: request.Limit,
		Int1:  request.Int1,
		Str1:  request.Str1,
		Int2:  request.Int2,
		Str2:  request.Str2,
		Repo:  repo,
	}

	output, err := createFizzBuzzCommand.Execute()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, output)
}
