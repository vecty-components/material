package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ComponentPage struct {
	component vecty.Component
	vecty.Core
}

func NewComponentPage(c vecty.Component) *ComponentPage {
	return &ComponentPage{component: c}
}

func (cp *ComponentPage) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("demo-panel"),
		),
		/* sidebar */
		elem.Div(
			vecty.Markup(
				vecty.Class(
					"demo-content", "mdc-drawer-app-content", "mdc-top-app-bar--fixed-adjust",
				),
			),
			cp.component,
		),
	)
}
