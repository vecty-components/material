package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"

	"github.com/vecty-material/material/button"
	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/menu"
	"github.com/vecty-material/material/ul"
)

func NewMenuPage() *components.ComponentPage {
	return components.NewComponentPage(
		"Menu",
		"Menus display a list of choices on a transient sheet of material.",
		"https://material.io/go/design-menus",
		"https://material.io/components/web/catalog/menus/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-menu",
		components.NewHeroComponent(&MenuHero{}), &MenuDemos{},
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
			Open: true,
			Root: vecty.Markup(
				vecty.Class("hero-menu"),
			),
			List: &ul.L{Items: []vecty.ComponentOrHTML{
				&ul.Item{Primary: vecty.Text("A Menu Item")},
				&ul.Item{Primary: vecty.Text("Another Menu Item")},
			}},
		},
	)
}

type MenuDemos struct {
	vecty.Core
}

func (bd *MenuDemos) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/MenuCatalog.css")

	menu := &menu.M{
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
	}

	return elem.Div(
		elem.Heading3(
			vecty.Markup(
				vecty.Class("mdc-typography--subtitle1"),
			),
			vecty.Text("Anchored Menu"),
		),
		&button.B{
			Label: elem.Anchor(
				vecty.Markup(
					event.Click(func(e *vecty.Event) {
						menu.Open = !menu.Open
						vecty.Rerender(menu)
					}),
				),
				vecty.Text("Open menu"),
			),
		},
		menu,
	)
}
