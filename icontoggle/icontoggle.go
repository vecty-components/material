package icontoggle

import (
	"agamigo.io/material/icontoggle"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/icon"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// IT is a vecty-material icontoggle component.
type IT struct {
	*base.MDC
	vecty.Core
	Root          vecty.MarkupOrChild
	ChangeHandler func(thisIT *IT, e *vecty.Event)
	On            bool
	Disabled      bool
	OnIcon        *icon.I
	OffIcon       *icon.I
	activeIcon    *icon.I
	OnLabel       string
	OffLabel      string
}

// Render implements the vecty.Component interface.
func (c *IT) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	if c.OffIcon == nil || c.OnIcon == nil {
		panic("OnIcon and/or OffIcon missing in icontoggle.")
	}

	// Built-in root element.
	return elem.Span(
		vecty.Markup(
			c,
			vecty.MarkupIf(rootMarkup != nil, *rootMarkup),
		),
		vecty.If(!c.On, c.OffIcon),
		vecty.If(c.On, c.OnIcon),
	)
}

func (c *IT) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = icontoggle.New()
	}

	var markup []vecty.Applyer
	var itElement *icon.I
	switch {
	case c.On:
		itElement = c.OnIcon
		markup = append(markup,
			vecty.Attribute("aria-label", c.OnLabel),
			vecty.Attribute("aria-pressed", "true"),
		)
	default:
		itElement = c.OffIcon
		markup = append(markup,
			vecty.Attribute("aria-label", c.OffLabel),
		)
	}
	switch {
	case itElement.ClassOverride != nil:
		markup = append(markup,
			vecty.Data("iconInnerSelector", "."+itElement.ClassOverride[0]),
		)
	default:
		markup = append(markup,
			vecty.Data("iconInnerSelector", ".material-icons"),
		)
	}
	// TODO: Add this attribute to icons
	// itElement.Markup = append(itElement.Markup,
	// 	vecty.Attribute("aria-hidden", "true"),
	// )

	vecty.Markup(
		vecty.Class("mdc-icon-toggle"),
		vecty.Attribute("role", "button"),
		vecty.Attribute("aria-pressed", c.On),
		vecty.MarkupIf(!c.Disabled,
			vecty.Attribute("tabindex", "0"),
		),
		vecty.MarkupIf(c.Disabled,
			vecty.Attribute("tabindex", "-1"),
			vecty.Class("mdc-icon-toggle--disabled"),
			vecty.Attribute("aria-hidden", true),
		),
		&vecty.EventListener{
			Name: "MDCIconToggle:change",
			Listener: func(e *vecty.Event) {
				c.On = !c.On
				vecty.Rerender(c)
			},
		},
		vecty.MarkupIf(c.ChangeHandler != nil,
			&vecty.EventListener{
				Name:     "MDCIconToggle:change",
				Listener: c.wrapChangeHandler(),
			},
		),
		vecty.Markup(markup...),
		vecty.MarkupIf(c.On,
			vecty.Class("mdc-icon-toggle--on"),
		),
	).Apply(h)
	c.MDC.RootElement = h
}

func (c *IT) wrapChangeHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.ChangeHandler(c, e)
	}
}
