package menu

import (
	"agamigo.io/material/menu"
	"agamigo.io/material/ripple"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/ul"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// M is a vecty-material menu component.
type M struct {
	*menu.M
	vecty.Core
	ID          string
	Markup      []vecty.Applyer
	rootElement *vecty.HTML
	Ripple      bool
	Basic       bool
	ripple      *ripple.R

	// Open is the visible state of the menu component.
	Open bool `js:"open"`

	// QuickOpen controls whether the menu should open and close without
	// animation. False uses animation, true does not.
	QuickOpen bool `js:"quickOpen"`

	// List is a HTMLUListElement containing the menu's items.
	List vecty.ComponentOrHTML

	// Set AnchorElement to embed the menu component inside an HTMLElement from
	// which the element will be anchored.
	AnchorElement vecty.ComponentOrHTML

	// Define OnSelect to handle "MDCMenu:selected" events. item is the
	// menu item that was selected.
	OnSelect func(index int, item vecty.ComponentOrHTML, e *vecty.Event)

	// Define OnCancel to handle "MDCMenu:selected" events. item is the
	// menu item that was selected.
	OnCancel func(e *vecty.Event)
}

// Render implements the vecty.Component interface.
func (c *M) Render() vecty.ComponentOrHTML {
	c.init()
	switch t := c.List.(type) {
	case *ul.L:
		t.Markup = append(t.Markup,
			vecty.Class("mdc-menu__items"),
			vecty.Attribute("role", "menu"),
		)
		if !c.Open {
			t.Markup = append(t.Markup, vecty.Attribute("aria-hidden", "true"))
		}
		for _, item := range t.Items {
			if i, ok := item.(*ul.Item); ok {
				i.Markup = append(i.Markup,
					vecty.Attribute("role", "menuitem"),
					vecty.Attribute("tabindex", 0),
				)
			}
		}
	case *vecty.HTML:
		vecty.Class("mdc-menu__items").Apply(t)
		vecty.Attribute("role", "menu").Apply(t)
		if c.Open {
			vecty.Attribute("aria-hidden", "false").Apply(t)
		}
	}

	menuMarkup := vecty.Markup(
		vecty.MarkupIf(c.ID != "", prop.ID(c.ID)),
		vecty.Class("mdc-menu"),
		vecty.MarkupIf(c.Open && c.rootElement == nil,
			vecty.Class("mdc-menu--open"),
		),
		vecty.Style("position", "absolute"),
		vecty.Attribute("tabindex", -1),
		&vecty.EventListener{
			Name:     "MDCMenu:selected",
			Listener: c.onSelect,
		},
		&vecty.EventListener{
			Name:     "MDCMenu:cancel",
			Listener: c.onCancel,
		},
	)

	if c.AnchorElement != nil {
		c.rootElement = elem.Div(
			menuMarkup,
			c.List,
		)
		return elem.Div(
			vecty.Markup(
				vecty.Class("mdc-menu-anchor"),
				vecty.Markup(c.Markup...),
			),
			c.AnchorElement,
			c.rootElement,
		)
	}
	c.rootElement = elem.Div(
		menuMarkup,
		vecty.Markup(vecty.Markup(c.Markup...)),
		c.List,
	)
	return c.rootElement
}

func (c *M) MDCRoot() *base.Base {
	return &base.Base{
		MDC:       c,
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.Ripple,
		Basic:     c.Basic,
		RippleC:   c.ripple,
	}
}

func (c *M) Mount() {
	c.MDCRoot().Mount()
}

func (c *M) Unmount() {
	c.MDCRoot().Unmount()
}

func (c *M) init() {
	if c.M == nil {
		// TODO: Make initial values work in material package
		open := js.InternalObject(c).Get("Open").Bool()
		quickOpen := js.InternalObject(c).Get("QuickOpen").Bool()
		c.M = menu.New()
		c.Open = open
		c.QuickOpen = quickOpen
	}
}

func (c *M) onSelect(e *vecty.Event) {
	if c.OnSelect != nil {
		var item vecty.ComponentOrHTML
		i := e.Get("detail").Get("index").Int()
		i = i + c.dividerCountBefore(i)
		switch t := c.List.(type) {
		case *ul.L:
			item = t.Items[i]
		case vecty.List:
			item = t[i]
		}
		c.OnSelect(i, item, e)
	}
}

func (c *M) onCancel(e *vecty.Event) {
	if c.OnCancel != nil {
		c.OnCancel(e)
	}
}

// TODO: Figure out how to use or replicate MDC behavior
// dividerCountBefore returns the number non-valid items (dividers) in the list
// before itemIndex, non-inclusive.
func (c *M) dividerCountBefore(itemIndex int) int {
	count := 0
	var items []vecty.ComponentOrHTML
	switch t := c.List.(type) {
	case *ul.L:
		items = t.Items
	case vecty.List:
		items = t
	}
	for i, item := range items {
		if i > itemIndex {
			break
		}
		if _, ok := item.(*base.StaticComponent); ok {
			count++
		}
	}
	return count
}
