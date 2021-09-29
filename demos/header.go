package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

/*
	TODO: move this into a separate component
*/

type HeaderIcon struct {
	vecty.Core
}

func (hi *HeaderIcon) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class(
				"mdc-icon-button", "material-icons", "mdc-top-app-bar__navigation-icon",
			),
			vecty.Attribute("title", "home"),
		),
		elem.Image(),
	)
}

type HeaderBar struct {
	vecty.Core
}

func (hb *HeaderBar) Render() vecty.ComponentOrHTML {
	return elem.Header(
		vecty.Markup(
			vecty.Class("mdc-top-app-bar", "catalog-top-app-bar"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-top-app-bar__row"),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class(
						"mdc-top-app-bar__section", "mdc-top-app-bar__section--align-start",
					),
				),
				/* HeaderIcon */

				elem.Span(
					vecty.Markup(
						vecty.Class(
							"mdc-top-app-bar__title", "catalog-top-app-bar__title",
						),
					),

					elem.Span(
						vecty.Markup(
							vecty.Class(
								"catalog-top-app-bar__title--small-screen",
							),
						),
						vecty.Text("MDC Web"),
					),
					elem.Span(
						vecty.Markup(
							vecty.Class(
								"catalog-top-app-bar__title--large-screen",
							),
						),
						vecty.Text("Material Components for the Web"),
					),
				),
			),
		),
	)
}
