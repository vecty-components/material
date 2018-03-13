// https://material.io/components/web/catalog/buttons/
package button // import "agamigo.io/vecty-material/button"

import (
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// B is a vecty-material button component.
type B struct {
	*base.Base
	*State
}

type State struct {
	Label        vecty.ComponentOrHTML
	Icon         vecty.ComponentOrHTML
	Disabled     bool
	Raised       bool
	Unelevated   bool
	Stroked      bool
	Dense        bool
	Href         string
	ClickHandler func(*vecty.Event)
}

func New(p *base.Props, s *State) *B {
	c := &B{}
	if s == nil {
		s = &State{}
	}
	c.State = s
	c.Base = base.New(p, nil)
	return c
}

// Render implements the vecty.Component interface.
func (c *B) Render() vecty.ComponentOrHTML {
	var ico *vecty.HTML
	switch t := c.Icon.(type) {
	case nil:
		ico = nil
	case vecty.Component:
		ico = t.Render().(*vecty.HTML)
	case *vecty.HTML:
		ico = t
	}
	if ico != nil {
		vecty.Class("mdc-button__icon").Apply(ico)
	}

	return c.Base.Render(elem.Button(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.Class("mdc-button"),
			prop.Type(prop.TypeButton),
			vecty.MarkupIf(c.ClickHandler != nil,
				event.Click(c.ClickHandler),
			),
			vecty.Property("disabled", c.Disabled),
			vecty.MarkupIf(c.Raised,
				vecty.Class("mdc-button--raised"),
			),
			vecty.MarkupIf(c.Unelevated,
				vecty.Class("mdc-button--unelevated"),
			),
			vecty.MarkupIf(c.Stroked,
				vecty.Class("mdc-button--stroked"),
			),
			vecty.MarkupIf(c.Dense,
				vecty.Class("mdc-button--dense"),
			),
		),
		vecty.If(ico != nil, ico),
		vecty.If(c.Label != nil, base.RenderStoredChild(c.Label)),
	))
}
