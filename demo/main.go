package main

import (
	"agamigo.io/material/vecty/checkbox"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.SetTitle("Material Components Go")
	vecty.RenderBody(&PageView{})
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("node_modules/material-components-web/dist/material-components-web.css")
	cb := checkbox.New()
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-form-field"),
			),
			cb,
			elem.Label(
				vecty.Markup(
					prop.For("native-js-checkbox"),
				),
				vecty.Text("Default checkbox"),
			),
		),
	)
}
