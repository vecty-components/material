// The linearprogress package implements a material linearprogress component.
//
// See: https://material.io/components/web/catalog/linear-progress/
package linearprogress // import "agamigo.io/material/component/linearprogress"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div role="progressbar" class="mdc-linear-progress">
  <div class="mdc-linear-progress__buffering-dots"></div>
  <div class="mdc-linear-progress__buffer"></div>
  <div class="mdc-linear-progress__bar mdc-linear-progress__primary-bar">
    <span class="mdc-linear-progress__bar-inner"></span>
  </div>
  <div class="mdc-linear-progress__bar mdc-linear-progress__secondary-bar">
    <span class="mdc-linear-progress__bar-inner"></span>
  </div>
</div>`
)

// LP is the interface for a material linearprogress component.
type LP interface {
	component.C
}

// linearProgress is the internal implementation of LP made available publicly via
// New().
type linearProgress struct {
	component.C
	html string
}

// New creates a material linearprogress component that implement the LP interface.
// It is a wrapper around component.New.
func New() (c LP, err error) {
	newLP, err := component.New(component.LinearProgress)
	if err != nil {
		return nil, err
	}
	return &linearProgress{newLP, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (lp *linearProgress) HTML() string {
	return lp.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (lp *linearProgress) SetHTML(html string) {
	lp.html = html
}
