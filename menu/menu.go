// menu implements a material menu component.
//
// See: https://material.io/components/web/catalog/menus/
package menu // import "agamigo.io/material/menu"

import (
	"agamigo.io/material/internal/base"
	"github.com/gopherjs/gopherjs/js"
)

type Corner int

const (
	TOP_LEFT     Corner = 0
	TOP_RIGHT           = 4
	BOTTOM_LEFT         = 1
	BOTTOM_RIGHT        = 5
	TOP_START           = 8
	TOP_END             = 12
	BOTTOM_START        = 9
	BOTTOM_END          = 13
)

// M is a material menu component.
type M struct {
	mdc *base.Component

	// Open is the visible state of the menu component.
	Open bool `js:"open"`

	// QuickOpen controls whether the menu should open and close without
	// animation. False uses animation, true does not.
	QuickOpen bool `js:"quickOpen"`

	// For now we give read-only access with the Items() method.
	items []*js.Object `js:"items"`

	// For now we give read-only access with the ItemsContainer() method.
	itemsContainer *js.Object `js:"itemsContainer_"`
}

// Margins holds margin values used to configure menu anchor margins via
// {Set}AnchorMargins() methods.
type Margins struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *M) Start(rootElem *js.Object) error {
	err := base.Start(c.Component(), rootElem)
	if err != nil {
		return err
	}
	err = c.afterStart()
	if err != nil {
		// TODO: handle afterStart + stop error
		_ = c.Stop()
		return err
	}
	return nil
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *M) Stop() error {
	return base.Stop(c.mdc)
}

// Component returns the component's underlying base.Component.
func (c *M) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{}
		c.mdc.Type = base.ComponentType{
			MDCClassName:     "MDCMenu",
			MDCCamelCaseName: "menu",
		}
	}
	return c.mdc
}

// OpenFocus opens the menu with an item at index given initial focus.
func (m *M) OpenFocus(index int) {
	m.Component().Call("show", index)
}

// Items returns the HTMLLIElements that represent the menu's items.
func (m *M) Items() []*js.Object {
	return m.items
}

// ItemsContainer is the HTMLUListElement that contains the menu's items
func (m *M) ItemsContainer() *js.Object {
	return m.itemsContainer
}

// AnchorCorner returns the Corner the menu is/will be attached to.
func (m *M) AnchorCorner() Corner {
	return Corner(m.Component().Get("foundation_").Get("anchorCorner_").Int())
}

// AnchorCorner sets the Corner the menu is/will be attached to.
func (m *M) SetAnchorCorner(c Corner) {
	m.Component().Call("setAnchorCorner", c)
}

// AnchorMargins returns the distance from the anchor point that the menu
// is/will be.
func (m *M) AnchorMargins() *Margins {
	o := m.Component().Get("foundation_").Get("anchorMargin_")
	return &Margins{
		Left:   o.Get("left").Int(),
		Right:  o.Get("right").Int(),
		Top:    o.Get("top").Int(),
		Bottom: o.Get("bottom").Int(),
	}
}

// AnchorMargins sets the distance from the anchor point that the menu is/will
// be.
func (m *M) SetAnchorMargins(ms *Margins) {
	o := &js.M{
		"left":   ms.Left,
		"right":  ms.Right,
		"top":    ms.Top,
		"bottom": ms.Bottom,
	}
	m.Component().Call("setAnchorMargin", o)
}
