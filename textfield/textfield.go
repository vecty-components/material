package textfield

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/components/textfield"
)

// https://github.com/material-components/material-components-web/tree/d6db793dfc2d9f3b085274c73c074bbc3a255587/packages/mdc-textfield

// TF is a vecty-material textfield component.
type TF struct {
	*base.MDC
	vecty.Core
	Root     vecty.MarkupOrChild
	Input    string
	Label    string
	AlignEnd bool
}

// Render implements the vecty.Component interface.
func (c *TF) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Input(
			vecty.Markup(
				prop.Type(prop.TypeText),
				vecty.Class("mdc-text-field__input"),
			),
		),
		elem.Label(
			vecty.Markup(
				vecty.Class("mdc-text-field__label"),
				// todo: autogenerate form field id?
				// vecty.MarkupIf(inputID != "", prop.For(inputID)),
			),
			vecty.Text(c.Label),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-text-field__bottom-line"),
			),
		),
	)
}

func (c *TF) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = textfield.New()
	}
	vecty.Markup(
		vecty.Class("mdc-text-field"),
		vecty.MarkupIf(c.AlignEnd,
			vecty.Class("mdc-form-field--align-end"),
		),
	).Apply(h)
	c.MDC.RootElement = h
}
