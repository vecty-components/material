// The textfield package implements a material textfield component.
//
// See: https://material.io/components/web/catalog/input-controls/text-field/
package textfield // import "agamigo.io/material/component/textfield"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-text-field">
  <input type="text" id="my-text-field" class="mdc-text-field__input">
  <label class="mdc-text-field__label" for="my-text-field">Hint text</label>
  <div class="mdc-text-field__bottom-line"></div>
</div>`
)

// T is the interface for a material textfield component.
type T interface {
	component.C
}

// textField is the internal implementation of T made available publicly via
// New().
type textField struct {
	component.C
	html string
}

// New creates a material textfield component that implement the T interface.
// It is a wrapper around component.New.
func New() (c T, err error) {
	newT, err := component.New(component.TextField)
	if err != nil {
		return nil, err
	}
	return &textField{newT, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (t *textField) HTML() string {
	return t.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (t *textField) SetHTML(html string) {
	t.html = html
}
