// https://material.io/components/web/catalog/buttons/
package button // import "github.com/vecty-material/material/button"

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/base"
)

// B is a vecty-material button component.
type B struct {
	*base.MDC
	vecty.Core
	Root       vecty.MarkupOrChild
	Label      vecty.ComponentOrHTML
	Icon       vecty.ComponentOrHTML
	Disabled   bool
	Raised     bool
	Unelevated bool
	Outlined   bool
	Dense      bool

	link *base.LinkMarkup
}

// Render implements the vecty.Component interface.
func (c *B) Render() vecty.ComponentOrHTML {
	c.link = base.ExtractMarkupFromLink(
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

	label := c.Label
	if c.link.Href != "" && c.link.Child != nil {
		label = c.link.Child
	}

	return elem.Button(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		ico,
		base.RenderStoredChild(label),
	)
}

func (c *B) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
	}
	c.MDC.Component = nil
	c.MDC.RootElement = h
	vecty.Markup(
		vecty.Class("mdc-button"),
		prop.Type(prop.TypeButton),
		vecty.MarkupIf(c.link.OnClick != nil && !c.link.PreventDefault,
			event.Click(c.link.OnClick),
		),
		vecty.MarkupIf(c.link.OnClick != nil && c.link.PreventDefault,
			event.Click(c.link.OnClick).PreventDefault(),
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
}
