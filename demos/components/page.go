package components

import (
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

type ComponentPage struct {
	routes  map[string]*ComponentCatalogPanel
	sidebar *ComponentSidebar
	vecty.Core
}

func NewComponentPage(routes map[string]*ComponentCatalogPanel, sidebar *ComponentSidebar) *ComponentPage {
	return &ComponentPage{
		routes:  routes,
		sidebar: sidebar,
	}
}

func (cp *ComponentPage) Render() vecty.ComponentOrHTML {
	path := js.Global().Get("window").Get("location").Get("pathname").String()
	var p *ComponentCatalogPanel
	for route, panel := range cp.routes {
		if route != path {
			continue
		}
		p = panel
	}

	if p == nil {
		router.Redirect("/")
		return elem.Div()
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("demo-panel"),
		),
		cp.sidebar,
		elem.Div(
			vecty.Markup(
				vecty.Class(
					"demo-content", "mdc-drawer-app-content", "mdc-top-app-bar--fixed-adjust",
				),
			),
			p,
		),
	)
}
