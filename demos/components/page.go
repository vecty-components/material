package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-components/material/base"
)

type ComponentPage struct {
	panel *ComponentCatalogPanel
	vecty.Core
}

func (cp *ComponentPage) Render() vecty.ComponentOrHTML {

	return elem.Div(
		vecty.Markup(
			vecty.Class("demo-content-transition"),
		),
		base.RenderStoredChild(cp.panel),
	)
}
