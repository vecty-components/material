package icontoggle

import (
	"agamigo.io/material/icontoggle"
	"agamigo.io/material/ripple"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/icon"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// IT is a vecty-material icontoggle component.
type IT struct {
	*icontoggle.IT
	vecty.Core
	ID            string
	Markup        []vecty.Applyer
	rootElement   *vecty.HTML
	Ripple        bool
	Basic         bool
	ripple        *ripple.R
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
	c.init()
	if c.OffIcon == nil || c.OnIcon == nil {
		panic("OnIcon and/or OffIcon missing in icontoggle.")
	}
	var itMarkup []vecty.Applyer
	var itElement *icon.I
	switch {
	case c.On:
		itElement = c.OnIcon
		itMarkup = append(itMarkup,
			vecty.Attribute("aria-label", c.OnLabel),
			vecty.Attribute("aria-pressed", "true"),
		)
	default:
		itElement = c.OffIcon
		itMarkup = append(itMarkup,
			vecty.Attribute("aria-label", c.OffLabel),
		)
	}
	switch {
	case itElement.ClassOverride != nil:
		itMarkup = append(itMarkup,
			vecty.Data("iconInnerSelector", "."+itElement.ClassOverride[0]),
		)
	default:
		itMarkup = append(itMarkup,
			vecty.Data("iconInnerSelector", ".material-icons"),
		)
	}
	itElement.Markup = append(itElement.Markup,
		vecty.Attribute("aria-hidden", "true"),
	)
	c.rootElement = elem.Span(
		vecty.Markup(
			vecty.Markup(c.Markup...),
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
			vecty.Markup(itMarkup...),
			vecty.MarkupIf(c.On,
				vecty.Class("mdc-icon-toggle--on"),
			),
		),
		vecty.If(!c.On, c.OffIcon),
		vecty.If(c.On, c.OnIcon),
	)
	return c.rootElement
}

func (c *IT) MDCRoot() *base.Base {
	return &base.Base{
		MDC:       c,
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		Basic:     c.Basic,
		RippleC:   c.ripple,
	}
}

func (c *IT) Mount() {
	c.MDCRoot().Mount()
}

func (c *IT) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *IT) init() {
	if c.IT == nil {
		c.IT = icontoggle.New()
	}
}

func (c *IT) wrapChangeHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.ChangeHandler(c, e)
	}
}
