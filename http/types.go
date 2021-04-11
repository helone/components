package http

import (
	"github.com/gofiber/fiber/v2"
)

type Option func(c *Client)

type Config struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

type Client struct {
	http   *fiber.App
	config *Config
}

type Groups struct {
	Prefix string
	Route  []Route
}

type Route struct {
	Method  string
	Path    string
	Handler func(ctx *fiber.Ctx) error
}
