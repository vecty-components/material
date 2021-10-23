package base

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type StaticComponentKeyer struct {
	vecty.Core
	Child vecty.ComponentOrHTML
	key   string
}

// RenderStoredChild is a helper which provides a Component which wraps the
// provided ComponentOrHTML. It exists as a workaround to a vecty issue.
//
// See: https://github.com/hexops/vecty/issues/191
func RenderStoredChildKeyer(child vecty.ComponentOrHTML) vecty.Component {
	switch c := child.(type) {
	case vecty.Component:
		return c
	}

	return &StaticComponentKeyer{Child: child}
}

func (c *StaticComponentKeyer) Render() vecty.ComponentOrHTML {
	switch t := c.Child.(type) {
	case vecty.List:
		return elem.Div(t)
	}
	return c.Child
}

func (c *StaticComponentKeyer) SkipRender(prev vecty.Component) bool {
	switch prev.(type) {
	case *StaticComponent:
		return true
	}
	return false
}

func (c *StaticComponentKeyer) Key() interface{} {
	return c.key
}
