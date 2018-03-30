// https://material.io/components/web/catalog/buttons/
package button // import "agamigo.io/vecty-material/button"

import (
	"agamigo.io/material/ripple"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// B is a vecty-material button component.
type B struct {
	vecty.Core
	ID          string
	Markup      []vecty.Applyer
	rootElement *vecty.HTML
	Ripple      bool
	ripple      *ripple.R
	Label       vecty.ComponentOrHTML
	Icon        vecty.ComponentOrHTML
	Disabled    bool
	Raised      bool
	Unelevated  bool
	Stroked     bool
	Dense       bool
	Href        string
	OnClick     func(this *B, e *vecty.Event)
}

// Render implements the vecty.Component interface.
func (c *B) Render() vecty.ComponentOrHTML {
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

	c.rootElement = elem.Button(
		vecty.Markup(
			vecty.Markup(c.Markup...),
			vecty.MarkupIf(c.ID != "",
				prop.ID(c.ID)),
			vecty.Class("mdc-button"),
			prop.Type(prop.TypeButton),
			vecty.MarkupIf(c.OnClick != nil,
				event.Click(c.wrapOnClick()),
			),
			vecty.Property("disabled", c.Disabled),
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
		),
		vecty.If(ico != nil, ico),
		vecty.If(c.Label != nil, base.RenderStoredChild(c.Label)),
	)
	return c.rootElement
}

func (c *B) MDCRoot() *base.Base {
	return &base.Base{
		MDC:       nil,
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		RippleC:   c.ripple,
	}
}

func (c *B) Mount() {
	c.MDCRoot().Mount()
}

func (c *B) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *B) wrapOnClick() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.OnClick(c, e)
	}
}
