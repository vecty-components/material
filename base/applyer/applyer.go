package applyer

import (
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
	if p.IsNull() {
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
		if id.IsNull() {
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
	if p.IsNull() {
		return
	}
	p.Get("Start").Invoke()
}

func findProp(key string, h *vecty.HTML) js.Value {
	k := h.Node()
	if k.IsUndefined() {
		return js.Null()
	}
	k = k.Get("properties")
	if k.IsUndefined() {
		return js.Null()
	}
	k = k.Get("$" + key)
	if k.IsUndefined() {
		return js.Null()
	}
	return k.Get("v").Get("$val")
}
