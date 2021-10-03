package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/drawer"
	"github.com/vecty-material/material/ul"
	router "marwan.io/vecty-router"
)

type DemoLink struct {
	Name  string
	Url   string
	Image string
}

type ComponentSidebar struct {
	drawer *drawer.D
	list   []DemoLink
	vecty.Core
}

func NewComponentSidebar(list []DemoLink) *ComponentSidebar {
	return &ComponentSidebar{
		list: list,
	}
}

func (cs *ComponentSidebar) Toggle() {
	if cs.drawer != nil {
		cs.drawer.Open = !cs.drawer.Open
		vecty.Rerender(cs.drawer)
	}
}

func (cs *ComponentSidebar) renderSidebarLink(link *DemoLink, index int) vecty.ComponentOrHTML {
	return &ul.Item{
		Primary: router.Link(
			link.Url,
			link.Name,
			router.LinkOptions{},
		),
	}
}

func (cs *ComponentSidebar) Render() vecty.ComponentOrHTML {
	if cs.drawer == nil {
		links := append([]DemoLink{
			{
				Name: "Home",
				Url:  "/",
			},
		}, cs.list...)

		items := make([]vecty.ComponentOrHTML, len(links))
		for i, link := range links {
			items[i] = cs.renderSidebarLink(&link, i)
		}

		cs.drawer = &drawer.D{
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

	return cs.drawer
}
