package applyer

import (
	"reflect"
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type nativeInputer interface {
	NativeInput() (*vecty.HTML, string)
}

func CSSOnly() vecty.Applyer {
	return vecty.Property("vecty-material-css-only", true)
}

func IsCSSOnly(h *vecty.HTML) bool {
	p := findProp("vecty-material-css-only", h)
	if p.IsNull() || p.IsUndefined() {
		return false
	}
	return p.Bool()
}

func FindID(moc vecty.MarkupOrChild) string {
	switch t := moc.(type) {
	case *vecty.MarkupList:
		return FindID(*t)
	case nativeInputer:
		_, id := t.NativeInput()
		return id
	case vecty.Applyer:
		d := elem.Div()
		t.Apply(d)
		return FindID(d)
	case *vecty.HTML:
		id := findProp("id", t)
		if id.IsNull() || id.IsUndefined() {
			return ""
		}
		return id.String()
	case vecty.Component:
		if h, ok := t.Render().(*vecty.HTML); ok {
			return FindID(h)
		}
	}
	return ""
}

func StartRipple(h *vecty.HTML) {
	p := findProp("vecty-material-ripple", h)
	if p.IsNull() || p.IsUndefined() {
		return
	}
	p.Get("Start").Invoke()
}

func findProp(key string, html *vecty.HTML) js.Value {
	h := reflect.ValueOf(*html)
	if !h.FieldByName("properties").IsZero() {
		for _, k := range h.FieldByName("properties").MapKeys() {
			if reflect.ValueOf(key) != k {
				continue
			}

			href := h.FieldByName("properties").
				MapIndex(reflect.ValueOf(key)).Elem().String()

			return js.ValueOf(href)
		}
	}

	return js.Undefined()
}
