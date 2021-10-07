package snackbar

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/button"
	"github.com/vecty-components/material/gojs"
)

// S is a vecty-material snackbar component.
type S struct {
	*base.MDC
	vecty.Core
	Root    vecty.MarkupOrChild `vecty:"prop"`
	Label   vecty.MarkupOrChild `vecty:"prop"`
	Buttons []*button.B         `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (c *S) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	buttons := make([]vecty.MarkupOrChild, len(c.Buttons))
	for i, b := range c.Buttons {
		markup := base.ExtractMarkupFromLink(b.Label.(*vecty.HTML))
		buttons[i] = elem.Button(
			vecty.Markup(
				prop.Type("button"),
				vecty.Class("mdc-button", "mdc-snackbar__action"),
				vecty.MarkupIf(
					markup.OnClick != nil,
					event.Click(markup.OnClick).PreventDefault(),
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-button__ripple"),
				),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class("mdc-button__label"),
				),
				markup.Child,
			),
		)
	}

	// Built-in root element
	return elem.Aside(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-snackbar__surface"),
				vecty.Attribute("role", "status"),
				vecty.Attribute("aria-relevant", "additions"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-snackbar__label"),
					vecty.Attribute("aria-atomic", "false"),
				),
				c.Label,
			),
			elem.Div(
				append([]vecty.MarkupOrChild{
					vecty.Markup(
						vecty.Class("mdc-snackbar__actions"),
						vecty.Attribute("aria-atomic", "true"),
					),
				}, buttons...)...,
			),
		),
	)
}

func (c *S) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCSnackbar",
				MDCCamelCaseName: "snackbar",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{
			// "timeoutMs":     c.Timeout,
			// "closeOnEscape": c.CloseOnEscape,
		})
	}

	vecty.Markup(
		vecty.Class("mdc-snackbar"),
	).Apply(h)
	c.MDC.RootElement = h
}

// Open displays the snackbar. If the configuration is invalid an error message
// will be returned and the snackbar will not be shown. For information on
// config requirements look at documentation for S.
func (c *S) Open() error {
	var err error
	gojs.CatchException(&err)

	c.MDC.Component.Component().Value.Call("open")
	return err
}
