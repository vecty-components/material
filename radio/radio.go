package radio

import (
	"agamigo.io/material/radio"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// R is a vecty-material radio component.
type R struct {
	*base.Base
	*State
}

type State struct {
	*radio.R
	ChangeHandler func(*vecty.Event)
	Name          string
	Checked       bool   `js:"checked"`
	Disabled      bool   `js:"disabled"`
	Value         string `js:"value"`
}

func New(p *base.Props, s *State) *R {
	c := &R{}
	if s == nil {
		s = &State{}
	}
	if s.R == nil {
		s.R = radio.New()
	}
	c.State = s
	c.Base = base.New(p, c)
	c.Checked = js.InternalObject(s).Get("Checked").Bool()
	c.Disabled = js.InternalObject(s).Get("Disabled").Bool()
	c.Value = js.InternalObject(s).Get("Value").String()
	return c
}

// Render implements the vecty.Component interface.
func (c *R) Render() vecty.ComponentOrHTML {
	return c.Base.Render(elem.Div(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.Class("mdc-radio"),
			vecty.MarkupIf(c.Disabled,
				vecty.Class("mdc-radio--disabled"),
			),
		),
		elem.Input(
			vecty.Markup(
				vecty.MarkupIf(c.ChangeHandler != nil,
					event.Change(c.ChangeHandler),
				),
				vecty.Class("mdc-radio__native-control"),
				vecty.MarkupIf(c.Props.ID != "",
					prop.ID(c.Props.ID),
				),
				prop.Type(prop.TypeRadio),
				prop.Checked(c.Checked),
				vecty.MarkupIf(c.Value != "",
					prop.Value(c.Value),
				),
				vecty.MarkupIf(c.Name != "", vecty.Property("name", c.Name)),
				vecty.Property("disabled", c.Disabled),
			),
		),
		elem.Div(
			vecty.Markup(vecty.Class("mdc-radio__background")),
			elem.Div(vecty.Markup(vecty.Class("mdc-radio__outer-circle"))),
			elem.Div(vecty.Markup(vecty.Class("mdc-radio__inner-circle"))),
		),
	))
}
