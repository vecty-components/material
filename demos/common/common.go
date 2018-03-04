package common

import (
	"path"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type ToolbarHeader struct {
	vecty.Core
	Title string
}

func main() {
}

func (c *ToolbarHeader) Render() vecty.ComponentOrHTML {
	pathname := js.Global.Get("window").Get("location").Get("pathname").String()
	pathname = path.Base(pathname)
	var toolbarNav *vecty.HTML
	if pathname == "/" || pathname == "." || pathname == "demos" {
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
	} else {
		toolbarNav = elem.Anchor(
			vecty.Markup(
				prop.Href("../"),
				vecty.Class("catalog-back"),
				vecty.Class("mdc-toolbar__menu-icon"),
			),
			elem.Italic(
				vecty.Markup(
					vecty.Class("material-icons"),
				),
				// vecty.Text("&#xE5C4;"),
				vecty.Text("arrow_back"),
			),
		)
	}
	return elem.Header(
		vecty.Markup(
			vecty.Class("mdc-toolbar"),
			vecty.Class("mdc-toolbar--fixed"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-toolbar__row"),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class("mdc-toolbar__section"),
					vecty.Class("mdc-toolbar__section--align-start"),
				),
				toolbarNav,
				elem.Span(
					vecty.Markup(
						vecty.Class("mdc-toolbar__title"),
						vecty.Class("catalog-title"),
					),
					vecty.Text(c.Title),
				),
			),
		),
	)
}
