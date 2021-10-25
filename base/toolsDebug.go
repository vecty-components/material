package base

import (
	"log"

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
//func RenderStoredChildKeyer(child vecty.ComponentOrHTML) vecty.Component {
func RenderStoredChildKeyer(child vecty.ComponentOrHTML) vecty.Component {
	switch c := child.(type) {
	case vecty.Component:
		log.Printf("%v %v %p %T %+v", "material/base/toolsDebug/base.RenderStoredChildKeyer()1", "child =", c, c, c)
		return c
	}
	sc := &StaticComponentKeyer{Child: child}
	log.Printf("%v %v %p %T %+v", "material/base/toolsDebug/base.RenderStoredChildKeyer()2", "StaticComponentKeyer =", sc, sc, sc)
	return sc
}

func (c *StaticComponentKeyer) Render() vecty.ComponentOrHTML {
	log.Printf("%v %v %p %T", "material/base/toolsDebug/base.StaticComponentKeyer.Render()1", "c.Child =", c.Child, c.Child)
	switch t := c.Child.(type) {
	case vecty.List:
		return elem.Div(t)
	}
	return c.Child
}

func (c *StaticComponentKeyer) SkipRender(prev vecty.Component) bool {
	switch prev.(type) {
	case *StaticComponentKeyer:
		log.Printf("%v %v %v %T", "material/base/toolsDebug/base.StaticComponentKeyer.SkipRender()1", "SkipRender = true", "prev =", prev)
		//return true
		return false
	}
	log.Printf("%v %v %v %T", "material/base/toolsDebug/base.StaticComponentKeyer.SkipRender()2", "SkipRender = false", "prev =", prev)
	return false
}

func (c *StaticComponentKeyer) Key() interface{} {
	if c.key == "" {
		c.key = Key()
	}
	log.Printf("%v %v %p %+v", "material/base/toolsDebug/base.StaticComponentKeyer.Key()1", "c =", c, c)
	return c.key
}
