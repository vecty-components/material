package dialog

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type D interface {
	component.C
}

type dialog struct {
	component.C
	html string
}

func New() (c D, err error) {
	newD, err := component.New(component.Dialog)
	if err != nil {
		return nil, err
	}
	return &dialog{newD, defaultHTML}, err
}

func (d *dialog) HTML() string {
	return ``
}

func (d *dialog) SetHTML(html string) {
	d.html = html
}
