package components

import (
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/icon"
	"github.com/vecty-material/material/iconbutton"
)

type HeaderBar struct {
	vecty.Core
	ticonb, miconb *iconbutton.IB
	sidebar        *ComponentSidebar
}

func pathname() string {
	return js.Global().Get("window").Get("location").
		Get("pathname").String()
}

func NewHeaderBar(sidebar *ComponentSidebar) *HeaderBar {
	return &HeaderBar{sidebar: sidebar}
}

func (hb *HeaderBar) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/HeaderBar.css")

	if hb.ticonb == nil || hb.miconb == nil {
		hb.ticonb = &iconbutton.IB{
			Root: vecty.Markup(
				vecty.Class("mdc-top-app-bar__navigation-icon"),
			),
			OnIcon: elem.Image(
				vecty.Markup(
					prop.Src("/assets/images/ic_component_24px_white.svg"),
					prop.Alt("Material logo"),
				),
			),
		}

		hb.miconb = &iconbutton.IB{
			Root: vecty.Markup(
				vecty.Class("mdc-top-app-bar__navigation-icon"),
			),
			OnIcon: &icon.I{
				Name: "menu",
			},
			OnClick: func(e *vecty.Event) {
				hb.sidebar.Toggle()
			},
		}

	}

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
				vecty.If(pathname() == "/", hb.ticonb),
				vecty.If(pathname() != "/", hb.miconb),
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
