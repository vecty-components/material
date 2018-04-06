package ripple

import (
	"agamigo.io/material/ripple"
	"github.com/gopherjs/vecty"
)

type R struct {
	*ripple.R
	Root      *vecty.HTML
	Disabled  bool `js:"disabled"`
	Unbounded bool `js:"unbounded"`
}

func (c *R) Apply(h *vecty.HTML) {
	c.R = ripple.New()
	vecty.Property("vecty-material-ripple", c.Start).Apply(h)
	c.Root = h
}

func (c *R) Start() error {
	return c.R.Start(c.Root.Node())
}
