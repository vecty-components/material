package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/drawer"
	"github.com/vecty-components/material/ul"
	router "marwan.io/vecty-router"
)

type DemoLink struct {
	Name  string
	Url   string
	Image string
}

func NewComponentSidebar(list []DemoLink) *drawer.D {
	links := append([]DemoLink{
		{
			Name: "Home",
			Url:  "/",
		},
	}, list...)

	items := make([]vecty.ComponentOrHTML, len(links))
	for i, link := range links {
		items[i] = &ul.Item{
			Primary: router.Link(
				link.Url,
				link.Name,
				router.LinkOptions{},
			),
		}
	}

	return &drawer.D{
		Root: vecty.Markup(
			prop.ID("demo-drawer"),
			vecty.Class("demo-drawer", "mdc-top-app-bar--fixed-adjust"),
		),
		Type: drawer.Dismissible,
		Content: &ul.L{
			Items: items,
		},
	}
}
