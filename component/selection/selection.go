package selection // import "agamigo.io/material/component/selection"

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

type selection struct {
	component.C
	html string
}

func New() (c S, err error) {
	newS, err := component.New(component.Select)
	if err != nil {
		return nil, err
	}
	return &selection{newS, defaultHTML}, err
}

func (s *selection) HTML() string {
	return ``
}

func (s *selection) SetHTML(html string) {
	s.html = html
}
