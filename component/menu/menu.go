package menu // import "agamigo.io/material/component/menu"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type M interface {
	component.C
}

type menu struct {
	component.C
	html string
}

func New() (c M, err error) {
	newM, err := component.New(component.Menu)
	if err != nil {
		return nil, err
	}
	return &menu{newM, defaultHTML}, err
}

func (m *menu) HTML() string {
	return ``
}

func (m *menu) SetHTML(html string) {
	m.html = html
}
