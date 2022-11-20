package app

import (
	"context"
	"disease-api/internal/app/config"
	"disease-api/internal/handler"
	"disease-api/internal/repository"
	"disease-api/internal/repository/connection"
	"disease-api/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func Run(filenames ...string) {
	cfg, err := config.New(filenames...)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	log := logrus.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"method": c.Request().Method,
				"URI":    values.URI,
				"status": values.Status,
			}).Info()

			return nil
		},
	}))

	ctx := context.Background()

	db, err := connection.DialPostgres(ctx, cfg.Postgres)
	if err != nil {
		logrus.Fatalf("unable to connect to postgres: %v", err)
	}

	repos := repository.New(db, cfg)
	services := service.New(repos)
	handlers := handler.New(services)
	handlers.Register(e)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
