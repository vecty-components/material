package temporarydrawer // import "agamigo.io/material/component/temporarydrawer"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type TD interface {
	component.C
}

type temporaryDrawer struct {
	component.C
	html string
}

func New() (c TD, err error) {
	newTD, err := component.New(component.TemporaryDrawer)
	if err != nil {
		return nil, err
	}
	return &temporaryDrawer{newTD, defaultHTML}, err
}

func (td *temporaryDrawer) HTML() string {
	return ``
}

func (td *temporaryDrawer) SetHTML(html string) {
	td.html = html
}
