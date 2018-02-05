// The slider package implements a material slider component.
//
// See: https://material.io/components/web/catalog/input-controls/sliders/
package slider // import "agamigo.io/material/component/slider"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-slider" tabindex="0" role="slider"
     aria-valuemin="0" aria-valuemax="100" aria-valuenow="0"
     aria-label="Select Value">
  <div class="mdc-slider__track-container">
    <div class="mdc-slider__track"></div>
  </div>
  <div class="mdc-slider__thumb-container">
    <svg class="mdc-slider__thumb" width="21" height="21">
      <circle cx="10.5" cy="10.5" r="7.875"></circle>
    </svg>
    <div class="mdc-slider__focus-ring"></div>
  </div>
</div>`
)

// S is the interface for a material slider component.
type S interface {
	component.C
}

// slider is the internal implementation of S made available publicly via
// New().
type slider struct {
	component.C
	html string
}

// New creates a material slider component that implement the S interface.
// It is a wrapper around component.New.
func New() (c S, err error) {
	newS, err := component.New(component.Slider)
	if err != nil {
		return nil, err
	}
	return &slider{newS, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (s *slider) HTML() string {
	return s.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (s *slider) SetHTML(html string) {
	s.html = html
}
