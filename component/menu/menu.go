// The menu package implements a material menu component.
//
// See: https://material.io/components/web/catalog/menus/
package menu // import "agamigo.io/material/component/menu"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-simple-menu" tabindex="-1">
  <ul class="mdc-simple-menu__items mdc-list" role="menu" aria-hidden="true">
    <li class="mdc-list-item" role="menuitem" tabindex="0">
      A Menu Item
    </li>
    <li class="mdc-list-item" role="menuitem" tabindex="0">
      Another Menu Item
    </li>
  </ul>
</div>`
)

// M is the interface for a material menu component.
type M interface {
	component.C
}

// menu is the internal implementation of M made available publicly via
// New().
type menu struct {
	component.C
	html string
}

// New creates a material menu component that implement the M interface.
// It is a wrapper around component.New.
func New() (c M, err error) {
	newM, err := component.New(component.Menu)
	if err != nil {
		return nil, err
	}
	return &menu{newM, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (m *menu) HTML() string {
	return m.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (m *menu) SetHTML(html string) {
	m.html = html
}
