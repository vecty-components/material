// The selection package implements a material selection component.
//
// See: https://material.io/components/web/catalog/input-controls/select-menus/
package selection // import "agamigo.io/material/component/selection"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-select" role="listbox">
  <div class="mdc-select__surface" tabindex="0">
    <div class="mdc-select__label">Pick a Food Group</div>
    <div class="mdc-select__selected-text"></div>
    <div class="mdc-select__bottom-line"></div>
  </div>
  <div class="mdc-simple-menu mdc-select__menu">
    <ul class="mdc-list mdc-simple-menu__items">
      <li class="mdc-list-item" role="option" tabindex="0">
        Bread, Cereal, Rice, and Pasta
      </li>
      <li class="mdc-list-item" role="option" tabindex="0">
        Vegetables
      </li>
    </ul>
  </div>
</div>`
)

// S is the interface for a material selection component.
type S interface {
	component.C
}

// selection is the internal implementation of S made available publicly via
// New().
type selection struct {
	component.C
	html string
}

// New creates a material selection component that implement the S interface.
// It is a wrapper around component.New.
func New() (c S, err error) {
	newS, err := component.New(component.Select)
	if err != nil {
		return nil, err
	}
	return &selection{newS, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (s *selection) HTML() string {
	return s.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (s *selection) SetHTML(html string) {
	s.html = html
}
