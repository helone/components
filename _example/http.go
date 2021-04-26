package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helone/components/http"
)

func main() {
	c := &http.Config{
		Port:        "3000",
		ReadTimeout: 1,
	}
	h := http.New(c)
	h.Inject(func(c *http.Client) {
		c.Use()
		c.Routes(web(), api())
	})
	h.StartServerWithGracefulShutdown()
}

func web() http.Groups {
	return http.Groups{
		Prefix: "/web",
		Route: []http.Route{
			{
				Method: "GET",
				Path:   "/",
				Handler: func(ctx *fiber.Ctx) error {
					return ctx.JSON(fiber.Map{"route": "web"})
				},
			},
		},
	}
}

func api() http.Groups {
	return http.Groups{
		Prefix: "/api",
		Route: []http.Route{
			{
				Method: "GET",
				Path:   "/",
				Handler: func(ctx *fiber.Ctx) error {
					return ctx.JSON(fiber.Map{"route": "api"})
				},
			},
		},
	}
}
