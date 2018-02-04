// The radio package implements a material radio component.
//
// See: https://material.io/components/web/catalog/input-controls/radio-buttons/
package radio // import "agamigo.io/material/component/radio"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-radio">
  <input class="mdc-radio__native-control" type="radio" id="radio-1" name="radios" checked>
  <div class="mdc-radio__background">
    <div class="mdc-radio__outer-circle"></div>
    <div class="mdc-radio__inner-circle"></div>
  </div>
</div>
<label id="radio-1-label" for="radio-1">Radio 1</label>`
)

// R is the interface for a material radio component.
type R interface {
	component.C
}

// radio is the internal implementation of R made available publicly via
// New().
type radio struct {
	component.C
	html string
}

// New creates a material radio component that implement the R interface.
// It is a wrapper around component.New.
func New() (c R, err error) {
	newR, err := component.New(component.Radio)
	if err != nil {
		return nil, err
	}
	return &radio{newR, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (r *radio) HTML() string {
	return r.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (r *radio) SetHTML(html string) {
	r.html = html
}
