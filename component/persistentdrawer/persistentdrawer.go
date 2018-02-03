package persistentdrawer

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type PD interface {
	component.C
}

type persistentDrawer struct {
	component.C
	html string
}

func New() (c PD, err error) {
	newPD, err := component.New(component.PersistentDrawer)
	if err != nil {
		return nil, err
	}
	return &persistentDrawer{newPD, defaultHTML}, err
}

func (pd *persistentDrawer) HTML() string {
	return ``
}

func (pd *persistentDrawer) SetHTML(html string) {
	pd.html = html
}
