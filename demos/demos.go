package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type NotFoundView struct {
	vecty.Core
}

func (nf *NotFoundView) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading1(
			vecty.Text("page not found"),
		),
	)
}

type CatalogView struct {
	vecty.Core
}

func (c *CatalogView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Material Components Web | Catalog")

	return elem.Div(
		&HeaderBar{},
	)
}
