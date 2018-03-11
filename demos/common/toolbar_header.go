package common

import (
	"path"

	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/toolbar"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type NavType int

const (
	NavBack NavType = iota
	NavRoot
	NavMenu
	NavNone
)

type ToolbarHeader struct {
	vecty.Core
	Title       string
	Navigation  NavType
	NoFixed     bool
	MenuHandler func(e *vecty.Event)
}

func (c *ToolbarHeader) Render() vecty.ComponentOrHTML {
	pathname := js.Global.Get("window").Get("location").Get("pathname").String()
	var toolbarNav vecty.ComponentOrHTML
	switch c.Navigation {
	case NavRoot:
		toolbarNav = elem.Span(
			vecty.Markup(
				vecty.Class("catalog-logo"),
				vecty.Class("mdc-toolbar__menu-icon"),
			),
			elem.Image(
				vecty.Markup(
					prop.Src("https://material-components-web.appspot.com/images/ic_component_24px_white.svg"),
				),
			),
		)
	case NavBack:
		toolbarNav = elem.Anchor(
			vecty.Markup(
				prop.Href(path.Clean(pathname+"/..")),
				vecty.Class("catalog-back"),
				vecty.Class("mdc-toolbar__menu-icon"),
			),
			elem.Italic(
				vecty.Markup(
					vecty.Class("material-icons"),
					vecty.UnsafeHTML("&#xE5C4;"),
				),
			),
		)
	case NavMenu:
		toolbarNav = button.New(
			&base.Props{
				Markup: []vecty.Applyer{
					vecty.Class("demo-menu"),
					vecty.Class("mdc-toolbar__menu-icon"),
				},
			},
			&button.State{
				Icon:         "menu",
				ClickHandler: c.MenuHandler,
			},
		)
	case NavNone:
		toolbarNav = nil
	}
	return toolbar.New(
		&base.Props{
			Markup: []vecty.Applyer{
				vecty.MarkupIf(c.NoFixed, vecty.Class("mdc-elevation--z4")),
			},
		},
		&toolbar.State{
			Fixed: !c.NoFixed,
			SectionStart: vecty.List{
				toolbarNav,
				toolbar.Title(
					c.Title,
					&base.Props{
						Markup: []vecty.Applyer{vecty.Class("catalog-title")},
					},
				),
			},
		},
	)
}
