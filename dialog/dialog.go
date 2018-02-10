// https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/vecty-material/dialog"

import (
	mdcD "agamigo.io/material/dialog"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

const (
	MDCClass  = "mdc-dialog"
	DefaultID = "vecty-dialog"
)

// D is a material dialog component. It should only be created using the New
// function.
type D struct {
	vecty.Core
	*mdcD.D
	Identity    string
	Role        string
	Label       string
	Description string
	AcceptBtn   string
	CancelBtn   string
}

// New creates a new vecty-material dialog component.
func New(id string) (*D, error) {
	c, err := mdcD.New()
	if err != nil {
		return nil, err
	}
	if id == "" {
		id = DefaultID
	}
	return &D{
		D:         c,
		Identity:  id,
		Role:      "dialog",
		AcceptBtn: "Accept",
		CancelBtn: "Deny",
	}, nil
}

// Render implements the vecty.Component interface.
func (c *D) Render() vecty.ComponentOrHTML {
	e := elem.Aside(
		vecty.Markup(
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
						prop.Type(prop.TypeButton),
						vecty.Class("mdc-button"),
						vecty.Class("mdc-dialog__footer__button"),
						vecty.Class("mdc-dialog__footer__button--cancel"),
						prop.Value(c.CancelBtn),
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
	return e
}

// Mount implements the vecty.Mounter interface and calls Start() on the
// underlying material dialog component.
func (c *D) Mount() {
	c.Start()
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

func (c *D) labelID() string {
	return c.Identity + "-label"
}

func (c *D) ariaLabelledBy() vecty.Applyer {
	return vecty.Attribute("aria-labelledby", c.labelID())
}

func (c *D) descriptionID() string {
	return c.Identity + "-description"
}

func (c *D) ariaDescribedBy() vecty.Applyer {
	return vecty.Attribute("aria-describedby", c.descriptionID())
}

func (c *D) headerID() vecty.Applyer {
	return prop.ID(c.Identity + "-header")
}
