// The selection package implements a material selection component.
//
// See: https://material.io/components/web/catalog/input-controls/select-menus/
package selection // import "agamigo.io/material/selection"

import (
	"agamigo.io/material/component"
	"github.com/gopherjs/gopherjs/js"
)

// S is a material selection component. It should only be created using the New
// function.
type S struct {
	*component.C
	SelectedIndex int  `js:"selectedIndex"`
	Disabled      bool `js:"disabled"`
}

// New creates a material selection component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*S, error) {
	newS, err := component.New(component.Select)
	if err != nil {
		return nil, err
	}
	return &S{C: newS}, err
}

// Selected returns the id of the currently selected option. If no id is present
// on the selected option, its textContent is used. Returns an empty string when
// no option is selected.
func (s *S) SelectedString() string {
	v := s.GetObject().Get("value").String()
	if v == "undefined" {
		return ""
	}
	return v
}

// SelectedElem returns a NodeList of either the currently selected option, or
// an empty js.S if nothing is selected.
func (s *S) SelectedElem() *js.Object {
	return s.GetObject().Get("selectedOptions")
}

// Options returns a slice of menu items comprising the select’s options.
func (s *S) Options() *js.Object {
	return s.GetObject().Get("options")
}
