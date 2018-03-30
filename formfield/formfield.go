package formfield

import (
	"agamigo.io/material/formfield"
	"agamigo.io/material/ripple"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// FF is a vecty-material formfield component.
type FF struct {
	*formfield.FF
	vecty.Core
	ID          string
	Markup      []vecty.Applyer
	rootElement *vecty.HTML
	Ripple      bool
	ripple      *ripple.R
	inputID     string
	Input       vecty.ComponentOrHTML
	Label       string
	AlignEnd    bool
}

// Render implements the vecty.Component interface.
func (c *FF) Render() vecty.ComponentOrHTML {
	c.init()
	input := c.Input
	if c.Input != nil {
		switch t := c.Input.(type) {
		case base.MDCRooter:
			c.inputID = t.MDCRoot().ID
		case *vecty.HTML:
			input = base.RenderStoredChild(c.Input)
		}
	}
	c.rootElement = elem.Div(
		vecty.Markup(
			vecty.Markup(c.Markup...),
			vecty.MarkupIf(c.ID != "",
				prop.ID(c.ID)),
			vecty.Class("mdc-form-field"),
			vecty.MarkupIf(c.AlignEnd,
				vecty.Class("mdc-form-field--align-end"),
			),
		),
		input,
		elem.Label(
			vecty.Markup(
				vecty.MarkupIf(c.inputID != "",
					prop.For(c.inputID),
				),
			),
			vecty.Text(c.Label),
		),
	)
	return c.rootElement
}

func (c *FF) MDCRoot() *base.Base {
	return &base.Base{
		MDC:       c,
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		RippleC:   c.ripple,
	}
}

func (c *FF) Mount() {
	c.MDCRoot().Mount()
}

func (c *FF) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *FF) init() {
	if c.FF == nil {
		c.FF = formfield.New()
	}
	c.FF.Input = c.Input
}
