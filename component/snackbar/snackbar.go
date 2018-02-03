package snackbar // import "agamigo.io/material/component/snackbar"

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

type snackbar struct {
	component.C
	html string
}

func New() (c S, err error) {
	newS, err := component.New(component.Snackbar)
	if err != nil {
		return nil, err
	}
	return &snackbar{newS, defaultHTML}, err
}

func (s *snackbar) HTML() string {
	return ``
}

func (s *snackbar) SetHTML(html string) {
	s.html = html
}
