// https://material.io/components/web/catalog/buttons/
package button // import "agamigo.io/vecty-material/button"

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// B is a vecty-material button component.
type B struct {
	vecty.Core
	Label        vecty.ComponentOrHTML
	ClickHandler func(*vecty.Event)
	Disabled     bool
	Raised       bool
	Unelevated   bool
	Stroked      bool
	Dense        bool
	Compact      bool
	Classes      vecty.ClassMap
	Icon         string
	IconClass    string
	Href         string
}

// Render implements the vecty.Component interface.
func (c *B) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class("mdc-button"),
			c.Classes,
			prop.Type(prop.TypeButton),
			vecty.MarkupIf(c.ClickHandler != nil,
				event.Click(c.ClickHandler),
			),
			vecty.MarkupIf(c.Disabled,
				vecty.Property("disabled", true),
			),
			vecty.MarkupIf(c.Raised,
				vecty.Class("mdc-button--raised"),
			),
			vecty.MarkupIf(c.Unelevated,
				vecty.Class("mdc-button--unelevated"),
			),
			vecty.MarkupIf(c.Stroked,
				vecty.Class("mdc-button--stroked"),
			),
			vecty.MarkupIf(c.Dense,
				vecty.Class("mdc-button--dense"),
			),
			vecty.MarkupIf(c.Compact,
				vecty.Class("mdc-button--compact"),
			),
		),
		vecty.If(c.Icon != "",
			elem.Italic(
				vecty.Markup(
					vecty.Class("mdc-button__icon"),
					vecty.MarkupIf(c.IconClass != "",
						vecty.Class(c.IconClass),
					),
				),
				vecty.Text(c.Icon),
			),
		),
		c.Label,
	)
}
