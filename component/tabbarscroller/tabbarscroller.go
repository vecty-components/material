package tabBarScroller

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type TBS interface {
	component.C
}

type tabBarScroller struct {
	component.C
	html string
}

func New() (c TBS, err error) {
	newTBS, err := component.New(component.TabBarScroller)
	if err != nil {
		return nil, err
	}
	return &tabBarScroller{newTBS, defaultHTML}, err
}

func (tbs *tabBarScroller) HTML() string {
	return ``
}

func (tbs *tabBarScroller) SetHTML(html string) {
	tbs.html = html
}
