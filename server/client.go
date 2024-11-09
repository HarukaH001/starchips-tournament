package server

import (
	"fmt"

	"github.com/HarukaH001/starchips-tournament/common"
	"github.com/HarukaH001/starchips-tournament/config"
	"github.com/HarukaH001/starchips-tournament/route.client"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ClientServerConfig struct {
	PostgresConfig *config.PostgresConfig `validate:"required"`
}

func RegisterClientServer(router fiber.Router, log logrus.FieldLogger, cfg *ClientServerConfig) error {
	if log == nil {
		log = logrus.New()
	}
	common.MustValid(cfg)

	db, err := cfg.PostgresConfig.NewClient(log.WithField("component", "postgres"))
	if err != nil {
		return fmt.Errorf("failed to create postgres client: %w", err)
	}

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	route.NewBaseRouter(router.Group("/"), &route.BaseRouterConfig{
		DB: db,
	})

	return nil
}
