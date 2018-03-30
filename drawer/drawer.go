package drawer

import (
	mbase "agamigo.io/material/base"
	"agamigo.io/material/persistentdrawer"
	"agamigo.io/material/ripple"
	"agamigo.io/material/temporarydrawer"
	"agamigo.io/vecty-material/base"
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
	D mbase.ComponentStartStopper
	vecty.Core
	ID          string
	Markup      []vecty.Applyer
	rootElement *vecty.HTML
	Ripple      bool
	Basic       bool
	ripple      *ripple.R
	Type
	Open          bool
	BelowToolbar  bool
	Toolbar       vecty.ComponentOrHTML
	Header        vecty.ComponentOrHTML
	ToolbarSpacer vecty.ComponentOrHTML
	Content       vecty.ComponentOrHTML
}

// Render implements the vecty.Component interface.
func (c *D) Render() vecty.ComponentOrHTML {
	c.init()
	markup := vecty.Markup(
		vecty.Markup(c.Markup...),
		vecty.Class("mdc-drawer"))
	switch c.Type {
	case Temporary:
		c.rootElement = elem.Aside(
			vecty.Markup(markup,
				vecty.Class("mdc-drawer--temporary"),
				vecty.MarkupIf(c.Open, vecty.Class("mdc-drawer--open"))),
			elem.Navigation(
				vecty.Markup(vecty.Class("mdc-drawer__drawer")),
				c.renderDrawer()))
	case Persistent:
		c.rootElement = elem.Aside(
			vecty.Markup(markup,
				vecty.Class("mdc-drawer--persistent"),
				vecty.MarkupIf(c.Open, vecty.Class("mdc-drawer--open"))),
			elem.Navigation(
				vecty.Markup(vecty.Class("mdc-drawer__drawer")),
				c.renderDrawer()))
	default: // Permanent
		c.rootElement = elem.Navigation(
			vecty.Markup(markup, vecty.Class("mdc-drawer--permanent")),
			c.renderDrawer())
	}
	return c.rootElement
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
			c.Header = base.RenderStoredChild(h)
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

func (c *D) MDCRoot() *base.Base {
	b := &base.Base{
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		Basic:     c.Basic,
		RippleC:   c.ripple,
	}
	if c.Type != Permanent {
		b.MDC = c.D
	}
	return b
}

func (c *D) Mount() {
	c.MDCRoot().Mount()
}

func (c *D) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *D) init() {
	switch {
	case c.D == nil && c.Type == Temporary:
		c.D = temporarydrawer.New()
	case c.D == nil && c.Type == Persistent:
		c.D = persistentdrawer.New()
	}
	switch t := c.D.(type) {
	case *temporarydrawer.TD:
		t.Open = c.Open
	case *persistentdrawer.PD:
		t.Open = c.Open
	}
}
