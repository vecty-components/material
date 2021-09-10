// https://material.io/components/web/catalog/dialogs/
package dialog // import "github.com/vecty-material/material/dialog"

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/base/applyer"
	"github.com/vecty-material/material/button"
	"github.com/vecty-material/material/material/dialog"
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
	OnAccept   func(this *D, e *vecty.Event)
	OnCancel   func(this *D, e *vecty.Event)
}

// Render implements the vecty.Component interface.
func (c *D) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Aside(c.Root)
	}

	// TODO: Make action buttons a type
	cancelButton := c.CancelBtn
	if cancelButton == nil {
		cancelButton = &button.B{}
	}
	if cancelButton.Label == nil {
		cancelButton.Label = vecty.Text("Cancel")
	}
	if cancelButton.Root == nil {
		cancelButton.Root = vecty.Markup(
			vecty.Class("mdc-dialog__footer__button"),
			vecty.Class("mdc-dialog__footer__button--cancel"),
			event.Click(c.onCancel),
		)
	}

	// TODO: Make action buttons a type
	acceptButton := c.AcceptBtn
	if acceptButton == nil {
		acceptButton = &button.B{}
	}
	if acceptButton.Label == nil {
		acceptButton.Label = vecty.Text("Accept")
	}
	if acceptButton.Root == nil {
		acceptButton.Root = vecty.Markup(
			vecty.Class("mdc-dialog__footer__button"),
			vecty.Class("mdc-dialog__footer__button--accept"),
			event.Click(c.onAccept),
		)
	}

	h := elem.Aside(
		vecty.Markup(
			vecty.MarkupIf(rootMarkup != nil, *rootMarkup),
		),
	)

	// Built-in root element.
	return elem.Aside(
		vecty.Markup(
			c,
			vecty.MarkupIf(rootMarkup != nil, *rootMarkup),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-dialog__surface"),
			),
			elem.Header(
				vecty.Markup(
					vecty.Class("mdc-dialog__header"),
				),
				elem.Heading2(
					vecty.Markup(
						vecty.Class("mdc-dialog__header__title"),
						vecty.MarkupIf(c.labelID(h) != "",
							prop.ID(c.labelID(h))),
					),
					vecty.Text(c.Header),
				),
			),
			elem.Section(
				vecty.Markup(
					prop.ID(c.descriptionID(h)),
					vecty.Class("mdc-dialog__body"),
					vecty.MarkupIf(c.Scrollable,
						vecty.Class("mdc-dialog__body--scrollable")),
				),
				base.RenderStoredChild(c.Body),
			),
			elem.Footer(
				vecty.Markup(
					vecty.Class("mdc-dialog__footer"),
				),
				cancelButton,
				acceptButton,
			),
		),
		vecty.If(!c.NoBackdrop,
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-dialog__backdrop"),
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
		vecty.MarkupIf(c.Role == "", vecty.Attribute("role", "dialog")),
		vecty.MarkupIf(c.Role != "", vecty.Attribute("role", c.Role)),
		vecty.MarkupIf(c.Open, vecty.Class("mdc-dialog--open")),
		vecty.MarkupIf(!c.Open, vecty.Attribute("aria-hidden", "true")),
		c.ariaLabelledBy(h),
		c.ariaDescribedBy(h),
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
	if d, ok := c.MDC.Component.(*dialog.D); ok {
		c.Open = d.Open
	}
	if c.OnCancel != nil {
		c.OnCancel(c, e)
	}
}

func (c *D) onAccept(e *vecty.Event) {
	if d, ok := c.MDC.Component.(*dialog.D); ok {
		c.Open = d.Open
	}
	if c.OnAccept != nil {
		c.OnAccept(c, e)
	}
}
