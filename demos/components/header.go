package components

import (
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/appbar"
	"github.com/vecty-material/material/base"
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

	var icon vecty.ComponentOrHTML
	if pathname() == "/" {
		icon = hb.ticonb
	} else {
		icon = hb.miconb
	}

	return &appbar.A{
		Root: vecty.Markup(
			vecty.Class("catalog-top-app-bar"),
		),
		SectionStart: vecty.List{
			base.RenderStoredChild(icon),
			appbar.Title(
				"Material Components for the Web",
				[]vecty.Applyer{
					vecty.Class(
						"catalog-top-app-bar__title",
					),
				},
			),
		},
	}
}
