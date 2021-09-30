package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

type ComponentPage struct {
	routes  []*router.Route
	sidebar *ComponentSidebar
	vecty.Core
}

func NewComponentPage(routes []*router.Route, sidebar *ComponentSidebar) *ComponentPage {
	return &ComponentPage{
		routes:  routes,
		sidebar: sidebar,
	}
}

func (cp *ComponentPage) Render() vecty.ComponentOrHTML {
	children := make([]vecty.MarkupOrChild, len(cp.routes)+1)
	children[0] = vecty.Markup(
		vecty.Class(
			"demo-content", "mdc-drawer-app-content", "mdc-top-app-bar--fixed-adjust",
		),
	)
	for i, route := range cp.routes {
		children[i] = route
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("demo-panel"),
		),
		cp.sidebar,
		elem.Div(
			children...,
		),
	)
}
