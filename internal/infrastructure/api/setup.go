package api

import (
	"fizzbuzz/api/internal/domain/repositories"
	"fizzbuzz/api/internal/infrastructure/api/handlers"
	"github.com/labstack/echo/v4"
)

type Api interface {
	Setup() (*Api, error)
	Run() error
}

type Result struct {
	Context echo.Echo
}

type Repositories struct {
	Stat repositories.IStat
}

func Setup(repositories Repositories) *Result {
	router := echo.New()

	router.POST("/api/v1/fizzbuzz", func(c echo.Context) error {
		return handlers.CreateFizzBuzzHandler(c, repositories.Stat)
	})
	router.GET("/api/v1/stats", func(c echo.Context) error {
		return handlers.GetStatsHandler(c, repositories.Stat)
	})
	return &Result{*router}
}

func (r *Result) Run() error {
	r.Context.Logger.Fatal(r.Context.Start(":8080"))
	return nil
}
