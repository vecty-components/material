package app

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-components/material/appbar"
	"github.com/vecty-components/material/drawer"
)

// A is a vecty-material appbar component.
type A struct {
	vecty.Core
	RootMarkup  vecty.MarkupList `vecty:"prop"`
	ChildMarkup vecty.MarkupList `vecty:"prop"`
	Appbar      *appbar.A        `vecty:"prop"`
	Sidebar     *drawer.D        `vecty:"prop"`
	Routes      vecty.List       `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (c *A) Render() vecty.ComponentOrHTML {
	routes := make([]vecty.MarkupOrChild, len(c.Routes))
	for i, route := range c.Routes {
		routes[i] = route
	}

	// Built in root element.
	return elem.Div(
		c.Appbar,
		elem.Div(
			c.RootMarkup,
			c.Sidebar,
			elem.Div(
				append(
					[]vecty.MarkupOrChild{
						vecty.Markup(
							vecty.Class(
								"mdc-drawer-app-content",
								"mdc-top-app-bar--fixed-adjust",
							),
						),
						c.ChildMarkup,
					},
					routes...,
				)...,
			),
		),
	)
}
