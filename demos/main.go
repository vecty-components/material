package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"

	"github.com/vecty-material/material/app"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/demos/views"
)

func main() {
	base.SetViewport()
	base.AddStyles()

	vecty.AddStylesheet("/assets/styles/App.css")
	vecty.AddStylesheet("/assets/styles/HeaderBar.css")
	vecty.AddStylesheet("/assets/styles/ComponentPage.css")
	vecty.AddStylesheet("/assets/styles/ComponentCatalogPanel.css")
	vecty.AddStylesheet("/assets/styles/CatalogPage.css")
	vecty.AddStylesheet("/assets/styles/HeroComponent.css")

	vecty.AddStylesheet("/assets/styles/ImageListCatalog.css")
	vecty.AddStylesheet("/assets/styles/ButtonCatalog.css")
	vecty.AddStylesheet("/assets/styles/MenuCatalog.css")
	vecty.AddStylesheet("/assets/styles/ListCatalog.css")
	vecty.AddStylesheet("/assets/styles/CheckboxCatalog.css")

	base.Boot()

	body := &Body{}
	vecty.RenderBody(body)
}

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() vecty.ComponentOrHTML {

	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		&CatalogPage{},
	)
}

type CatalogPage struct {
	vecty.Core
}

func (c *CatalogPage) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Material Components Web | Catalog")

	componentList := []components.DemoLink{
		{
			Name:  "Button",
			Url:   "/button",
			Image: "/assets/images/buttons_180px.svg",
		}, {
			Name:  "Card",
			Url:   "/card",
			Image: "/assets/images/cards_180px.svg",
		}, {
			Name:  "Checkbox",
			Url:   "/checkbox",
			Image: "/assets/images/checkbox_180px.svg",
		}, {
			Name:  "Chips",
			Url:   "/chips",
			Image: "/assets/images/chips_180px.svg",
		}, {
			Name:  "Data Table",
			Url:   "/data-table",
			Image: "/assets/images/data_table_180px.svg",
		}, {
			Name:  "Dialog",
			Url:   "/dialog",
			Image: "/assets/images/dialog_180px.svg",
		}, {
			Name:  "Drawer",
			Url:   "/drawer",
			Image: "/assets/images/drawer_180px.svg",
		}, {
			Name:  "Elevation",
			Url:   "/elevation",
			Image: "/assets/images/elevation_180px.svg",
		}, {
			Name:  "FAB",
			Url:   "/fab",
			Image: "/assets/images/floating_action_button_180px.svg",
		}, {
			Name:  "Icon Button",
			Url:   "/icon-button",
			Image: "/assets/images/icon_button_180px.svg",
		}, {
			Name:  "Image List",
			Url:   "/image-list",
			Image: "/assets/images/image_list_180px.svg",
		}, {
			Name:  "Layout Grid",
			Url:   "/layout-grid",
			Image: "/assets/images/layout_grid_180px.svg",
		}, {
			Name:  "Linear Progress Indicator",
			Url:   "/linear-progress-indicator",
			Image: "/assets/images/linear_progress_180px.svg",
		}, {
			Name:  "List",
			Url:   "/list",
			Image: "/assets/images/list_180px.svg",
		}, {
			Name:  "Menu",
			Url:   "/menu",
			Image: "/assets/images/menu_180px.svg",
		}, {
			Name:  "Radio Button",
			Url:   "/radio",
			Image: "/assets/images/radio_180px.svg",
		}, {
			Name:  "Ripple",
			Url:   "/ripple",
			Image: "/assets/images/ripple_180px.svg",
		}, {
			Name:  "Select",
			Url:   "/select",
			Image: "/assets/images/form_field_180px.svg",
		}, {
			Name:  "Slider",
			Url:   "/slider",
			Image: "/assets/images/slider_180px.svg",
		}, {
			Name:  "Snackbar",
			Url:   "/snackbar",
			Image: "/assets/images/snackbar_180px.svg",
		}, {
			Name:  "Switch",
			Url:   "/switch",
			Image: "/assets/images/switch_180px.svg",
		}, {
			Name:  "Tab Bar",
			Url:   "/tabs",
			Image: "/assets/images/tabs_180px.svg",
		}, {
			Name:  "Text Field",
			Url:   "/text-field",
			Image: "/assets/images/ic_theme_24px.svg",
		}, {
			Name:  "Top App Bar",
			Url:   "/top-app-bar",
			Image: "/assets/images/top_app_bar_180px.svg",
		}, {
			Name:  "Typography",
			Url:   "/typography",
			Image: "/assets/images/fonts_180px.svg",
		},
	}

	sidebar := components.NewComponentSidebar(componentList)
	appbar := components.NewHeaderBar(sidebar)

	a := &app.A{
		RootMarkup: vecty.Markup(
			vecty.Class("demo-panel"),
		),
		ChildMarkup: vecty.Markup(
			vecty.Class("demo-content"),
		),
		Appbar:  appbar,
		Sidebar: sidebar,
		Routes: []vecty.ComponentOrHTML{
			router.NewRoute(
				"/", views.NewComponentImageList(componentList),
				router.NewRouteOpts{ExactMatch: true},
			),
			router.NewRoute(
				"/button", views.NewButtonPage(), router.NewRouteOpts{ExactMatch: true},
			),
			router.NewRoute(
				"/menu", views.NewMenuPage(), router.NewRouteOpts{ExactMatch: true},
			),
			router.NewRoute(
				"/list", views.NewListPage(), router.NewRouteOpts{ExactMatch: true},
			),
			router.NewRoute(
				"/checkbox", views.NewCheckboxPage(), router.NewRouteOpts{ExactMatch: true},
			),
			router.NewRoute(
				"/tabs", views.NewTabbarPage(), router.NewRouteOpts{ExactMatch: true},
			),
			router.NewRoute(
				"/data-table", views.NewDatatablePage(), router.NewRouteOpts{ExactMatch: true},
			),
			router.NewRoute(
				"/snackbar", views.NewSnackbarPage(), router.NewRouteOpts{ExactMatch: true},
			),
		},
	}

	return a
}
