package radio

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type R interface {
	component.C
}

type radio struct {
	component.C
	html string
}

func New() (c R, err error) {
	newR, err := component.New(component.Radio)
	if err != nil {
		return nil, err
	}
	return &radio{newR, defaultHTML}, err
}

func (r *radio) HTML() string {
	return ``
}

func (r *radio) SetHTML(html string) {
	r.html = html
}
