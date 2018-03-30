// https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/vecty-material/dialog"

import (
	"agamigo.io/material/dialog"
	"agamigo.io/material/ripple"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/button"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// D is a material dialog component.
type D struct {
	*dialog.D
	vecty.Core
	ID          string
	Markup      []vecty.Applyer
	rootElement *vecty.HTML
	Ripple      bool
	Basic       bool
	ripple      *ripple.R
	Header      string
	Body        vecty.ComponentOrHTML
	Role        string
	Open        bool
	NoBackdrop  bool
	Scrollable  bool
	AcceptBtn   *button.B
	CancelBtn   *button.B
	OnAccept    func(this *D, e *vecty.Event)
	OnCancel    func(this *D, e *vecty.Event)
}

// Render implements the vecty.Component interface.
func (c *D) Render() vecty.ComponentOrHTML {
	c.init()
	cancelButton := c.CancelBtn
	if cancelButton == nil {
		cancelButton = &button.B{}
	}
	cancelButton.Label = vecty.Text("Cancel")
	cancelButton.Markup = append(cancelButton.Markup,
		vecty.Class("mdc-dialog__footer__button"),
		vecty.Class("mdc-dialog__footer__button--cancel"),
	)
	if c.OnCancel != nil {
		cancelButton.Markup = append(cancelButton.Markup,
			event.Click(c.wrapCancelHandler()))
	} else {
		cancelButton.Markup = append(cancelButton.Markup,
			event.Click(func(e *vecty.Event) {
			}).StopPropagation())
	}

	acceptButton := c.AcceptBtn
	if acceptButton == nil {
		acceptButton = &button.B{}
	}
	acceptButton.Label = vecty.Text("Accept")
	acceptButton.Markup = append(acceptButton.Markup,
		vecty.Class("mdc-dialog__footer__button"),
		vecty.Class("mdc-dialog__footer__button--accept"),
	)
	if c.OnAccept != nil {
		acceptButton.Markup = append(acceptButton.Markup,
			event.Click(c.wrapAcceptHandler()))
	} else {
		acceptButton.Markup = append(acceptButton.Markup,
			event.Click(func(e *vecty.Event) {
			}).StopPropagation())
	}

	c.rootElement = elem.Aside(
		vecty.Markup(
			vecty.Markup(c.Markup...),
			vecty.MarkupIf(c.ID != "", prop.ID(c.ID)),
			vecty.Class("mdc-dialog"),
			vecty.MarkupIf(c.Role == "", vecty.Attribute("role", "dialog")),
			vecty.MarkupIf(c.Role != "", vecty.Attribute("role", c.Role)),
			vecty.MarkupIf(c.Open, vecty.Class("mdc-dialog--open")),
			vecty.MarkupIf(!c.Open, vecty.Attribute("aria-hidden", "true")),
			c.ariaLabelledBy(),
			c.ariaDescribedBy(),
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
						vecty.MarkupIf(c.ID != "",
							prop.ID(c.labelID())),
					),
					vecty.Text(c.Header),
				),
			),
			elem.Section(
				vecty.Markup(
					prop.ID(c.descriptionID()),
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
	return c.rootElement
}

func (c *D) MDCRoot() *base.Base {
	return &base.Base{
		MDC:       c,
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		Basic:     c.Basic,
		RippleC:   c.ripple,
	}
}

func (c *D) Mount() {
	c.MDCRoot().Mount()
}

func (c *D) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *D) init() {
	if c.D == nil {
		c.D = dialog.New()
	}
	c.D.Open = c.Open
}

func (c *D) labelID() string {
	if c.ID == "" {
		return ""
	}
	return c.ID + "-label"
}

func (c *D) ariaLabelledBy() vecty.Applyer {
	if c.labelID() == "" {
		return nil
	}
	return vecty.Attribute("aria-labelledby", c.labelID())
}

func (c *D) descriptionID() string {
	if c.ID == "" {
		return ""
	}
	return c.ID + "-description"
}

func (c *D) ariaDescribedBy() vecty.Applyer {
	if c.descriptionID() == "" {
		return nil
	}
	return vecty.Attribute("aria-describedby", c.descriptionID())
}

func (c *D) headerID() string {
	if c.ID == "" {
		return ""
	}
	return c.ID + "-header"
}

func (c *D) wrapCancelHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.OnCancel(c, e)
	}
}

func (c *D) wrapAcceptHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.OnAccept(c, e)
	}
}
