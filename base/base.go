package base

import (
	"agamigo.io/material/base"
	"agamigo.io/material/ripple"
	"github.com/gopherjs/vecty"
)

type MDCRooter interface {
	MDCRoot() *Base
}

type Base struct {
	MDC       base.ComponentStartStopper
	ID        string
	Basic     bool
	Element   *vecty.HTML
	HasRipple bool
	RippleC   *ripple.R
}

func (b *Base) Mount() {
	var isRippler bool
	if b.MDC == nil || b.Basic {
		if !b.HasRipple {
			return
		}
		b.MDC = ripple.New()
		isRippler = true
	}
	err := b.MDC.Start(b.Element.Node())
	if err != nil {
		panic(err)
	}
	if isRippler || !b.HasRipple || b.Basic {
		return
	}
	if b.RippleC != nil {
		print("Warning: Stopping lingering Rippler. " +
			"Use Stop() to stop vecty-material components..")
		err = b.RippleC.Stop()
		if err != nil {
			print(err)
		}
		b.RippleC = nil
	}
	b.RippleC = ripple.New()
	err = b.RippleC.Start(b.Element.Node())
	if err != nil {
		print(err)
	}
}

func (b *Base) Unmount() {
	if b.RippleC != nil {
		err := b.RippleC.Stop()
		if err != nil {
			print(err)
		}
	}
	if b.MDC != nil {
		err := b.MDC.Stop()
		if err != nil {
			panic(err)
		}
	}
}
