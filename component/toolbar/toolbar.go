// The toolbar package implements a material toolbar component.
//
// See: https://material.io/components/web/catalog/toolbar/
package toolbar

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<header class="mdc-toolbar">
  <div class="mdc-toolbar__row">
    <section class="mdc-toolbar__section mdc-toolbar__section--align-start">
      <a href="#" class="material-icons mdc-toolbar__menu-icon">menu</a>
      <span class="mdc-toolbar__title">Title</span>
    </section>
  </div>
</header>`
)

// T is the interface for a material toolbar component.
type T interface {
	component.C
}

// toolbar is the internal implementation of T made available publicly via
// New().
type toolbar struct {
	component.C
	html string
}

// New creates a material toolbar component that implement the T interface.
// It is a wrapper around component.New.
func New() (c T, err error) {
	newT, err := component.New(component.Toolbar)
	if err != nil {
		return nil, err
	}
	return &toolbar{newT, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (t *toolbar) HTML() string {
	return t.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (t *toolbar) SetHTML(html string) {
	t.html = html
}
