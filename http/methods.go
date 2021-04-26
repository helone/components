package http

import (
	"log"
	"os"
	"os/signal"
)

func (c *Client) Use(args ...interface{}) *Client {
	for _, arg := range args {
		c.http.Use(arg)
	}
	return c
}

func (c *Client) Inject(opts ...Option) *Client {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) StartServerWithGracefulShutdown() {
	idleConnClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := c.http.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}
		close(idleConnClosed)
	}()
	if err := c.http.Listen(":" + c.config.Port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
	<-idleConnClosed
}

func (c *Client) Routes(groups ...Groups) {
	for _, group := range groups {
		g := c.http.Group(group.Prefix)
		for _, route := range group.Route {
			g.Add(route.Method, route.Path, route.Handler)
		}
	}
}
