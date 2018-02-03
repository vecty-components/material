package textfield

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

type textfield struct {
	component.C
	html string
}

func New() (c T, err error) {
	newT, err := component.New(component.Textfield)
	if err != nil {
		return nil, err
	}
	return &textfield{newT, defaultHTML}, err
}

func (t *textfield) HTML() string {
	return ``
}

func (t *textfield) SetHTML(html string) {
	t.html = html
}
