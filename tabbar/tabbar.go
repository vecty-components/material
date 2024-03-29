package tabbar

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-components/material/base"
)

type T struct {
	Label  vecty.ComponentOrHTML
	Markup vecty.MarkupList
}

func (c *T) renderTab(active bool) *vecty.HTML {
	return elem.Button(
		c.Markup,
		vecty.Markup(
			vecty.Class("mdc-tab"),
			vecty.Attribute("role", "tab"),
			vecty.MarkupIf(
				active,
				vecty.Attribute("aria-selected", "true"),
				vecty.Attribute("tabIndex", "0"),
				vecty.Class("mdc-tab--active"),
			),
			vecty.MarkupIf(
				active,
				vecty.Attribute("aria-selected", "false"),
				vecty.Attribute("tabIndex", "-1"),
			),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class("mdc-tab__content"),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class("mdc-tab__text-label"),
				),
				c.Label,
			),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class("mdc-tab-indicator"),
				vecty.MarkupIf(
					active,
					vecty.Class("mdc-tab-indicator--active"),
				),
			),
			elem.Span(
				vecty.Markup(
					vecty.Class(
						"mdc-tab-indicator__content",
						"mdc-tab-indicator__content--underline",
					),
				),
			),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class("mdc-tab__ripple"),
			),
		),
	)
}

// TB is a vecty-material tabbar component.
type TB struct {
	*base.MDC
	vecty.Core
	Root     vecty.MarkupOrChild `vecty:"prop"`
	Tabs     []*T                `vecty:"prop"`
	Active   uint                `vecty:"prop"`
	OnChange func(index int)     `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (c *TB) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	if int(c.Active)-1 > len(c.Tabs) {
		panic("active tab index must not exceed len(c.Tabs)-1")
	}

	tabs := make([]vecty.MarkupOrChild, len(c.Tabs))
	for i, t := range c.Tabs {
		tabs[i] = t.renderTab(int(c.Active) == i)
	}

	// Built in root element.
	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-tab-scroller"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-tab-scroller__scroll-area"),
				),
				elem.Div(
					append(
						[]vecty.MarkupOrChild{
							vecty.Markup(
								vecty.Class("mdc-tab-scroller__scroll-content"),
							),
						},
						tabs...,
					)...,
				),
			),
		),
	)
}

func (c *TB) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCTabBar",
				MDCCamelCaseName: "tabBar",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{})
	}
	c.MDC.RootElement = h
	vecty.Markup(
		vecty.Class("mdc-tab-bar"),
		vecty.Attribute("role", "tablist"),
		&vecty.EventListener{
			Name:     "MDCTabBar:activated",
			Listener: c.onChange,
		},
	).Apply(h)
}

func (c *TB) onChange(e *vecty.Event) {
	if c.OnChange != nil {
		c.OnChange(
			e.Get("detail").Get("index").Int(),
		)
	}
}
