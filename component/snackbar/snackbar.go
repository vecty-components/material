// The snackbar package implements a material snackbar component.
//
// See: https://material.io/components/web/catalog/snackbars/
package snackbar // import "agamigo.io/material/component/snackbar"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-snackbar"
     aria-live="assertive"
     aria-atomic="true"
     aria-hidden="true">
  <div class="mdc-snackbar__text"></div>
  <div class="mdc-snackbar__action-wrapper">
    <button type="button" class="mdc-snackbar__action-button"></button>
  </div>
</div>`
)

// S is the interface for a material snackbar component.
type S interface {
	component.C
}

// snackbar is the internal implementation of S made available publicly via
// New().
type snackbar struct {
	component.C
	html string
}

// New creates a material snackbar component that implement the S interface.
// It is a wrapper around component.New.
func New() (c S, err error) {
	newS, err := component.New(component.Snackbar)
	if err != nil {
		return nil, err
	}
	return &snackbar{newS, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (s *snackbar) HTML() string {
	return s.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (s *snackbar) SetHTML(html string) {
	s.html = html
}
