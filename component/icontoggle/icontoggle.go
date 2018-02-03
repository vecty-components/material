package icontoggle

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type IT interface {
	component.C
}

type iconToggle struct {
	component.C
	html string
}

func New() (c IT, err error) {
	newIT, err := component.New(component.IconToggle)
	if err != nil {
		return nil, err
	}
	return &iconToggle{newIT, defaultHTML}, err
}

func (it *iconToggle) HTML() string {
	return ``
}

func (it *iconToggle) SetHTML(html string) {
	it.html = html
}
