package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ComponentPage struct {
	panel *ComponentCatalogPanel `vecty:"prop"`
	vecty.Core
}

func (cp *ComponentPage) Render() vecty.ComponentOrHTML {

	return elem.Div(
		vecty.Markup(
			vecty.Class("demo-content-transition"),
		),
		cp.panel,
	)
}
