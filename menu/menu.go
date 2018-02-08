// The menu package implements a material menu component.
//
// See: https://material.io/components/web/catalog/menus/
package menu // import "agamigo.io/material/menu"

import (
	"agamigo.io/material/component"
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

const ITEMS_SELECTOR = ".mdc-menu__items"

// M is a material menu component. It should only be created using the New
// function.
type M struct {
	*component.C
	*State
	Open           bool       `js:"open"`
	Items          *js.Object `js:"items"`
	ItemsContainer *js.Object `js:"itemsContainer_"`
}

// State holds some of the menu's r/w settings. It is in its own struct due to
// the structure of the MDC component storing properties in the "foundation_".
// The properties here can still be accessed directly from an instance of M.
//
// TODO: Find a way to put these properties into the M struct directly.
type State struct {
	*state
	*AnchorMargins
	QuickOpen bool `js:"quickOpen"`
}

// state exists as the unexposed portion of State.
type state struct {
	object *js.Object
}

type AnchorMargins struct {
	*anchorMargins
	LeftMargin   int `js:"left"`
	RightMargin  int `js:"right"`
	TopMargin    int `js:"top"`
	BottomMargin int `js:"bottom"`
}

type anchorMargins struct {
	object *js.Object
}

// New creates a material menu component. It is a wrapper around component.New
// which instantiates the component from the MDC library.
func New() (*M, error) {
	newM, err := component.New(component.Menu)
	if err != nil {
		return nil, err
	}
	m := &M{
		C: newM,
		State: &State{
			state:         &state{},
			AnchorMargins: &AnchorMargins{anchorMargins: &anchorMargins{}}},
	}
	return m, err
}

// Start wraps component.Start.
func (m *M) Start() (err error) {
	err = m.C.Start()
	if err != nil {
		return err
	}

	m.State.object = m.GetObject().Get("foundation_")
	m.State.anchorMargins.object = m.State.object.Get("anchorMargin_")
	return err
}

// StartWith wraps component.StartWith.
func (m *M) StartWith(querySelector string) (err error) {
	err = m.C.StartWith(querySelector)
	if err != nil {
		return err
	}

	m.State = &State{state: &state{object: m.GetObject().Get("foundation_")}}
	m.State.anchorMargins.object = m.State.object.Get("anchorMargin_")
	return err
}

// StartWithElement wraps component.StartWithElement.
func (m *M) StartWithElement(e *js.Object) (err error) {
	err = m.C.StartWithElement(e)
	if err != nil {
		return err
	}

	m.State = &State{state: &state{object: m.GetObject().Get("foundation_")}}
	m.State.anchorMargins.object = m.State.object.Get("anchorMargin_")
	return err
}

// func (m *M) IsOpen() bool {
// 	return m.GetObject().Get("open").Bool()
// }

// TODO: Should we manipulate menu items via JS?
// func (m *M) SetItems(elems ...*js.Object) {
// 	m.GetObject().Set("items", js.S(elems))
// }

// func (m *M) AddItems(elems ...*js.Object) {
// 	return
// }

// func (m *M) DelItems(elems ...*js.Object) {
// 	return
// }

// func (m *M) Items() js.S {
// 	items := m.GetObject().Get("items").Interface().([]interface{})
// 	// if !ok {
// 	// 	panic("Unable to convert element list in menu.Items.")
// 	// }
// 	return items
// }

// func (m *M) ItemsContainer() *js.Object {
// 	return m.GetObject().Get("root_").Call("querySelector", ITEMS_SELECTOR)
// }

// func (m *M) IsQuickOpen() bool {
// 	return m.GetObject().Get("quickOpen").Bool()
// }

// func (m *M) SetQuickOpen(q bool) {
// 	m.GetObject().Set("quickOpen", q)
// }

// func (m *M) Open() {
// 	m.GetObject().Call("show")
// }

func (m *M) OpenFocus(index int) {
	m.GetObject().Call("show", index)
}

// func (m *M) Close() {
// 	m.GetObject().Call("hide")
// }

func (m *M) AnchorCorner() Corner {
	return Corner(m.GetObject().Get("foundation_").Get("anchorCorner_").Int())
}

func (m *M) SetAnchorCorner(c Corner) {
	m.GetObject().Call("setAnchorCorner", c)
}

// func (m *M) AnchorMargin() (pixels int) {
// 	return m.GetObject().Get("foundation_").Get("anchorMargin_").Int()
// }

// func (m *M) SetAnchorMargin(margins *js.M) {
// 	m.GetObject().Call("setAnchorMargin", pixels)
// }
