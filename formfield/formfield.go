package formfield

import (
	mbase "agamigo.io/material/base"
	"agamigo.io/material/formfield"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// FF is a vecty-material formfield component.
type FF struct {
	*base.Base
	*State
}

type State struct {
	*formfield.FF
	inputID  string
	Input    vecty.ComponentOrHTML
	Label    string
	AlignEnd bool
}

func New(p *base.Props, s *State) *FF {
	c := &FF{}
	if s == nil {
		s = &State{}
	}
	if s.FF == nil {
		s.FF = formfield.New()
	}
	c.State = s
	c.Base = base.New(p, c)
	return c
}

// Render implements the vecty.Component interface.
func (c *FF) Render() vecty.ComponentOrHTML {
	input := c.Input
	if c.Input != nil {
		switch t := c.Input.(type) {
		case *base.Base:
			c.inputID = t.ID
		case *vecty.HTML:
			input = base.RenderStoredChild(c.Input)
		}
		switch t := c.Input.(type) {
		case mbase.Componenter:
			c.FF.Input = t.Component()
		}
	}
	return c.Base.Render(elem.Div(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.MarkupIf(c.Props.ID != "",
				prop.ID(c.Props.ID)),
			vecty.Class("mdc-form-field"),
			vecty.MarkupIf(c.AlignEnd,
				vecty.Class("mdc-form-field--align-end"),
			),
		),
		input,
		elem.Label(
			vecty.Markup(
				vecty.MarkupIf(c.inputID != "",
					prop.For(c.inputID),
				),
			),
			vecty.Text(c.Label),
		),
	))
}
