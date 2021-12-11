package base

import (
	"reflect"
	"unsafe"

	"github.com/hexops/vecty"
)

/*
	Represents a simplified version of an element
*/
type LinkMarkup struct {
	Child   vecty.ComponentOrHTML
	Href    string
	OnClick func(*vecty.Event)
}

func ExtractMarkupFromLink(ht vecty.ComponentOrHTML) *LinkMarkup {
	switch ht.(type) {
	case vecty.Component, vecty.List:
		return &LinkMarkup{Child: ht}
	}

	html := ht.(*vecty.HTML)
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
