// The tab package implements a material tab component.
//
// See: https://material.io/components/web/catalog/tabs/
package tab // import "agamigo.io/material/component/tab"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<nav id="basic-tab-bar" class="mdc-tab-bar">
  <a class="mdc-tab mdc-tab--active" href="#one">Home</a>
  <a class="mdc-tab" href="#two">Merchandise</a>
  <a class="mdc-tab" href="#three">About Us</a>
  <span class="mdc-tab-bar__indicator"></span>
</nav>`
)

// T is the interface for a material tab component.
type T interface {
	component.C
}

// tab is the internal implementation of T made available publicly via
// New().
type tab struct {
	component.C
	html string
}

// New creates a material tab component that implement the T interface.
// It is a wrapper around component.New.
func New() (c T, err error) {
	newT, err := component.New(component.Tab)
	if err != nil {
		return nil, err
	}
	return &tab{newT, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (t *tab) HTML() string {
	return t.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (t *tab) SetHTML(html string) {
	t.html = html
}
