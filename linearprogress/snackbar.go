package linearprogress

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/gojs"
)

// LP is a vecty-material snackbar component.
type LP struct {
	*base.MDC
	vecty.Core
	Root          vecty.MarkupOrChild `vecty:"prop"`
	Indeterminate bool                `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (c *LP) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	// Built-in root element
	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-linear-progress__buffer"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-linear-progress__buffer-bar"),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-linear-progress__buffer-dots"),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class(
					"mdc-linear-progress__bar",
					"mdc-linear-progress__primary-bar",
				),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class("mdc-linear-progress__bar-inner"),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class(
					"mdc-linear-progress__bar",
					"mdc-linear-progress__secondary-bar",
				),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class("mdc-linear-progress__bar-inner"),
				),
			),
		),
	)
}

func (c *LP) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCLinearProgress",
				MDCCamelCaseName: "linearProgress",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{})
	}

	vecty.Markup(
		vecty.Class("mdc-linear-progress"),
		vecty.MarkupIf(
			c.Indeterminate,
			vecty.Class("mdc-linear-progress--indeterminate"),
		),
		vecty.Attribute("aria-label", "Progress Bar"),
		vecty.Attribute("aria-valuemin", "0"),
		vecty.Attribute("aria-valuemax", "1"),
		vecty.Attribute("aria-valuenow", "0"),
	).Apply(h)
	c.MDC.RootElement = h
}

func (c *LP) Open() error {
	var err error
	gojs.CatchException(&err)

	c.MDC.Component.Component().Value.Call("open")
	return err
}

func (c *LP) Close() error {
	var err error
	gojs.CatchException(&err)

	c.MDC.Component.Component().Value.Call("close")
	return err
}
