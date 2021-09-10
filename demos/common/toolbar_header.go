package common

import (
	"path"

	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/icon"
	"github.com/vecty-material/material/toolbar"
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
	pathname := js.Global().Get("window").Get("location").Get("pathname").String()
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
			&icon.I{Name: "&#xE5C4;"},
		)
	case NavMenu:
		toolbarNav = elem.Button(
			vecty.Markup(
				vecty.Class("mdc-toolbar__menu-icon"),
				vecty.Class("material-icons"),
				vecty.Class("demo-menu"),
				event.Click(c.MenuHandler),
			),
			vecty.Text("menu"),
		)
	case NavNone:
		toolbarNav = nil
	}
	t := &toolbar.T{
		Root: vecty.Markup(
			vecty.MarkupIf(
				c.NoFixed,
				vecty.Class("mdc-elevation--z4"),
			),
		),
		Fixed: !c.NoFixed,
		SectionStart: vecty.List{
			toolbarNav,
			toolbar.Title(
				c.Title,
				[]vecty.Applyer{vecty.Class("catalog-title")},
			),
		},
	}
	return t.Render()
}
