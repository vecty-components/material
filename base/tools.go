package base

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type childComponent struct {
	vecty.Core
	child vecty.ComponentOrHTML
}

// RenderStoredChild is a helper which provides a Component which wraps the
// provided ComponentOrHTML. It exists as a workaround to a vecty issue.
//
// See: https://github.com/gopherjs/vecty/issues/191
func RenderStoredChild(child vecty.ComponentOrHTML) *childComponent {
	return &childComponent{child: child}
}

func (s *childComponent) Render() vecty.ComponentOrHTML {
	switch t := s.child.(type) {
	case vecty.List:
		return elem.Div(t)
	}
	return s.child
}

func (s *childComponent) SkipRender(prev vecty.Component) bool {
	switch prev.(type) {
	case *childComponent:
		return true
	}
	return false
}
