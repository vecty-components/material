package toolbar

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

type toolbar struct {
	component.C
	html string
}

func New() (c T, err error) {
	newT, err := component.New(component.Toolbar)
	if err != nil {
		return nil, err
	}
	return &toolbar{newT, defaultHTML}, err
}

func (t *toolbar) HTML() string {
	return ``
}

func (t *toolbar) SetHTML(html string) {
	t.html = html
}
