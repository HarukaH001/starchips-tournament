package route

import (
	"github.com/HarukaH001/starchips-tournament/common"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type BaseRouterConfig struct {
	DB bun.IDB `validate:"required"`
}

func NewBaseRouter(router fiber.Router, cfg *BaseRouterConfig) {
	common.MustValid(cfg)
}
