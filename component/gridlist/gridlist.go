package gridlist

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type GL interface {
	component.C
}

type gridList struct {
	component.C
	html string
}

func New() (c GL, err error) {
	newGL, err := component.New(component.GridList)
	if err != nil {
		return nil, err
	}
	return &gridList{newGL, defaultHTML}, err
}

func (gl *gridList) HTML() string {
	return ``
}

func (gl *gridList) SetHTML(html string) {
	gl.html = html
}
