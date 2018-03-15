package drawer

import (
	mbase "agamigo.io/material/base"
	"agamigo.io/material/persistentdrawer"
	"agamigo.io/material/temporarydrawer"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Type int

const (
	Temporary Type = iota
	Persistent
	Permanent
)

// D is a vecty-material drawer component.
type D struct {
	*base.Base
	*State
}

type State struct {
	object *js.Object
	mbase.ComponentStartStopper
	Type
	Open          bool `js:"open"`
	BelowToolbar  bool
	Toolbar       vecty.ComponentOrHTML
	Header        vecty.ComponentOrHTML
	ToolbarSpacer vecty.ComponentOrHTML
	Content       vecty.ComponentOrHTML
}

func New(p *base.Props, s *State) *D {
	c := &D{}
	if s == nil {
		s = &State{}
	}
	switch {
	case s.ComponentStartStopper != nil || s.Type == Permanent:
		break
	case s.Type == Temporary:
		s.ComponentStartStopper = temporarydrawer.New()
		s.object = s.ComponentStartStopper.Component().Object
	case s.Type == Persistent:
		s.ComponentStartStopper = persistentdrawer.New()
		s.object = s.ComponentStartStopper.Component().Object
	}
	c.State = s
	if c.Type == Permanent {
		c.Base = base.New(p, nil)
	} else {
		c.Base = base.New(p, c)
	}
	return c
}

func (c *D) Mount() {
	c.Base.Mount()
	if c.Type == Permanent {
		return
	}
	c.object = c.Component().Object
}

// Render implements the vecty.Component interface.
func (c *D) Render() vecty.ComponentOrHTML {
	markup := vecty.Markup(
		vecty.Markup(c.Props.Markup...),
		vecty.Class("mdc-drawer"))
	var element *vecty.HTML
	switch c.Type {
	case Temporary:
		element = elem.Aside(
			vecty.Markup(markup,
				vecty.Class("mdc-drawer--temporary"),
				vecty.MarkupIf(c.Open, vecty.Class("mdc-drawer--open"))),
			elem.Navigation(
				vecty.Markup(vecty.Class("mdc-drawer__drawer")),
				c.renderDrawer()))
	case Persistent:
		element = elem.Aside(
			vecty.Markup(markup,
				vecty.Class("mdc-drawer--persistent"),
				vecty.MarkupIf(c.Open, vecty.Class("mdc-drawer--open"))),
			elem.Navigation(
				vecty.Markup(vecty.Class("mdc-drawer__drawer")),
				c.renderDrawer()))
	default: // Permanent
		element = elem.Navigation(
			vecty.Markup(markup, vecty.Class("mdc-drawer--permanent")),
			c.renderDrawer())
	}
	return c.Base.Render(element)
}

func (c *D) renderDrawer() vecty.List {
	var elements []vecty.ComponentOrHTML
	if c.ToolbarSpacer != nil {
		var h *vecty.HTML
		var ok bool
		h, ok = c.ToolbarSpacer.(*vecty.HTML)
		if h != nil && ok {
			vecty.Class("mdc-drawer__toolbar-spacer").Apply(h)
			elements = append(elements, elem.Div(
				vecty.Markup(vecty.Class("mdc-drawer__toolbar-spacer")),
				vecty.If(ok, h),
				vecty.If(!ok, c.ToolbarSpacer),
			))
		}
	}
	if c.Header != nil {
		var h *vecty.HTML
		var ok bool
		h, ok = c.Header.(*vecty.HTML)
		if h != nil && ok {
			vecty.Class("mdc-drawer__header-content").Apply(h)
			elements = append(elements, elem.Header(
				vecty.Markup(vecty.Class("mdc-drawer__header")),
				vecty.If(ok, h),
				vecty.If(!ok, c.Header),
			))
		}
	}
	if c.Content != nil {
		elements = append(elements, elem.Navigation(
			vecty.Markup(vecty.Class("mdc-drawer__content")),
			c.Content,
		))
	}
	return elements
}
