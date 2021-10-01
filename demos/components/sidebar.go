package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/drawer"
	"github.com/vecty-material/material/ul"
)

type link struct {
	content string
	url     string
}

type ComponentSidebar struct {
	drawer *drawer.D
	vecty.Core
}

func NewComponentSidebar() *ComponentSidebar {
	return &ComponentSidebar{}
}

func (cs *ComponentSidebar) Toggle() {
	if cs.drawer != nil {
		cs.drawer.Open = !cs.drawer.Open
		vecty.Rerender(cs.drawer)
	}
}

func (cs *ComponentSidebar) renderSidebarLink(link link, index int) vecty.ComponentOrHTML {
	return ul.ItemLink(
		link.url,
		link.content,
	)
}

func (cs *ComponentSidebar) renderDrawer() vecty.ComponentOrHTML {
	if cs.drawer != nil {
		return elem.Div(
			cs.drawer,
		)
	}

	links := []link{
		{
			content: "Home",
			url:     "/",
		}, {
			content: "Button",
			url:     "/button",
		}, {
			content: "Card",
			url:     "/card",
		}, {
			content: "Checkbox",
			url:     "/checkbox",
		}, {
			content: "Chips",
			url:     "/chips",
		}, {
			content: "Data Table",
			url:     "/data-table",
		}, {
			content: "Dialog",
			url:     "/dialog",
		}, {
			content: "Drawer",
			url:     "/drawer",
		}, {
			content: "Elevation",
			url:     "/elevation",
		}, {
			content: "FAB",
			url:     "/fab",
		}, {
			content: "Icon Button",
			url:     "/icon-button",
		}, {
			content: "Image List",
			url:     "/image-list",
		}, {
			content: "Layout Grid",
			url:     "/layout-grid",
		}, {
			content: "Linear Progress Indicator",
			url:     "/linear-progress-indicator",
		}, {
			content: "List",
			url:     "/list",
		}, {
			content: "Menu",
			url:     "/menu",
		}, {
			content: "Radio Button",
			url:     "/radio",
		}, {
			content: "Ripple",
			url:     "/ripple",
		}, {
			content: "Select",
			url:     "/select",
		}, {
			content: "Slider",
			url:     "/slider",
		}, {
			content: "Snackbar",
			url:     "/snackbar",
		}, {
			content: "Switch",
			url:     "/switch",
		}, {
			content: "Tab Bar",
			url:     "/tabs",
		}, {
			content: "Text Field",
			url:     "/text-field",
		}, {
			content: "Top App Bar",
			url:     "/top-app-bar",
		}, {
			content: "Typography",
			url:     "/typography",
		},
	}
	items := make([]vecty.ComponentOrHTML, len(links))
	for i, link := range links {
		items[i] = cs.renderSidebarLink(link, i)
	}

	cs.drawer = &drawer.D{
		Type: drawer.Dismissible,
		Content: &ul.L{
			Items: items,
		},
	}

	return elem.Div(
		cs.drawer,
	)
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
