package drawer

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/components/persistentdrawer"
	"github.com/vecty-material/material/components/temporarydrawer"
)

type Type int

const (
	Temporary Type = iota
	Persistent
	Permanent
)

// D is a vecty-material drawer component.
type D struct {
	*base.MDC
	vecty.Core
	Root vecty.MarkupOrChild
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
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	markup := vecty.Markup(
		c,
		base.MarkupIfNotNil(rootMarkup),
	)

	// Built-in root element.
	if c.Type == Permanent {
		return elem.Navigation(
			markup,
			c.renderDrawer(),
		)
	}
	// Persistent or Temporary drawer.
	return elem.Aside(
		markup,
		elem.Navigation(
			vecty.Markup(vecty.Class("mdc-drawer__drawer")),
			c.renderDrawer(),
		),
	)
}

func (c *D) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		switch c.Type {
		case Permanent:
		case Temporary:
			c.MDC.Component = temporarydrawer.New()
			c.MDC.Component.(*temporarydrawer.TD).Open = c.Open
		case Persistent:
			c.MDC.Component = persistentdrawer.New()
			c.MDC.Component.(*persistentdrawer.PD).Open = c.Open
		}
	}

	markup := []vecty.Applyer{
		vecty.Class("mdc-drawer"),
		vecty.MarkupIf(c.Open, vecty.Class("mdc-drawer--open")),
	}
	switch c.Type {
	case Permanent:
		markup = append(markup, vecty.Class("mdc-drawer--permanent"))
	case Temporary:
		markup = append(markup, vecty.Class("mdc-drawer--temporary"))
	case Persistent:
		markup = append(markup, vecty.Class("mdc-drawer--persistent"))
	}

	vecty.Markup(markup...).Apply(h)
	c.MDC.RootElement = h
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
