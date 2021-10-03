package iconbutton

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/components/iconbutton"
	"github.com/vecty-material/material/icon"
)

// IB is a vecty-material iconbutton component.
type IB struct {
	*base.MDC
	vecty.Core
	Root          vecty.MarkupOrChild
	ChangeHandler func(thisIB *IB, e *vecty.Event)
	On            bool
	OnIcon        *icon.I
	OffIcon       *icon.I
	OnLabel       string
	OffLabel      string
}

// Render implements the vecty.Component interface.
func (c *IB) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	if c.OffIcon == nil || c.OnIcon == nil {
		panic("OnIcon and/or OffIcon missing in iconbutton.")
	}

	// Built-in root element.
	return elem.Span(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		vecty.If(!c.On, c.OffIcon),
		vecty.If(c.On, c.OnIcon),
	)
}

func (c *IB) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = iconbutton.New()
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
		vecty.Class("mdc-icon-button"),
		vecty.Attribute("role", "button"),
		vecty.Attribute("aria-pressed", c.On),
		&vecty.EventListener{
			Name: "MDCIconButtonToggle:change",
			Listener: func(e *vecty.Event) {
				c.On = !c.On
				vecty.Rerender(c)
			},
		},
		vecty.MarkupIf(c.ChangeHandler != nil,
			&vecty.EventListener{
				Name:     "MDCIconButtonToggle:change",
				Listener: c.wrapChangeHandler(),
			},
		),
		vecty.Markup(markup...),
		vecty.MarkupIf(c.On,
			vecty.Class("mdc-icon-button--on"),
		),
	).Apply(h)
	c.MDC.RootElement = h
}

func (c *IB) wrapChangeHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.ChangeHandler(c, e)
	}
}
