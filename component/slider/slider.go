package slider

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type S interface {
	component.C
}

type slider struct {
	component.C
	html string
}

func New() (c S, err error) {
	newS, err := component.New(component.Slider)
	if err != nil {
		return nil, err
	}
	return &slider{newS, defaultHTML}, err
}

func (s *slider) HTML() string {
	return ``
}

func (s *slider) SetHTML(html string) {
	s.html = html
}
