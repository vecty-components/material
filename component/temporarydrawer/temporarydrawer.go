// The temporarydrawer package implements a material temporarydrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package temporarydrawer // import "agamigo.io/material/component/temporarydrawer"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<aside class="mdc-drawer mdc-drawer--temporary mdc-typography">
  <nav class="mdc-drawer__drawer">
    <header class="mdc-drawer__header">
      <div class="mdc-drawer__header-content">
        Header here
      </div>
    </header>
    <nav id="icon-with-text-demo" class="mdc-drawer__content mdc-list">
      <a class="mdc-list-item mdc-list-item--activated" href="#">
        <i class="material-icons mdc-list-item__graphic" aria-hidden="true">inbox</i>Inbox
      </a>
      <a class="mdc-list-item" href="#">
        <i class="material-icons mdc-list-item__graphic" aria-hidden="true">star</i>Star
      </a>
    </nav>
  </nav>
</aside>`
)

// TD is the interface for a material temporarydrawer component.
type TD interface {
	component.C
}

// temporarydrawer is the internal implementation of TD made available publicly via
// New().
type temporaryDrawer struct {
	component.C
	html string
}

// New creates a material temporarydrawer component that implement the TD interface.
// It is a wrapper around component.New.
func New() (c TD, err error) {
	newTD, err := component.New(component.TemporaryDrawer)
	if err != nil {
		return nil, err
	}
	return &temporaryDrawer{newTD, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (td *temporaryDrawer) HTML() string {
	return td.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (td *temporaryDrawer) SetHTML(html string) {
	td.html = html
}
