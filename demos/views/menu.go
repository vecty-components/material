package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-material/material/button"
	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/menu"
	"github.com/vecty-material/material/ul"
)

func NewMenuPage() *components.ComponentCatalogPanel {
	return components.NewComponentCatalogPanel(
		"Menu",
		"Menus display a list of choices on a transient sheet of material.",
		"https://material.io/go/design-menus",
		"https://material.io/components/web/catalog/menus/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-menu",
		components.NewHeroComponent(
			&MenuHero{},
		),
		NewMenuDemos(),
	)
}

type MenuHero struct {
	vecty.Core
}

func (bh *MenuHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),
		&menu.M{
			Root: vecty.Markup(
				vecty.Class("hero-menu"),
			),
			AnchorElement: elem.Div(),
			List: &ul.L{Items: []vecty.ComponentOrHTML{
				&ul.Item{Primary: vecty.Text("A Menu Item")},
				&ul.Item{Primary: vecty.Text("Another Menu Item")},
			}},
		},
	)
}

type MenuDemos struct {
	menu *menu.M
	vecty.Core
}

func NewMenuDemos() *MenuDemos {
	return &MenuDemos{
		menu: &menu.M{
			AnchorElement: elem.Div(),
			List: &ul.L{Items: []vecty.ComponentOrHTML{
				&ul.Item{Primary: vecty.Text("PassionFruit")},
				&ul.Item{Primary: vecty.Text("Orange")},
				&ul.Item{Primary: vecty.Text("Guava")},
				&ul.Item{Primary: vecty.Text("Pitaya")},
				ul.ItemDivider(),
				&ul.Item{Primary: vecty.Text("Pineapple")},
				&ul.Item{Primary: vecty.Text("Mango")},
				&ul.Item{Primary: vecty.Text("Papaya")},
				&ul.Item{Primary: vecty.Text("Lychee")},
			}},
		},
	}
}

func (bd *MenuDemos) handleOpenClick(this *button.B, e *vecty.Event) {
	bd.menu.Open = !bd.menu.Open
	vecty.Rerender(bd.menu)
}

func (bd *MenuDemos) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/MenuCatalog.css")

	return elem.Div(
		elem.Heading3(
			vecty.Markup(
				vecty.Class("mdc-typography--subtitle1"),
			),
			vecty.Text("Anchored Menu"),
		),
		&button.B{
			Label:   vecty.Text("Open menu"),
			OnClick: bd.handleOpenClick,
		},
		bd.menu,
	)
}
