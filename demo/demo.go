package main

import (
	"log"
	"time"

	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/dialog"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.SetTitle("Vecty Material Components")

	myCheckbox, err := checkbox.New()
	if err != nil {
		log.Fatalf("Failed to create checkbox component: %v", err)
	}

	myDialog, err := dialog.New("mydialog1")
	if err != nil {
		log.Fatalf("Failed to create dialog component: %v", err)
	}
	myDialog.Label = "My Dialog"
	myDialog.Description = "This is a material dialog example."
	myDialog.Role = "alertdialog"

	vecty.RenderBody(&PageView{
		Checkbox: myCheckbox,
		Dialog:   myDialog,
	})
}

// PageView is our main page component.
type PageView struct {
	vecty.Core
	Checkbox *checkbox.CB
	Dialog   *dialog.D
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("mcw.css")
	return elem.Body(
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-form-field"),
			),
			p.Checkbox,
			elem.Label(
				vecty.Markup(
					prop.For("native-js-checkbox"),
				),
				vecty.Text("Default checkbox"),
			),
		),
		elem.Div(
			elem.Button(
				vecty.Markup(
					prop.Type(prop.TypeButton),
					vecty.Class("mdc-button"),
					vecty.Class("mdc-button--raised"),
					prop.Value("Activate Dialog"),
					event.Click(p.Dialog.OpenHandler),
				),
				vecty.Text("Show Dialog"),
			),
		),
		elem.Div(
			p.Dialog,
		),
	)
}

func (p *PageView) Mount() {
	go testCB(p.Checkbox)
}

func testCB(c *checkbox.CB) {
	for _ = range time.Tick(1 * time.Second) {
		switch {
		case c.Checked == true:
			c.Checked = false
		default:
			c.Checked = true
		}
		print(c.Checked)
		print(c.GetObject().Get("checked"))
	}
}
