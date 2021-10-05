package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"

	"github.com/vecty-material/material/button"
	ssnackbar "github.com/vecty-material/material/components/snackbar"
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
	)
}

type SnackbarDemos struct {
	vecty.Core
}

func (bd *SnackbarDemos) Render() vecty.ComponentOrHTML {
	s := &snackbar.S{
		Label: vecty.Text("Can't send photo. Retry in 5 seconds."),
	}

	return elem.Div(
		s,
		&button.B{
			Root: vecty.Markup(
				vecty.Class("hero-button"),
			),
			Label: elem.Anchor(
				vecty.Markup(
					event.Click(func(e *vecty.Event) {
						s.MDC.Component.(*ssnackbar.S).Open()
					}),
				),
				vecty.Text("Baseline"),
			),
		},
	)

}
