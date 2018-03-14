// https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/vecty-material/dialog"

import (
	"agamigo.io/material/dialog"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/button"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// D is a material dialog component.
type D struct {
	*base.Base
	*State
}

type State struct {
	*dialog.D
	Header        string
	Body          vecty.ComponentOrHTML
	Role          string
	Open          bool `js:"open"`
	NoBackdrop    bool
	Scrollable    bool
	AcceptBtn     *button.B
	CancelBtn     *button.B
	AcceptHandler func(this *D, e *vecty.Event)
	CancelHandler func(this *D, e *vecty.Event)
}

func New(p *base.Props, s *State) (c *D) {
	open := js.InternalObject(s).Get("Open").Bool()
	c = &D{}
	if s == nil {
		s = &State{}
	}
	c.State = s
	if c.D == nil {
		c.D = dialog.New()
	}
	c.Base = base.New(p, c)
	c.Open = open
	return c
}

// Render implements the vecty.Component interface.
func (c *D) Render() vecty.ComponentOrHTML {
	cancelButton := c.CancelBtn
	if cancelButton == nil {
		cancelButton = button.New(nil, nil)
	}
	cancelButton.Label = vecty.Text("Cancel")
	cancelButton.Props.Markup = append(cancelButton.Props.Markup,
		vecty.Class("mdc-dialog__footer__button"),
		vecty.Class("mdc-dialog__footer__button--cancel"),
	)
	if c.CancelHandler != nil {
		cancelButton.Props.Markup = append(cancelButton.Props.Markup,
			event.Click(c.wrapCancelHandler()))
	} else {
		cancelButton.Props.Markup = append(cancelButton.Props.Markup,
			event.Click(func(e *vecty.Event) {
			}).StopPropagation())
	}

	acceptButton := c.AcceptBtn
	if acceptButton == nil {
		acceptButton = button.New(nil, nil)
	}
	acceptButton.Label = vecty.Text("Accept")
	acceptButton.Props.Markup = append(acceptButton.Props.Markup,
		vecty.Class("mdc-dialog__footer__button"),
		vecty.Class("mdc-dialog__footer__button--accept"),
	)
	if c.AcceptHandler != nil {
		acceptButton.Props.Markup = append(acceptButton.Props.Markup,
			event.Click(c.wrapAcceptHandler()))
	} else {
		acceptButton.Props.Markup = append(acceptButton.Props.Markup,
			event.Click(func(e *vecty.Event) {
			}).StopPropagation())
	}

	return c.Base.Render(elem.Aside(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.MarkupIf(c.Props.ID != "", prop.ID(c.Props.ID)),
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
						vecty.MarkupIf(c.Props.ID != "",
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
	))
}

func (c *D) labelID() string {
	if c.Props.ID == "" {
		return ""
	}
	return c.Props.ID + "-label"
}

func (c *D) ariaLabelledBy() vecty.Applyer {
	if c.labelID() == "" {
		return nil
	}
	return vecty.Attribute("aria-labelledby", c.labelID())
}

func (c *D) descriptionID() string {
	if c.Props.ID == "" {
		return ""
	}
	return c.Props.ID + "-description"
}

func (c *D) ariaDescribedBy() vecty.Applyer {
	if c.descriptionID() == "" {
		return nil
	}
	return vecty.Attribute("aria-describedby", c.descriptionID())
}

func (c *D) headerID() string {
	if c.Props.ID == "" {
		return ""
	}
	return c.Props.ID + "-header"
}

func (c *D) wrapCancelHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.CancelHandler(c, e)
	}
}

func (c *D) wrapAcceptHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.AcceptHandler(c, e)
	}
}
