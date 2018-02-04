// The ripple package implements a material ripple component.
//
// See: https://material.io/components/web/catalog/ripples/
package ripple // import "agamigo.io/material/component/ripple"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-ripple-surface mdc-ripple-surface--primary my-surface" tabindex="0">
  Surface with a primary-colored ripple.
</div>`
)

// R is the interface for a material ripple component.
type R interface {
	component.C
}

// ripple is the internal implementation of R made available publicly via
// New().
type ripple struct {
	component.C
	html string
}

// New creates a material ripple component that implement the R interface.
// It is a wrapper around component.New.
func New() (c R, err error) {
	newR, err := component.New(component.Ripple)
	if err != nil {
		return nil, err
	}
	return &ripple{newR, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (r *ripple) HTML() string {
	return r.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (r *ripple) SetHTML(html string) {
	r.html = html
}
