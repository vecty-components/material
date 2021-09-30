package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"

	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/demos/views"
)

func main() {
	base.SetViewport()
	base.AddResources()

	body := &Body{}
	vecty.RenderBody(body)
}

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/App.css")

	return elem.Body(&CatalogPage{})
}

type CatalogPage struct {
	vecty.Core
}

func (c *CatalogPage) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Material Components Web | Catalog")

	vecty.AddStylesheet("/assets/styles/CatalogPage.css")

	return elem.Div(
		components.NewHeaderBar(),
		router.NewRoute(
			"/button", views.NewButtonPage(), router.NewRouteOpts{ExactMatch: true},
		),
		router.NewRoute("/", views.NewComponentImageList(), router.NewRouteOpts{ExactMatch: true}),
		router.NotFoundHandler(views.NewComponentImageList()),
	)
}
