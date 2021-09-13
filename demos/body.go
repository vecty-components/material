package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"

	dbutton "github.com/vecty-material/material/demos/button"
	dcheckbox "github.com/vecty-material/material/demos/checkbox"
	ddialog "github.com/vecty-material/material/demos/dialog"

	router "marwan.io/vecty-router"
)

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			router.NewRoute("/", &DemosCatalogView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/button", &dbutton.ButtonDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/checkbox", dcheckbox.NewCheckboxDemoView(), router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/dialog", ddialog.NewDialogDemoView(), router.NewRouteOpts{ExactMatch: true}),
			router.NotFoundHandler(&notFound{}),
		),
	)
}

type notFound struct {
	vecty.Core
}

func (nf *notFound) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("home-view")),
		elem.Div(
			vecty.Markup(prop.ID("home-top")),
			elem.Heading1(
				vecty.Text("page not found"),
			),
		),
	)
}
