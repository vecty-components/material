package main

import (
	"time"

	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/dialog"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

const (
	MDC_CSS = "https://unpkg.com/material-components-web@0.30.0/dist/material-components-web.min.css"
)

// PageView is our main page component.
type PageView struct {
	vecty.Core
	cb           *checkbox.CB
	dialog       *dialog.D
	cbStatus     string
	dialogStatus string
}

type cbStatusText struct {
	*checkbox.CB
	status string
}

func main() {
	vecty.SetTitle("Vecty Material Components")
	vecty.AddStylesheet(MDC_CSS)
	pv := &PageView{}
	pv.cb = checkbox.New()
	pv.cbStatus = cbStatus(pv.cb)
	pv.dialog = dialog.New()
	pv.dialog.CancelHandler = event.Click(
		func(e *vecty.Event) {
			pv.dialogStatus = "Denied"
			vecty.Rerender(pv)
		},
	)
	pv.dialog.AcceptHandler = event.Click(
		func(e *vecty.Event) {
			pv.dialogStatus = "Accepted"
			vecty.Rerender(pv)
		},
	)
	vecty.RenderBody(pv)
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	dialogHTML := elem.Div(
		p.dialog,
	)

	dialogBtnHTML := elem.Div(
		elem.Button(
			vecty.Markup(
				vecty.Class("mdc-button"),
				vecty.Class("mdc-button--raised"),
				prop.Type(prop.TypeButton),
				event.Click(p.dialog.OpenHandler),
			),
			vecty.Text("Show Dialog"),
		),
		vecty.Text(" Dialog Status: "+p.dialogStatus),
	)

	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-form-field"),
				event.Change(
					func(e *vecty.Event) {
						p.cbStatus = cbStatus(p.cb)
						vecty.Rerender(p)
					},
				),
			),
			p.cb,
			elem.Label(
				vecty.Markup(
					prop.For("native-js-checkbox"),
				),
				vecty.Text(p.cbStatus),
			),
		),
		dialogHTML,
		dialogBtnHTML,
	)
}

func cbStatus(cb *checkbox.CB) string {
	s := "Checkbox Status - [ID: \"" + cb.String() + "\"]"
	if cb.Component().Object == nil {
		return s + " [Value: \"\"]"
	}
	if cb.Value != "" {
		s = s + " [Value: \"" + cb.Value + "\"]"
	}
	if cb.Checked {
		s = s + " [Checked]"
	}
	if cb.Disabled {
		s = s + " [Disabled]"
	}
	if cb.Indeterminate {
		s = s + " [Indeterminate]"
	}
	return s
}

func (p *PageView) Mount() {
	go testCB(p)
}

func testCB(p *PageView) {
	for _ = range time.Tick(5 * time.Second) {
		switch {
		case !p.cb.Checked:
			p.cb.Checked = true
		case !p.cb.Disabled:
			p.cb.Disabled = true
		case !p.cb.Indeterminate:
			p.cb.Indeterminate = true
		default:
			p.cb.Checked = false
			p.cb.Disabled = false
			p.cb.Indeterminate = false
		}
		p.cbStatus = cbStatus(p.cb)
		vecty.Rerender(p)
	}
}
