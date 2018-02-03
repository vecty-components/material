package formrfield

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type FF interface {
	component.C
}

type formField struct {
	component.C
	html string
}

func New() (c FF, err error) {
	newFF, err := component.New(component.FormField)
	if err != nil {
		return nil, err
	}
	return &formField{newFF, defaultHTML}, err
}

func (ff *formField) HTML() string {
	return ``
}

func (ff *formField) SetHTML(html string) {
	ff.html = html
}
