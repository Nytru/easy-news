package cmd

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nytru/easy-news/src/internal/handler"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v3"
)

type app struct {
	e *echo.Echo
	l *lecho.Logger
}

func (a *app) Run() error {
	a.e.Use(lecho.Middleware(lecho.Config{
		Logger: a.l,
	}))
	a.e.Use(middleware.Recover())
	a.e.GET("/health", handler.Health)

	return a.e.Start("0.0.0.0:" + os.Getenv("PORT"))
}

func NewApp() *app {
	e := echo.New()
	l := lecho.From(log.Logger)
	e.Logger = l

	return &app{
		e,
		l,
	}
}
