// The persistentdrawer package implements a material persistentdrawer component.
//
// See: https://material.io/components/web/catalog/drawers/
package persistentdrawer // import "agamigo.io/material/component/persistentdrawer"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<aside class="mdc-drawer mdc-drawer--persistent mdc-typography">
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

// PD is the interface for a material persistentdrawer component.
type PD interface {
	component.C
}

// persistentDrawer is the internal implementation of PD made available publicly via
// New().
type persistentDrawer struct {
	component.C
	html string
}

// New creates a material persistentdrawer component that implement the PD interface.
// It is a wrapper around component.New.
func New() (c PD, err error) {
	newPD, err := component.New(component.PersistentDrawer)
	if err != nil {
		return nil, err
	}
	return &persistentDrawer{newPD, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (pd *persistentDrawer) HTML() string {
	return pd.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (pd *persistentDrawer) SetHTML(html string) {
	pd.html = html
}
