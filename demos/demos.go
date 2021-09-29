package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

type CatalogPage struct {
	vecty.Core
}

func (c *CatalogPage) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Material Components Web | Catalog")

	vecty.AddStylesheet("/assets/styles/CatalogPage.css")

	return elem.Div(
		&HeaderBar{},
		router.NewRoute("/", &ComponentImageList{
			images: make(map[string]string),
		}, router.NewRouteOpts{ExactMatch: true}),
		router.NotFoundHandler(&ComponentImageList{
			images: make(map[string]string),
		}),
	)
}
