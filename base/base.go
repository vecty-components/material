package base

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/vecty-components/material/base/applyer"
)

var key int = 0

func Key() string {
	key += 1
	return "material-component-" + strconv.Itoa(key)
}

type MDC struct {
	Component   ComponentStartStopper
	RootElement *vecty.HTML
}

func (b *MDC) Mount() {
	applyer.StartRipple(b.RootElement)
	switch {
	case b.Component == nil:
		fallthrough
	case applyer.IsCSSOnly(b.RootElement):
		return
	}
	err := b.Component.Start(b.RootElement.Node())
	if err != nil {
		panic(err)
	}
}

func (b *MDC) Unmount() {
	if b.Component != nil {
		err := b.Component.Stop()
		if err != nil {
			// panic(err)
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
