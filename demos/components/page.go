package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ComponentPage struct {
	panel *ComponentCatalogPanel
	vecty.Core
}

func NewComponentPage(
	designLink string,
	description string,
	docsLink string,
	sourceLink string,
	title string,
	hero vecty.ComponentOrHTML,
	demos vecty.ComponentOrHTML,
) *ComponentPage {
	return &ComponentPage{
		panel: &ComponentCatalogPanel{
			designLink:  designLink,
			description: description,
			demos:       demos,
			docsLink:    docsLink,
			hero:        hero,
			sourceLink:  sourceLink,
			title:       title,
		},
	}
}

func (cp *ComponentPage) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("demo-panel"),
		),
		&ComponentSidebar{},
		elem.Div(
			vecty.Markup(
				vecty.Class(
					"demo-content", "mdc-drawer-app-content", "mdc-top-app-bar--fixed-adjust",
				),
			),
			cp.panel,
		),
	)
}
