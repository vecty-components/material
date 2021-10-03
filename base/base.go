package base

import (
	"reflect"
	"unsafe"

	"github.com/hexops/vecty"
	"github.com/vecty-material/material/base/applyer"
	"github.com/vecty-material/material/components/base"
)

type MDC struct {
	Component   base.ComponentStartStopper
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

/*
	Represents a simplified version of an element
*/
type LinkMarkup struct {
	Child          vecty.ComponentOrHTML
	Href           string
	OnClick        func(*vecty.Event)
	PreventDefault bool
}

func ExtractMarkupFromLink(html *vecty.HTML) *LinkMarkup {
	sm := &LinkMarkup{}

	h := reflect.ValueOf(*html)
	tag := h.FieldByName("tag").String()
	href := ""
	if !h.FieldByName("properties").IsZero() {
		href = h.FieldByName("properties").
			MapIndex(reflect.ValueOf("href")).Elem().String()
	}

	if tag == "a" && href != "" {
		sm.Href = href
	}

	for i := 0; i < h.FieldByName("eventListeners").Len(); i++ {
		if h.FieldByName("eventListeners").Index(i).
			Elem().FieldByName("Name").String() != "click" {
			continue
		}

		sm.OnClick = *(*func(*vecty.Event))(
			unsafe.Pointer(
				h.FieldByName("eventListeners").Index(i).
					Elem().FieldByName("Listener").UnsafeAddr(),
			),
		)

		sm.PreventDefault = h.FieldByName("eventListeners").Index(i).
			Elem().FieldByName("callPreventDefault").Bool()

		break
	}

	if tag == "a" && h.FieldByName("children").Len() > 0 {
		sm.Child = *(*vecty.ComponentOrHTML)(
			unsafe.Pointer(
				h.FieldByName("children").Index(0).UnsafeAddr(),
			),
		)
	} else {
		sm.Child = html
	}

	return sm
}
