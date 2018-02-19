// https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/vecty-material/dialog"

import (
	"math/rand"

	mdcD "agamigo.io/material/dialog"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

const (
	MDCClass = "mdc-dialog"
)

// D is a material dialog component. It should only be created using the New
// function.
type D struct {
	*mdcD.D
	vecty.Core
	id            string
	Role          string
	Label         string
	Description   string
	AcceptBtn     string
	CancelBtn     string
	AcceptHandler *vecty.EventListener
	CancelHandler *vecty.EventListener
}

func New() (c *D) {
	c = &D{}
	c.D = &mdcD.D{}
	c.Role = "dialog"
	c.AcceptBtn = "Accept"
	c.CancelBtn = "Deny"
	return c
}

// Render implements the vecty.Component interface.
func (c *D) Render() vecty.ComponentOrHTML {
	return elem.Aside(
		vecty.Markup(
			prop.ID(c.String()),
			vecty.Class(MDCClass),
			vecty.Attribute("role", c.Role),
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
						prop.ID(c.labelID()),
					),
					vecty.Text(c.Label),
				),
			),
			elem.Section(
				vecty.Markup(
					prop.ID(c.descriptionID()),
					vecty.Class("mdc-dialog__body"),
				),
				vecty.Text(c.Description),
			),
			elem.Footer(
				vecty.Markup(
					vecty.Class("mdc-dialog__footer"),
				),
				elem.Button(
					vecty.Markup(
						vecty.Class("mdc-button"),
						vecty.Class("mdc-dialog__footer__button"),
						vecty.Class("mdc-dialog__footer__button--cancel"),
						prop.Value(c.CancelBtn),
						vecty.MarkupIf(c.CancelHandler != nil,
							c.CancelHandler),
					),
					vecty.Text(c.CancelBtn),
				),
				elem.Button(
					vecty.Markup(
						prop.Type(prop.TypeButton),
						vecty.Class("mdc-button"),
						vecty.Class("mdc-dialog__footer__button"),
						vecty.Class("mdc-dialog__footer__button--accept"),
						prop.Value(c.AcceptBtn),
						vecty.MarkupIf(c.AcceptHandler != nil,
							c.AcceptHandler),
					),
					vecty.Text(c.AcceptBtn),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-dialog__backdrop"),
			),
		),
	)
}

// Mount implements the vecty.Mounter interface and calls Start() on the
// underlying material dialog component.
func (c *D) Mount() {
	c.Start(js.Global.Get("document").Call("getElementById", c.String()))
}

// Unmount implements the vecty.Unmounter interface and calls Stop() on the
// underlying material dialog component.
func (c *D) Unmount() {
	c.Stop()
}

// OpenHandler opens the dialog. It fits the signature expected by the
// vecty/event package for creating a vecty.EventListener.
func (c *D) OpenHandler(e *vecty.Event) {
	err := c.D.Open()
	if err != nil {
		println("[ERROR] Unable to open dialog: %v", err)
	}
}

func (c *D) String() string {
	if c.id == "" {
		rand.Seed(13)
		c.id = c.Component().Type.MDCClassName + string(rand.Int())
	}
	return c.id
}

func (c *D) labelID() string {
	return c.String() + "-label"
}

func (c *D) ariaLabelledBy() vecty.Applyer {
	return vecty.Attribute("aria-labelledby", c.labelID())
}

func (c *D) descriptionID() string {
	return c.String() + "-description"
}

func (c *D) ariaDescribedBy() vecty.Applyer {
	return vecty.Attribute("aria-describedby", c.descriptionID())
}

func (c *D) headerID() vecty.Applyer {
	return prop.ID(c.String() + "-header")
}
