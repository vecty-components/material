package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

type ComponentImageList struct {
	vecty.Core
}

func (cl *ComponentImageList) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.UnorderedList(
			vecty.Markup(
				prop.ID("catalog-image-list"),
				vecty.Class(
					"mdc-image-list", "standard-image-list",
					"mdc-top-app-bar--fixed-adjust",
				),
			),
		),
	)
}
