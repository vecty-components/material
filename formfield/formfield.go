package formfield

import (
	"agamigo.io/material/formfield"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/base/applyer"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// FF is a vecty-material formfield component.
type FF struct {
	*base.MDCRoot
	vecty.Core
	Root     vecty.MarkupOrChild
	Input    vecty.ComponentOrHTML
	Label    string
	AlignEnd bool
}

// Render implements the vecty.Component interface.
func (c *FF) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	inputID := applyer.FindID(c.Input)
	return elem.Div(
		vecty.Markup(
			c,
			vecty.MarkupIf(rootMarkup != nil, *rootMarkup),
		),
		c.Input,
		elem.Label(
			vecty.Markup(
				vecty.MarkupIf(inputID != "",
					prop.For(inputID),
				),
			),
			vecty.Text(c.Label),
		),
	)
}

func (c *FF) Apply(h *vecty.HTML) {
	switch {
	case c.MDCRoot == nil:
		c.MDCRoot = &base.MDCRoot{}
		fallthrough
	case c.MDCRoot.MDC == nil:
		c.MDCRoot.MDC = formfield.New()
	}
	vecty.Markup(
		vecty.Class("mdc-form-field"),
		vecty.MarkupIf(c.AlignEnd,
			vecty.Class("mdc-form-field--align-end"),
		),
	).Apply(h)
	c.MDCRoot.Element = h
}
