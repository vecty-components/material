package formfield

import (
	"agamigo.io/material/formfield"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/checkbox"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// FF is a vecty-material formfield component.
type FF struct {
	*formfield.FF
	vecty.Core
	id      string
	inputID string
	classes vecty.ClassMap
	basic   bool
	started bool
	element *vecty.HTML
	*Config
}

type Config struct {
	Input    vecty.ComponentOrHTML
	Label    string
	AlignEnd bool
}

func New() *FF {
	c := &FF{}
	return c.WithConfig(&Config{}).WithClass("")
}

func (c *FF) WithBasic() *FF {
	if c.started {
		err := c.Stop()
		if err != nil {
			print(err)
		}
	}
	c.basic = true
	return c
}

func (c *FF) WithID(id string) *FF {
	c.id = id
	return c
}

func (c *FF) ID() string {
	return c.id
}

func (c *FF) WithClass(class string) *FF {
	if c.classes == nil {
		c.classes = make(vecty.ClassMap, 1)
	}
	if class != "" {
		c.classes[class] = true
	}
	return c
}

func (c *FF) WithConfig(s *Config) *FF {
	c.FF = formfield.New()
	c.Config = s
	return c
}

// Render implements the vecty.Component interface.
func (c *FF) Render() vecty.ComponentOrHTML {
	if c.Input != nil && c.inputID == "" {
		switch t := c.Input.(type) {
		case base.IDer:
			c.inputID = t.ID()
		}
	}
	c.element = elem.Div(
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
				vecty.MarkupIf(c.inputID != "",
					prop.For(c.inputID),
				),
			),
			vecty.Text(c.Label),
		),
	)
	return c.element
}

func (c *FF) Mount() {
	if c.FF == nil {
		c.FF = formfield.New()
	}
	if c.basic {
		return
	}
	if c.element == nil {
		panic("Element is nil during Mount().")
	}
	e := c.element.Node()
	if e == nil || e == js.Undefined {
		panic("Element is nil during Mount().")
	}
	err := c.Start(e)
	if err != nil {
		panic(err)
	}
	if c.Input != nil {
		switch t := c.Input.(type) {
		case *checkbox.CB:
			c.SetInput(t.Component())
		}
	}
	c.started = true
}

func (c *FF) Unmount() {
	err := c.FF.Stop()
	if err != nil {
		panic(err)
	}
	c.started = false
}
