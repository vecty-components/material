// The tabbar package implements a material tabbar component.
//
// See: https://material.io/components/web/catalog/tabs/
package tabbar // import "agamigo.io/material/component/tabbar"

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

// TB is the interface for a material tabbar component.
type TB interface {
	component.C
}

// tabBar is the internal implementation of TB made available publicly via
// New().
type tabBar struct {
	component.C
	html string
}

// New creates a material tabbar component that implement the TB interface.
// It is a wrapper around component.New.
func New() (c TB, err error) {
	newTB, err := component.New(component.TabBar)
	if err != nil {
		return nil, err
	}
	return &tabBar{newTB, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (tb *tabBar) HTML() string {
	return tb.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (tb *tabBar) SetHTML(html string) {
	tb.html = html
}
