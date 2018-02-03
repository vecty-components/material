package tabbar

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type TB interface {
	component.C
}

type tabBar struct {
	component.C
	html string
}

func New() (c TB, err error) {
	newTB, err := component.New(component.TabBar)
	if err != nil {
		return nil, err
	}
	return &tabBar{newTB, defaultHTML}, err
}

func (tb *tabBar) HTML() string {
	return ``
}

func (tb *tabBar) SetHTML(html string) {
	tb.html = html
}
