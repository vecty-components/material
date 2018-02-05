// The gridlist package implements a material gridlist component.
//
// See: https://material.io/components/web/catalog/grid-lists/
package gridlist // import "agamigo.io/material/component/gridlist"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-grid-list">
  <ul class="mdc-grid-list__tiles">
    <li class="mdc-grid-tile">
      <div class="mdc-grid-tile__primary">
        <img class="mdc-grid-tile__primary-content" src="my-image.jpg" />
      </div>
      <span class="mdc-grid-tile__secondary">
        <span class="mdc-grid-tile__title">Title</span>
      </span>
    </li>
  </ul>
</div>`
)

// GL is the interface for a material gridlist component.
type GL interface {
	component.C
}

// gridlist is the internal implementation of GL made available publicly via
// New().
type gridList struct {
	component.C
	html string
}

// New creates a material gridlist component that implement the GL interface.
// It is a wrapper around component.New.
func New() (c GL, err error) {
	newGL, err := component.New(component.GridList)
	if err != nil {
		return nil, err
	}
	return &gridList{newGL, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (gl *gridList) HTML() string {
	return gl.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (gl *gridList) SetHTML(html string) {
	gl.html = html
}
