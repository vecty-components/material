package tab

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type T interface {
	component.C
}

type tab struct {
	component.C
	html string
}

func New() (c T, err error) {
	newT, err := component.New(component.Tab)
	if err != nil {
		return nil, err
	}
	return &tab{newT, defaultHTML}, err
}

func (t *tab) HTML() string {
	return ``
}

func (t *tab) SetHTML(html string) {
	t.html = html
}
