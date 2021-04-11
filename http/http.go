package http

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func New(config *Config) *Client {
	return &Client{
		http: fiber.New(fiber.Config{
			ReadTimeout:           time.Second * time.Duration(config.ReadTimeout),
			WriteTimeout:          time.Second * time.Duration(config.ReadTimeout),
			DisableStartupMessage: true,
		}),
		config: config,
	}
}
