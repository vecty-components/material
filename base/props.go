package base

import (
	mbase "agamigo.io/material/base"
	"agamigo.io/material/ripple"
	"github.com/gopherjs/vecty"
)

type Props struct {
	mdc      mbase.ComponentStartStopper
	ID       string
	Markup   vecty.MarkupList
	Element  *vecty.HTML
	NoRipple bool
	ripple   *ripple.R
}

func NewProps() *Props {
	return &Props{}
}

func (p *Props) WithID(id string) *Props {
	p.ID = id
	return p
}

func (p *Props) WithMarkup(m ...vecty.Applyer) *Props {
	p.Markup = vecty.Markup(m...)
	return p
}
