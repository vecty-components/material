package ripple

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

type ripple struct {
	component.C
	html string
}

func New() (c R, err error) {
	newR, err := component.New(component.Ripple)
	if err != nil {
		return nil, err
	}
	return &ripple{newR, defaultHTML}, err
}

func (r *ripple) HTML() string {
	return ``
}

func (r *ripple) SetHTML(html string) {
	r.html = html
}
