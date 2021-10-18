package iconbutton

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/icon"
)

// IB is a vecty-material iconbutton component.
type IB struct {
	*base.MDC
	vecty.Core
	Root     vecty.MarkupOrChild   `vecty:"prop"`
	OnChange func(e *vecty.Event)  `vecty:"prop"`
	On       bool                  `vecty:"prop"`
	OnIcon   vecty.ComponentOrHTML `vecty:"prop"`
	OffIcon  vecty.ComponentOrHTML `vecty:"prop"`
	OnLabel  string                `vecty:"prop"`
	OffLabel string                `vecty:"prop"`
	OnClick  func(*vecty.Event)    `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (c *IB) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	if c.OnIcon == nil {
		panic("OnIcon missing in iconbutton (and is required).")
	}

	// Built-in root element.
	return elem.Button(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		vecty.If(!c.On, c.OffIcon),
		vecty.If(
			c.On || c.OffIcon == nil, c.OnIcon,
		),
	)
}

func (c *IB) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCIconButtonToggle",
				MDCCamelCaseName: "iconButton",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{
			"on": &c.On,
		})
	}

	var markup []vecty.Applyer
	var itElement vecty.ComponentOrHTML
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
	itIcon, _ := itElement.(*icon.I)
	switch {
	case itIcon != nil && itIcon.ClassOverride != nil:
		markup = append(markup,
			vecty.Data("iconInnerSelector", "."+itIcon.ClassOverride[0]),
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
		vecty.Class("mdc-icon-button", "material-icons"),
		vecty.Attribute("role", "button"),
		vecty.Attribute("aria-pressed", c.On),
		vecty.MarkupIf(c.OnChange != nil,
			&vecty.EventListener{
				Name:     "MDCIconButtonToggle:change",
				Listener: c.OnChange,
			},
		),
		vecty.MarkupIf(c.OnClick != nil,
			event.Click(c.OnClick),
		),
		vecty.Markup(markup...),
		vecty.MarkupIf(c.On,
			vecty.Class("mdc-icon-button--on"),
		),
	).Apply(h)
	c.MDC.RootElement = h
}
