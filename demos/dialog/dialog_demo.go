package main

import (
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/dialog"
	"agamigo.io/vecty-material/formfield"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

// dialogDemo is our main page component.
type dialogDemoView struct {
	vecty.Core
	dialogs map[string]*dialog.D
}

func main() {
	vecty.RenderBody(&dialogDemoView{dialogs: make(map[string]*dialog.D, 0)})
}

// Render implements the vecty.Component interface.
func (c *dialogDemoView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(vecty.Class("mdc-typography")),
		&common.ToolbarHeader{
			Title:      "Dialog",
			Navigation: common.NavBack,
		},
		elem.Main(
			elem.Div(vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust"))),
			elem.Section(
				vecty.Markup(vecty.Class("hero")),
				c.newDemoDialog(
					&dialog.D{
						ID: "mdc-dialog-hero",
						Markup: []vecty.Applyer{
							vecty.Class("catalog-dialog-demo")},
						Header: "Are you happy?",
						Body: vecty.Text("Please check the left and right side " +
							"of this element for fun."),
						Open:       true,
						NoBackdrop: true,
						Basic:      true,
					},
				),
			),
			elem.Div(
				vecty.Markup(vecty.Class("demo-body")),
				c.newDemoDialog(
					&dialog.D{
						ID:     "mdc-dialog-default",
						Header: "Use Google's location service?",
						Body: vecty.Text("Let Google help apps determine " +
							"location. This means sending anonymous location " +
							"data to Google, even when no apps are running."),
						Role: "alertdialog",
						OnCancel: func(thisD *dialog.D, e *vecty.Event) {
							thisD.Open = false
							vecty.Rerender(thisD)
						},
						OnAccept: func(thisD *dialog.D, e *vecty.Event) {
							thisD.Open = false
							vecty.Rerender(thisD)
						},
					},
				),
				c.newDemoDialog(
					&dialog.D{
						ID:     "mdc-dialog-colored-footer-buttons",
						Header: "Use Google's location service?",
						Body: vecty.Text("Let Google help apps determine " +
							"location. This means sending anonymous location " +
							"data to Google, even when no apps are running."),
						Role: "alertdialog",
						OnCancel: func(thisD *dialog.D, e *vecty.Event) {
							thisD.Open = false
						},
						OnAccept: func(thisD *dialog.D, e *vecty.Event) {
							thisD.Open = false
						},
					},
				),
				c.newDemoDialog(
					&dialog.D{
						ID:         "mdc-dialog-with-list",
						Header:     "Choose a Ringtone",
						Role:       "alertdialog",
						Scrollable: true,
						Body: renderList(
							"None",
							"Callisto",
							"Ganymede",
							"Luna",
							"Marimba",
							"Schwifty",
							"Callisto",
							"Ganymede",
							"Luna",
							"Marimba",
							"Schwifty",
						),
						OnCancel: func(thisD *dialog.D, e *vecty.Event) {
							thisD.Open = false
						},
						OnAccept: func(thisD *dialog.D, e *vecty.Event) {
							thisD.Open = false
						},
					},
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				&button.B{
					ID:     "default-dialog-activation",
					Label:  vecty.Text("Show Dialog"),
					Raised: true,
					OnClick: func(thisB *button.B, e *vecty.Event) {
						c.dialogs["mdc-dialog-default"].Open = true
						vecty.Rerender(c.dialogs["mdc-dialog-default"])
					},
				},
				&button.B{
					ID:     "colored-footer-button-dialog-activation",
					Label:  vecty.Text("Show Colored Footer Button Dialog"),
					Raised: true,
					OnClick: func(thisB *button.B, e *vecty.Event) {
						class := "mdc-dialog-colored-footer-buttons"
						c.dialogs[class].Open = true
						vecty.Rerender(c.dialogs[class])
					},
				},
				&button.B{
					ID:     "dialog-with-list-activation",
					Label:  vecty.Text("Show Scrolling Dialog"),
					Raised: true,
					OnClick: func(thisB *button.B, e *vecty.Event) {
						c.dialogs["mdc-dialog-with-list"].Open = true
						vecty.Rerender(c.dialogs["mdc-dialog-with-list"])
					},
				},
				&formfield.FF{
					Label: "Toggle RTL",
					Input: &checkbox.CB{
						ID: "toggle-rtl",
						Markup: []vecty.Applyer{
							event.Change(func(e *vecty.Event) {
								checked := e.Target.Get("checked").Bool()
								for _, v := range c.dialogs {
									el := v.MDCRoot().Element.Node()
									if checked {
										el.Call("setAttribute",
											"dir", "rtl")
										return
									}
									el.Call("removeAttribute", "dir")
								}
							}),
						},
					},
				},
			),
		),
	)
}

func renderList(itemsText ...string) vecty.ComponentOrHTML {
	items := make(vecty.List, len(itemsText))
	for i, v := range itemsText {
		items[i] = elem.ListItem(
			vecty.Markup(vecty.Class("mdc-list-item")),
			vecty.Text(v),
		)
	}
	return elem.UnorderedList(
		vecty.Markup(vecty.Class("mdc-list")),
		items,
	)
}

func (c *dialogDemoView) newDemoDialog(d *dialog.D) *dialog.D {
	if d.ID == "" {
		panic("newDemoDialog got a Props with empty ID.")
	}
	c.dialogs[d.ID] = d
	return d
}
