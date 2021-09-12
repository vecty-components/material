package menu

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/components/menu"
	"github.com/vecty-material/material/ul"
)

// M is a vecty-material menu component.
type M struct {
	*menu.M
	*base.MDC
	vecty.Core
	Root       vecty.MarkupOrChild
	menuAnchor *vecty.HTML

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
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	// TODO: Make initial values work in material package
	open := c.RootElement.Node().Get("Open").Bool()

	listMarkup := []vecty.Applyer{
		vecty.Class("mdc-menu__items"),
		vecty.Attribute("role", "menu"),
	}
	if !open {
		listMarkup = append(listMarkup, vecty.Attribute("aria-hidden", "true"))
	}
	switch t := c.List.(type) {
	case *ul.L:
		if mu := base.MarkupOnly(t.Root); mu != nil {
			listMarkup = append(listMarkup, mu)
		}
		t.Root = vecty.Markup(listMarkup...)
		for _, item := range t.Items {
			if i, ok := item.(*ul.Item); ok {
				itemMarkup := []vecty.Applyer{
					vecty.Attribute("role", "menuitem"),
					vecty.Attribute("tabindex", 0),
				}
				if mu := base.MarkupOnly(i.Root); mu != nil {
					itemMarkup = append(itemMarkup, mu)
				}
				i.Root = vecty.Markup(itemMarkup...)
			}
		}
	case *vecty.HTML:
		vecty.Class("mdc-menu__items").Apply(t)
		vecty.Attribute("role", "menu").Apply(t)
		if open {
			vecty.Attribute("aria-hidden", "false").Apply(t)
		}
	}

	menuElement := elem.Div(
		vecty.Markup(
			c,
			vecty.MarkupIf(rootMarkup != nil, *rootMarkup),
		),
		c.List,
	)

	if c.AnchorElement != nil {
		c.menuAnchor = elem.Div(
			vecty.Markup(
				vecty.Class("mdc-menu-anchor"),
			),
			c.AnchorElement,
			menuElement,
		)
		return c.menuAnchor
	}
	return menuElement
}

func (c *M) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.M == nil, c.MDC.Component == nil:
		// TODO: Make initial values work in material package
		open := c.RootElement.Node().Get("Open").Bool()
		quickOpen := c.RootElement.Node().Get("QuickOpen").Bool()
		c.M = menu.New()
		c.MDC.Component = c.M
		c.Open = open
		c.QuickOpen = quickOpen
	}

	vecty.Markup(
		vecty.Class("mdc-menu"),
		vecty.MarkupIf(c.Open,
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
	).Apply(h)

	if c.menuAnchor != nil {
		c.MDC.RootElement = c.menuAnchor
		return
	}
	c.MDC.RootElement = h
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
