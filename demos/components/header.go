package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/appbar"
	"github.com/vecty-material/material/drawer"
	"github.com/vecty-material/material/icon"
	"github.com/vecty-material/material/iconbutton"
)

type HeaderIcon struct {
	vecty.Core
	sidebar *drawer.D
}

func (hi *HeaderIcon) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/HeaderBar.css")

	if false {
		return &iconbutton.IB{
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
	}

	return &iconbutton.IB{
		Root: vecty.Markup(
			vecty.Class("mdc-top-app-bar__navigation-icon"),
		),
		OnIcon: &icon.I{
			Name: "menu",
		},
		OnClick: func(e *vecty.Event) {
			hi.sidebar.Open = !hi.sidebar.Open
			vecty.Rerender(hi.sidebar)
		},
	}
}

type HeaderBar struct {
	vecty.Core
	HeaderIcon *HeaderIcon
}

func NewHeaderBar(sidebar *drawer.D) *appbar.A {
	return &appbar.A{
		Root: vecty.Markup(
			vecty.Class("catalog-top-app-bar"),
		),
		SectionStart: vecty.List{
			&HeaderIcon{sidebar: sidebar},
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
