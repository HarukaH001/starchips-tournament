package main

import (
	"github.com/HarukaH001/starchips-tournament/config"
	"github.com/HarukaH001/starchips-tournament/handler"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	envconfig.MustProcess("", &cfg)

	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if cfg.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}

	logger.Info("Starting Starchips Tournament server")

	serve, err := handler.New(&cfg, logger)
	if err != nil {
		logger.WithError(err).Fatal("failed to create handler")
	}

	if err := serve(); err != nil {
		logger.WithError(err).Fatal("failed to serve application")
	}
}
