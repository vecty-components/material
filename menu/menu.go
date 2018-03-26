package menu

import (
	"agamigo.io/material/menu"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/ul"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// M is a vecty-material menu component.
type M struct {
	*base.Base
	*State
}

type State struct {
	*menu.M

	// Open is the visible state of the menu component.
	Open bool `js:"open"`

	// QuickOpen controls whether the menu should open and close without
	// animation. False uses animation, true does not.
	QuickOpen bool `js:"quickOpen"`

	// List is a HTMLUListElement containing the menu's items.
	List vecty.ComponentOrHTML `vecty:"prop"`

	// Set AnchorElement to embed the menu component inside an HTMLElement from
	// which the element will be anchored.
	AnchorElement vecty.ComponentOrHTML

	// Define SelectHandler to handle "MDCMenu:selected" events. item is the
	// menu item that was selected.
	SelectHandler func(index int, item vecty.ComponentOrHTML, e *vecty.Event)
}

func New(p *base.Props, s *State) *M {
	open := js.InternalObject(s).Get("Open").Bool()
	c := &M{}
	if s == nil {
		s = &State{}
	}
	if s.M == nil {
		s.M = menu.New()
	}
	c.State = s
	c.Base = base.New(p, c)
	c.Open = open
	return c
}

// Render implements the vecty.Component interface.
func (c *M) Render() vecty.ComponentOrHTML {
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
		if !c.Open {
			vecty.Attribute("aria-hidden", "true").Apply(t)
		}
	}

	menuMarkup := vecty.Markup(
		vecty.MarkupIf(c.Props.ID != "", prop.ID(c.Props.ID)),
		vecty.Class("mdc-menu"),
		vecty.MarkupIf(c.Open, vecty.Class("mdc-menu--open")),
		vecty.Style("position", "absolute"),
		vecty.Attribute("tabindex", -1),
		vecty.MarkupIf(c.SelectHandler != nil,
			&vecty.EventListener{
				Name:     "MDCMenu:selected",
				Listener: c.wrapSelectHandler(),
			},
		),
	)

	if c.AnchorElement != nil {
		return elem.Div(
			vecty.Markup(
				vecty.Class("mdc-menu-anchor"),
				vecty.Markup(c.Props.Markup...),
			),
			c.AnchorElement,
			c.Base.Render(elem.Div(
				menuMarkup,
				c.List,
			)),
		)
	}
	return c.Base.Render(elem.Div(
		menuMarkup,
		vecty.Markup(vecty.Markup(c.Props.Markup...)),
		c.List,
	))
}

func (c *M) wrapSelectHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		var item vecty.ComponentOrHTML
		i := e.Get("detail").Get("index").Int()
		i = i + c.dividerCountBefore(i)
		switch t := c.List.(type) {
		case *ul.L:
			item = t.Items[i]
		case vecty.List:
			item = t[i]
		}
		c.SelectHandler(i, item, e)
	}
}

// dividerCountBefore returns the number valid items in the list before itemIndex,
// non-inclusive. Dividers in the List slice do not count as valid items.
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
