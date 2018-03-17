package base

import (
	mbase "agamigo.io/material/base"
	"agamigo.io/material/ripple"
	"github.com/gopherjs/vecty"
)

type Base struct {
	vecty.Core
	*Props
}

func New(p *Props, mdc mbase.ComponentStartStopper) *Base {
	if p == nil {
		p = &Props{}
	}
	b := &Base{Props: p}
	b.Props.mdc = mdc
	return b
}

func (b *Base) Render(h *vecty.HTML) vecty.ComponentOrHTML {
	b.Props.Element = h
	return h
}

func (b *Base) Mount() {
	var isRippler bool
	if b.Props.mdc == nil || b.mdc.Component().Basic {
		if !b.Props.Ripple {
			return
		}
		b.Props.mdc = ripple.New()
		isRippler = true
	}
	err := b.Props.mdc.Start(b.Props.Element.Node())
	if err != nil {
		panic(err)
	}
	if isRippler || !b.Props.Ripple ||
		b.Props.mdc.Component().MDCState.Basic {
		return
	}
	if b.Props.ripple != nil {
		print("Warning: Stopping lingering Rippler. " +
			"Use Stop() to stop vecty-material components..")
		err = b.Props.ripple.Stop()
		if err != nil {
			print(err)
		}
		b.Props.ripple = nil
	}
	b.Props.ripple = ripple.New()
	err = b.Props.ripple.Start(b.Props.Element.Node())
	if err != nil {
		print(err)
	}
}

func (b *Base) Unmount() {
	if b.Props.ripple != nil {
		err := b.Props.ripple.Stop()
		if err != nil {
			print(err)
		}
	}
	err := b.Props.mdc.Stop()
	if err != nil {
		panic(err)
	}
}
