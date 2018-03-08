package base

import (
	mbase "agamigo.io/material/base"
	"agamigo.io/material/ripple"
	"github.com/gopherjs/vecty"
)

type Props struct {
	mdc      mbase.ComponentStartStopper
	ID       string
	Markup   []vecty.Applyer
	Element  *vecty.HTML
	NoRipple bool
	ripple   *ripple.R
}
