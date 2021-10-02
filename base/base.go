package base

import (
	"reflect"

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
	Returns:
		- If the element is an anchor, the associated href
		- The associated event listener
		- Whether PreventDefault is enabled for such listener
*/
func ExtractLinkAndListeners(html *vecty.HTML) (
	vecty.ComponentOrHTML, string, func(*vecty.Event), bool,
) {
	h := reflect.ValueOf(*html)
	tag := h.FieldByName("tag").String()
	properties := h.FieldByName("properties").
		Interface().(map[string]interface{})
	listeners := h.FieldByName("eventListeners").
		Interface().([]*vecty.EventListener)
	children := h.FieldByName("children").
		Interface().([]vecty.ComponentOrHTML)

	var OnClick *vecty.EventListener
	for _, listener := range listeners {
		if listener.Name != "OnClick" {
			continue
		}

		OnClick = listener
	}

	var callPreventDefault bool
	var f func(*vecty.Event)
	if OnClick != nil {
		callPreventDefault = reflect.ValueOf(OnClick).
			FieldByName("callPreventDefault").Bool()
	}

	var href string
	if h, ok := properties["href"]; tag == "a" && ok {
		href = h.(string)
	}

	return children[0], href, f, callPreventDefault
}
