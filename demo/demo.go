package main

import (
	"log"
	"time"

	mdccheckbox "agamigo.io/material/checkbox"
	"agamigo.io/vecty-material/checkbox"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type checkboxes []checkbox.CB

var cbs checkboxes

func main() {
	cb, err := checkbox.New()
	if err != nil {
		log.Fatalf("Failed to create checkbox component: %v", err)
	}
	cbs = append(cbs, cb)
	vecty.SetTitle("Material Components Go")
	vecty.RenderBody(&PageView{})
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("mcw.css")
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-form-field"),
			),
			cbs[0],
			elem.Label(
				vecty.Markup(
					prop.For("native-js-checkbox"),
				),
				vecty.Text("Default checkbox"),
			),
		),
	)
}

func (p *PageView) Mount() {
	go cbs.testCB()
}

func (cbs checkboxes) testCB() {
	for _ = range time.Tick(1 * time.Second) {
		for _, c := range cbs {
			s := c.State()
			print(s)
			if s == mdccheckbox.INDETERMINATE_DISABLED {
				c.SetState(mdccheckbox.UNCHECKED)
				continue
			}
			c.SetState(s + mdccheckbox.DISABLED)
		}
	}
}
