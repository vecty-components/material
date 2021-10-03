package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/base"
)

type ComponentPage struct {
	panel *ComponentCatalogPanel
	vecty.Core
}

func (cp *ComponentPage) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/ComponentPage.css")

	return elem.Div(
		vecty.Markup(
			vecty.Class(
				"demo-content", "mdc-drawer-app-content", "mdc-top-app-bar--fixed-adjust",
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("demo-content-transition"),
			),
			base.RenderStoredChild(cp.panel),
		),
	)
}
