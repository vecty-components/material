package base

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type StaticComponent struct {
	vecty.Core
	Child vecty.ComponentOrHTML
}

// RenderStoredChild is a helper which provides a Component which wraps the
// provided ComponentOrHTML. It exists as a workaround to a vecty issue.
//
// See: https://github.com/gopherjs/vecty/issues/191
func RenderStoredChild(child vecty.ComponentOrHTML) *StaticComponent {
	return &StaticComponent{Child: child}
}

func (c *StaticComponent) Render() vecty.ComponentOrHTML {
	switch t := c.Child.(type) {
	case vecty.List:
		return elem.Div(t)
	}
	return c.Child
}

func (c *StaticComponent) SkipRender(prev vecty.Component) bool {
	switch prev.(type) {
	case *StaticComponent:
		return true
	}
	return false
}
