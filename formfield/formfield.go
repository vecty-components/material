package formfield

import (
	"strconv"

	"agamigo.io/material/formfield"
	"agamigo.io/vecty-material/internal/base"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

var (
	autoID int
)

// FF is a vecty-material formfield component.
type FF struct {
	vecty.Core
	mdc      *formfield.FF
	id       string
	autoID   int
	Input    vecty.ComponentOrHTML
	InputID  string
	Label    string
	AlignEnd bool
}

// Render implements the vecty.Component interface.
func (c *FF) Render() vecty.ComponentOrHTML {
	c.updateInputID()
	return elem.Div(
		vecty.Markup(
			vecty.Class("mdc-form-field"),
			prop.ID(c.ID()),
			vecty.MarkupIf(c.AlignEnd,
				vecty.Class("mdc-form-field--align-end"),
			),
		),
		c.Input,
		elem.Label(
			vecty.Markup(
				vecty.MarkupIf(c.InputID != "",
					prop.For(c.InputID),
				),
			),
			vecty.Text(c.Label),
		),
	)
}

func (c *FF) Mount() {
	if c.mdc == nil {
		c.mdc = &formfield.FF{}
	}
	e := js.Global.Get("document").Call("getElementById", c.ID())
	if e == nil || e == js.Undefined {
		panic("Unable to find element during Mount()")
	}
	err := c.mdc.Start(e)
	if err != nil {
		panic(err)
	}
}

func (c *FF) Unmount() {
	err := c.mdc.Stop()
	if err != nil {
		panic(err)
	}
}

func (c *FF) ID() string {
	if c.id != "" {
		return c.id
	}
	// Set up new component autoID if needed
	if c.autoID == 0 {
		autoID = autoID + 1
		c.autoID = autoID
	}
	// Generate an ID, increment global autoID
	c.id = "vecty-material-formfield-gen-" + strconv.Itoa(c.autoID)
	return c.id
}

func (c *FF) SetID(id string) {
	if id == "" {
		return // Don't allow clearing the ID
	}
	c.id = id
}

func (c *FF) updateInputID() {
	if c.InputID != "" {
		return
	}
	switch t := c.Input.(type) {
	case base.IDer:
		c.InputID = t.ID()
	}
}
