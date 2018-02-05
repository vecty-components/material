// The icontoggle package implements a material icontoggle component.
//
// See: https://material.io/components/web/catalog/buttons/icon-toggle-buttons/
package icontoggle // import "agamigo.io/material/component/icontoggle"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
<i class="mdc-icon-toggle material-icons" role="button" aria-pressed="false"
   aria-label="Add to favorites" tabindex="0"
   data-toggle-on='{"label": "Remove from favorites", "content": "favorite"}'
   data-toggle-off='{"label": "Add to favorites", "content": "favorite_border"}'>
  favorite_border
</i>`
)

// IT is the interface for a material icontoggle component.
type IT interface {
	component.C
}

// icontoggle is the internal implementation of IT made available publicly via
// New().
type iconToggle struct {
	component.C
	html string
}

// New creates a material icontoggle component that implement the IT interface.
// It is a wrapper around component.New.
func New() (c IT, err error) {
	newIT, err := component.New(component.IconToggle)
	if err != nil {
		return nil, err
	}
	return &iconToggle{newIT, defaultHTML}, err
}

// HTML implements the material component.HTMLElementer interface.
func (it *iconToggle) HTML() string {
	return it.html
}

// SetHTML implements the material component.HTMLElementer interface.
func (it *iconToggle) SetHTML(html string) {
	it.html = html
}
