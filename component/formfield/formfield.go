// The formfield package implements a material formfield component.
//
// See: https://material.io/components/web/catalog/input-controls/form-fields/
package formfield // import "agamigo.io/material/component/formfield"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<div class="mdc-form-field">
  <div class="mdc-checkbox">
    <input type="checkbox" id="my-checkbox"
    class="mdc-checkbox__native-control"/>
    </div>
  </div>
  <label for="my-checkbox" id="my-checkbox-label">This is my checkbox</label>
</div>`
)

// FF is the interface for a material formfield component.
type FF interface {
	component.C
}

// formField is the internal implementation of FF made available publicly via
// New().
type formField struct {
	component.C
	html string
}

// New creates a material formfield component that implement the FF interface.
// It is a wrapper around component.New.
func New() (c FF, err error) {
	newFF, err := component.New(component.FormField)
	if err != nil {
		return nil, err
	}
	return &formField{newFF, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (ff *formField) HTML() string {
	return ff.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (ff *formField) SetHTML(html string) {
	ff.html = html
}
