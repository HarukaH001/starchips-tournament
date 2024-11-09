package handler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/HarukaH001/starchips-tournament/config"
	"github.com/HarukaH001/starchips-tournament/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

type Serve func() error

func New(cfg *config.Config, log *logrus.Logger) (Serve, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	adminApp := fiber.New()
	adminApp.Use(cors.New())
	adminApp.Use(recover.New())
	adminApp.Use(logger.New(logger.Config{
		Output: log.Writer(),
	}))

	clientApp := fiber.New()
	clientApp.Use(cors.New())
	clientApp.Use(recover.New())
	clientApp.Use(logger.New(logger.Config{
		Output: log.Writer(),
	}))

	if err := server.RegisterAdminServer(adminApp, log.WithField("server", "admin"), &server.AdminServerConfig{
		PostgresConfig: &cfg.Postgres,
	}); err != nil {
		return nil, fmt.Errorf("failed to register admin server: %w", err)
	}

	if err := server.RegisterClientServer(clientApp, log.WithField("server", "client"), &server.ClientServerConfig{
		PostgresConfig: &cfg.Postgres,
	}); err != nil {
		return nil, fmt.Errorf("failed to register client server: %w", err)
	}

	return func() error {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-ch
			adminApp.Shutdown()
			clientApp.Shutdown()
		}()

		errCh := make(chan error, 2)
		go func() {
			errCh <- adminApp.Listen(cfg.AdminServer.GetAddress())
		}()

		go func() {
			errCh <- clientApp.Listen(cfg.ClientServer.GetAddress())
		}()

		return <-errCh
	}, nil
}
