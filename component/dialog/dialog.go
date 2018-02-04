// The dialog package implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/material/component/dialog"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<aside id="my-mdc-dialog"
  class="mdc-dialog"
  role="alertdialog"
  aria-labelledby="my-mdc-dialog-label"
  aria-describedby="my-mdc-dialog-description">
  <div class="mdc-dialog__surface">
    <header class="mdc-dialog__header">
      <h2 id="my-mdc-dialog-label" class="mdc-dialog__header__title">
        Dialog header
      </h2>
    </header>
    <section id="my-mdc-dialog-description" class="mdc-dialog__body">
      Dialog description
    </section>
    <footer class="mdc-dialog__footer">
      <button type="button" class="mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--cancel">Decline</button>
      <button type="button" class="mdc-button mdc-dialog__footer__button mdc-dialog__footer__button--accept">Accept</button>
    </footer>
  </div>
  <div class="mdc-dialog__backdrop"></div>
</aside>`
)

// D is the interface for a material dialog component.
type D interface {
	component.C
}

// dialog is the internal implementation of D made available publicly via
// New().
type dialog struct {
	component.C
	html string
}

// New creates a material dialog component that implement the D interface.
// It is a wrapper around component.New.
func New() (c D, err error) {
	newD, err := component.New(component.Dialog)
	if err != nil {
		return nil, err
	}
	return &dialog{newD, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (d *dialog) HTML() string {
	return d.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (d *dialog) SetHTML(html string) {
	d.html = html
}
