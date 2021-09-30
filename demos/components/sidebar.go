package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ComponentSidebar struct {
	vecty.Core
}

func (cs *ComponentSidebar) renderDrawer() vecty.ComponentOrHTML {
	return elem.Div()
}

func (cs *ComponentSidebar) renderScrim() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("mdc-drawer-scrim"),
		),
	)
}

func (cs *ComponentSidebar) Render() vecty.ComponentOrHTML {
	return elem.Div(
		cs.renderDrawer(),
		cs.renderScrim(),
	)
}
