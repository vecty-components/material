package checkbox

import (
	"agamigo.io/material/checkbox"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// CB is a vecty-material checkbox component.
type CB struct {
	*base.Base
	*State
}

type State struct {
	*checkbox.CB
	ChangeHandler func(*vecty.Event)
	Checked       bool   `js:"checked"`
	Indeterminate bool   `js:"indeterminate"`
	Disabled      bool   `js:"disabled"`
	Value         string `js:"value"`
}

func New(p *base.Props, s *State) *CB {
	c := &CB{}
	if s == nil {
		s = &State{}
	}
	if s.CB == nil {
		s.CB = checkbox.New()
	}
	c.State = s
	c.Base = base.New(p, c)
	return c
}

// Render implements the vecty.Component interface.
func (c *CB) Render() vecty.ComponentOrHTML {
	return c.Base.Render(elem.Div(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.Class("mdc-checkbox"),
			vecty.MarkupIf(c.Disabled,
				vecty.Class("mdc-checkbox--disabled"),
			),
		),
		elem.Input(
			vecty.Markup(
				vecty.MarkupIf(c.ChangeHandler != nil,
					event.Change(c.ChangeHandler),
				),
				vecty.Class("mdc-checkbox__native-control"),
				vecty.MarkupIf(c.Props.ID != "",
					prop.ID(c.Props.ID),
				),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(c.CB.Checked),
				vecty.MarkupIf(c.Value != "",
					prop.Value(c.Value),
				),
				vecty.Property("disabled", c.Disabled),
				vecty.Property("indeterminate", c.Indeterminate),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-checkbox__background"),
				vecty.UnsafeHTML(
					`<svg class="mdc-checkbox__checkmark"
							viewBox="0 0 24 24">
							<path class="mdc-checkbox__checkmark-path"
								fill="none"
								stroke="white"
								d="M1.73,12.91 8.1,19.28 22.79,4.59"/>
							</svg>`,
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-checkbox__mixedmark"),
				),
			),
		),
	))
}
