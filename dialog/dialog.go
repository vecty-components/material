// https://material.io/components/web/catalog/dialogs/
package dialog // import "github.com/vecty-components/material/dialog"

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/base/applyer"
	"github.com/vecty-components/material/button"
	"github.com/vecty-components/material/components/dialog"
)

// D is a material dialog component.
type D struct {
	*base.MDC
	vecty.Core
	Root       vecty.MarkupOrChild
	Header     string
	Body       vecty.ComponentOrHTML
	Role       string
	Open       bool
	NoBackdrop bool
	Scrollable bool
	AcceptBtn  *button.B
	CancelBtn  *button.B
	OnAccept   func(e *vecty.Event)
	OnCancel   func(e *vecty.Event)
}

// Render implements the vecty.Component interface.
func (c *D) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Aside(c.Root)
	}

	// TODO: Make action buttons a type
	if c.CancelBtn != nil {
		if c.CancelBtn.Label == nil {
			c.CancelBtn.Label = vecty.Text("Cancel")
		}
		if c.CancelBtn.Root == nil {
			c.CancelBtn.Root = vecty.Markup(
				vecty.Class("mdc-dialog__button"),
				vecty.Attribute("data-mdc-dialog-action", "cancel"),
				event.Click(c.onCancel),
			)
		}
	}

	if c.AcceptBtn != nil {
		if c.AcceptBtn.Label == nil {
			c.AcceptBtn.Label = vecty.Text("Accept")
		}
		if c.AcceptBtn.Root == nil {
			c.AcceptBtn.Root = vecty.Markup(
				vecty.Class("mdc-dialog__button"),
				vecty.Attribute("data-mdc-dialog-action", "accept"),
				event.Click(c.onCancel),
			)
		}
	}

	h := elem.Aside(
		vecty.Markup(
			base.MarkupIfNotNil(rootMarkup),
		),
	)

	// Built-in root element.
	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-dialog__container"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-dialog__surface"),
					vecty.MarkupIf(c.Role == "", vecty.Attribute("role", "alertdialog")),
					vecty.MarkupIf(c.Role != "", vecty.Attribute("role", c.Role)),
					vecty.Attribute("aria-modal", "true"),
					vecty.Attribute("aria-labelledby", c.ariaLabelledBy(h)),
					vecty.Attribute("aria-describedby", c.ariaDescribedBy(h)),
				),
				//				elem.Header(
				//					vecty.Markup(
				//						vecty.Class("mdc-dialog__header"),
				//					),
				elem.Heading2(
					vecty.Markup(
						vecty.Class("mdc-dialog__title"),
						vecty.MarkupIf(c.labelID(h) != "",
							prop.ID(c.labelID(h))),
					),
					vecty.Text(c.Header),
				),
				//				),
				elem.Div(
					vecty.Markup(
						prop.ID(c.descriptionID(h)),
						vecty.Class("mdc-dialog__content"),
					),
					base.RenderStoredChild(c.Body),
				),
				vecty.If(
					c.CancelBtn != nil || c.AcceptBtn != nil,
					elem.Footer(
						vecty.Markup(
							vecty.Class("mdc-dialog__actions"),
						),
						vecty.If(
							c.CancelBtn != nil,
							base.RenderStoredChild(c.CancelBtn),
						),
						vecty.If(
							c.AcceptBtn != nil,
							base.RenderStoredChild(c.AcceptBtn),
						),
					),
				),
			),
			vecty.If(!c.NoBackdrop,
				elem.Div(
					vecty.Markup(
						vecty.Class("mdc-dialog__scrim"),
					),
				),
			),
		),
	)
}

func (c *D) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = dialog.New()
	}
	c.MDC.Component.(*dialog.D).Open = c.Open
	vecty.Markup(
		vecty.Class("mdc-dialog"),
		vecty.MarkupIf(c.Open, vecty.Class("mdc-dialog--open")),
		vecty.MarkupIf(!c.Open, vecty.Attribute("aria-hidden", "true")),
	).Apply(h)
	c.MDC.RootElement = h
}

func (c *D) labelID(h *vecty.HTML) string {
	id := applyer.FindID(h)
	if id == "" {
		return ""
	}
	return id + "-label"
}

func (c *D) ariaLabelledBy(h *vecty.HTML) vecty.Applyer {
	if c.labelID(h) == "" {
		return nil
	}
	return vecty.Attribute("aria-labelledby", c.labelID(h))
}

func (c *D) descriptionID(h *vecty.HTML) string {
	id := applyer.FindID(h)
	if id == "" {
		return ""
	}
	return id + "-description"
}

func (c *D) ariaDescribedBy(h *vecty.HTML) vecty.Applyer {
	if c.descriptionID(h) == "" {
		return nil
	}
	return vecty.Attribute("aria-describedby", c.descriptionID(h))
}

func (c *D) onCancel(e *vecty.Event) {
	c.Open = false
	if d, ok := c.MDC.Component.(*dialog.D); ok {
		d.Open = false
	}

	vecty.Rerender(c)
	if c.OnCancel != nil {
		c.OnCancel(e)
	}
}

func (c *D) onAccept(e *vecty.Event) {
	c.Open = false
	if d, ok := c.MDC.Component.(*dialog.D); ok {
		d.Open = false
	}

	vecty.Rerender(c)
	if c.OnAccept != nil {
		c.OnAccept(e)
	}
}
