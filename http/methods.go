package http

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

func (c *Client) Run() (err error) {
	err = c.http.Listen(":" + c.config.Port)
	return
}

func (c *Client) Routes(groups ...Groups) {
	for _, group := range groups {
		g := c.http.Group(group.Prefix)
		for _, route := range group.Route {
			g.Add(route.Method, route.Path, route.Handler)
		}
	}
}
