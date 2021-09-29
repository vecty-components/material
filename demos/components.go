package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type ComponentImageList struct {
	vecty.Core
}

func (cl *ComponentImageList) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.UnorderedList(
			vecty.Markup(
				vecty.Class(
					"mdc-image-list", "standard-image-list",
					"mdc-top-app-bar--fixed-adjust",
				),
				vecty.Attribute("id", "catalog-image-list"),
			),
		),
	)
}
