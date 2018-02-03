package linearprogress // import "agamigo.io/material/component/linearprogress"

import (
	"agamigo.io/material/component"
)

const (
	defaultHTML = `
	`
)

type LP interface {
	component.C
}

type linearProgress struct {
	component.C
	html string
}

func New() (c LP, err error) {
	newLP, err := component.New(component.LinearProgress)
	if err != nil {
		return nil, err
	}
	return &linearProgress{newLP, defaultHTML}, err
}

func (lp *linearProgress) HTML() string {
	return ``
}

func (lp *linearProgress) SetHTML(html string) {
	lp.html = html
}
