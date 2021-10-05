package checkbox

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/base/applyer"
	"github.com/vecty-material/material/components/checkbox"
)

// CB is a vecty-material checkbox component.
type CB struct {
	*base.MDC
	vecty.Core
	Root          vecty.MarkupOrChild
	Input         vecty.MarkupOrChild
	Background    vecty.MarkupOrChild
	OnChange      func(e *vecty.Event)
	Checked       bool
	Indeterminate bool
	Disabled      bool
	Value         string
}

// Render implements the vecty.Component interface.
func (c *CB) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	var bg vecty.ComponentOrHTML
	bgMarkup := base.MarkupOnly(c.Background)
	if c.Background != nil && bgMarkup == nil {
		// User supplied background element.
		bg = elem.Div(c.Background)
	} else {
		// Built-in background element.
		bg = elem.Div(
			vecty.Markup(
				vecty.MarkupIf(bgMarkup != nil, bgMarkup),
				vecty.Class("mdc-checkbox__background"),
				vecty.UnsafeHTML(
					`<svg class="mdc-checkbox__checkmark" viewBox="0 0 24 24">
						<path class="mdc-checkbox__checkmark-path"
							fill="none"
							stroke="white"
							d="M1.73,12.91 8.1,19.28 22.79,4.59"/>
					</svg>`,
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-checkbox__mixedmark"),
				),
			),
		)

	}

	input, _ := c.NativeInput()

	// Built-in root element
	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		input,
		bg,
	)
}

func (c *CB) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = checkbox.New()
		if cb, ok := c.MDC.Component.(*checkbox.CB); ok {
			cb.Checked = c.Checked
			cb.Indeterminate = c.Indeterminate
			cb.Disabled = c.Disabled
			cb.Value = c.Value
		}
	}

	vecty.Markup(
		vecty.Class("mdc-checkbox"),
		vecty.MarkupIf(c.Disabled, vecty.Class("mdc-checkbox--disabled")),
	).Apply(h)
	c.MDC.RootElement = h
}

func (c *CB) onChange(e *vecty.Event) {
	if cb, ok := c.MDC.Component.(*checkbox.CB); ok {
		c.Checked = cb.Checked
		c.Indeterminate = cb.Indeterminate
		c.Disabled = cb.Disabled
		c.Value = cb.Value
	}
	if c.OnChange != nil {
		c.OnChange(e)
	}
}

func (c *CB) NativeInput() (element *vecty.HTML, id string) {
	niMarkup := base.MarkupOnly(c.Input)
	if c.Input != nil && niMarkup == nil {
		// User supplied input element.
		element = elem.Input(c.Input)
		id = applyer.FindID(element)
		return
	}

	// Built-in input element.
	element = elem.Input(
		vecty.Markup(
			vecty.MarkupIf(niMarkup != nil, niMarkup),
			event.Change(c.onChange),
			vecty.Class("mdc-checkbox__native-control"),
			prop.Type(prop.TypeCheckbox),
			vecty.MarkupIf(c.Value != "", prop.Value(c.Value)),
			prop.Checked(c.Checked),
			vecty.Property("disabled", c.Disabled),
			vecty.Property("indeterminate", c.Indeterminate),
		),
		c.Input,
	)
	id = applyer.FindID(element)
	return
}
