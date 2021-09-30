package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/drawer"
	"github.com/vecty-material/material/ul"
	router "marwan.io/vecty-router"
)

type link struct {
	content string
	url     string
}

type ComponentSidebar struct {
	vecty.Core
}

func (cs *ComponentSidebar) renderSidebarLink(link link, index int) vecty.ComponentOrHTML {
	return router.Link(
		link.url,
		link.content,
		router.LinkOptions{},
	)
}

func (cs *ComponentSidebar) renderDrawer() vecty.ComponentOrHTML {

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

	return elem.Div(
		&drawer.D{
			Type: drawer.Dismissible,
			Content: &ul.L{
				Items: items,
			},
		},
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
