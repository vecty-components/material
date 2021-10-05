package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-material/material/demos/components"
	"github.com/vecty-material/material/snackbar"
)

func NewSnackbarPage() *components.ComponentPage {
	return components.NewComponentPage(
		"Snackbar",
		"Snackbares allow the user to select multiple options from a set.",
		"https://material.io/go/design-checkboxes",
		"https://material.io/components/web/catalog/checkboxes/",
		"https://github.com/material-components/material-components-web/tree/master/packages/mdc-checkbox",
		components.NewHeroComponent(&SnackbarHero{}), &SnackbarDemos{},
	)
}

type SnackbarHero struct {
	vecty.Core
}

func (bh *SnackbarHero) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("tab-content"),
		),
		&snackbar.S{
			Label: vecty.Text("Can't send photo. Retry in 5 seconds."),
			Open:  true,
		},
	)
}

type SnackbarDemos struct {
	vecty.Core
}

func (bd *SnackbarDemos) Render() vecty.ComponentOrHTML {

	return elem.Div()

}
