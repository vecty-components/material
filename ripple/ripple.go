package ripple

import (
	"github.com/hexops/vecty"
	"github.com/vecty-material/material/material/ripple"
)

type R struct {
	*ripple.R
	Root      *vecty.HTML
	Disabled  bool `js:"disabled"`
	Unbounded bool `js:"unbounded"`
}

func (c *R) Apply(h *vecty.HTML) {
	c.R = ripple.New()
	c.Root = h
	// start := func() error {
	// 	return c.R.Start(c.Root.Node())
	// }
	// vecty.Property("vecty-material-ripple", js.InternalObject(start)).Apply(h)
	// vecty.Property("vecty-material-ripple", &start).Apply(h)
	vecty.Property("vecty-material-ripple", c).Apply(h)
}

func (c *R) Start() error {
	return c.R.Start(c.Root.Node())
}
