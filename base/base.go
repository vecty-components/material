package base

import (
	"agamigo.io/material/base"
	"agamigo.io/vecty-material/base/applyer"
	"github.com/gopherjs/vecty"
)

type MDCRooter interface {
	MDCRoot() *MDCRoot
}

type MDCRoot struct {
	MDC     base.ComponentStartStopper
	Element *vecty.HTML
}

func (b *MDCRoot) Mount() {
	applyer.StartRipple(b.Element)
	switch {
	case b.MDC == nil:
		fallthrough
	case applyer.IsCSSOnly(b.Element):
		return
	}
	err := b.MDC.Start(b.Element.Node())
	if err != nil {
		panic(err)
	}
}

func (b *MDCRoot) Unmount() {
	if b.MDC != nil {
		err := b.MDC.Stop()
		if err != nil {
			panic(err)
		}
	}
}

// MarkupOnly returns the vecty.MarkupList contained in moc, or nil if none is
// found. It also returns nil if moc is a vecty.List that contains one or more
// vecty.ComponentOrHTML.  If nil is returned, it is then safe to assert the
// type of moc as a vecty.ComponentOrHTML.
func MarkupOnly(moc vecty.MarkupOrChild) *vecty.MarkupList {
	// TODO: handle vecty.List
	switch t := moc.(type) {
	case vecty.ComponentOrHTML:
		return nil
	case vecty.MarkupList:
		return &t
	}
	return nil
}
