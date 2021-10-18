// https://material.io/components/web/catalog/buttons/
package button // import "github.com/vecty-components/material/button"

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
)

// B is a vecty-material button component.
type B struct {
	*base.MDC
	vecty.Core
	Root       vecty.MarkupOrChild   `vecty:"prop"`
	Label      vecty.ComponentOrHTML `vecty:"prop"`
	Icon       vecty.ComponentOrHTML `vecty:"prop"`
	Disabled   bool                  `vecty:"prop"`
	Raised     bool                  `vecty:"prop"`
	Unelevated bool                  `vecty:"prop"`
	Outlined   bool                  `vecty:"prop"`
	Dense      bool                  `vecty:"prop"`

	markup *base.LinkMarkup
}

// Render implements the vecty.Component interface.
func (c *B) Render() vecty.ComponentOrHTML {
	c.markup = base.ExtractMarkupFromLink(
		c.Label.(*vecty.HTML),
	)

	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	var ico *vecty.HTML
	switch t := c.Icon.(type) {
	case nil:
		ico = nil
	case vecty.Component:
		ico = t.Render().(*vecty.HTML)
	case *vecty.HTML:
		ico = t
	}
	if ico != nil {
		vecty.Class("mdc-button__icon").Apply(ico)
	}

	return elem.Button(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class("mdc-button__ripple"),
			),
		),
		ico,
		base.RenderStoredChild(c.markup.Child),
	)
}

func (c *B) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCRipple",
				MDCCamelCaseName: "ripple",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{})
	}

	vecty.Markup(
		vecty.Class("mdc-button"),
		prop.Type(prop.TypeButton),
		vecty.MarkupIf(c.markup.OnClick != nil,
			event.Click(c.markup.OnClick).PreventDefault(),
		),
		vecty.Property("disabled", c.Disabled),
		vecty.MarkupIf(c.Raised,
			vecty.Class("mdc-button--raised"),
		),
		vecty.MarkupIf(c.Unelevated,
			vecty.Class("mdc-button--unelevated"),
		),
		vecty.MarkupIf(c.Outlined,
			vecty.Class("mdc-button--outlined"),
		),
		vecty.MarkupIf(c.Dense,
			vecty.Class("mdc-button--dense"),
		),
	).Apply(h)
	c.MDC.RootElement = h
}
