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
	IsTop   bool
}

func (hi *HeaderIcon) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/HeaderBar.css")

	if hi.IsTop {
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

func NewHeaderBar(sidebar *drawer.D) (*appbar.A, *HeaderIcon) {
	hicon := &HeaderIcon{
		sidebar: sidebar,
	}

	return &appbar.A{
		Root: vecty.Markup(
			vecty.Class("catalog-top-app-bar"),
		),
		SectionStart: vecty.List{
			hicon,
			appbar.Title(
				"Material Components for the Web",
				[]vecty.Applyer{
					vecty.Class(
						"catalog-top-app-bar__title",
					),
				},
			),
		},
	}, hicon
}
