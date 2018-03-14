package main

import (
	"agamigo.io/vecty-material/base"
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
					&base.Props{
						ID: "mdc-dialog-hero",
						Markup: []vecty.Applyer{
							vecty.Class("catalog-dialog-demo")},
					},
					&dialog.State{
						Header: "Are you happy?",
						Body: vecty.Text("Please check the left and right side " +
							"of this element for fun."),
						Open:          true,
						NoBackdrop:    true,
						CancelHandler: func(e *vecty.Event) {},
						AcceptHandler: func(e *vecty.Event) {},
					},
				),
			),
			elem.Div(
				vecty.Markup(vecty.Class("demo-body")),
				c.newDemoDialog(
					&base.Props{
						ID: "mdc-dialog-default",
					},
					&dialog.State{
						Header: "Use Google's location service?",
						Body: vecty.Text("Let Google help apps determine " +
							"location. This means sending anonymous location " +
							"data to Google, even when no apps are running."),
						Role: "alertdialog",
					},
				),
				c.newDemoDialog(
					&base.Props{
						ID: "mdc-dialog-colored-footer-buttons",
					},
					&dialog.State{
						Header: "Use Google's location service?",
						Body: vecty.Text("Let Google help apps determine " +
							"location. This means sending anonymous location " +
							"data to Google, even when no apps are running."),
						Role: "alertdialog",
					},
				),
				c.newDemoDialog(
					&base.Props{
						ID: "mdc-dialog-with-list",
					},
					&dialog.State{
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
					},
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				button.New(
					&base.Props{ID: "default-dialog-activation"},
					&button.State{
						Label:  vecty.Text("Show Dialog"),
						Raised: true,
						ClickHandler: func(e *vecty.Event) {
							c.dialogs["mdc-dialog-default"].Open = true
							vecty.Rerender(c.dialogs["mdc-dialog-default"])
						},
					},
				),
				button.New(
					&base.Props{ID: "colored-footer-button-dialog-activation"},
					&button.State{
						Label:  vecty.Text("Show Colored Footer Button Dialog"),
						Raised: true,
						ClickHandler: func(e *vecty.Event) {
							c.dialogs["mdc-dialog-colored-footer-buttons"].Open = true
							vecty.Rerender(
								c.dialogs["mdc-dialog-colored-footer-buttons"])
						},
					},
				),
				button.New(
					&base.Props{ID: "dialog-with-list-activation"},
					&button.State{
						Label:  vecty.Text("Show Scrolling Dialog"),
						Raised: true,
						ClickHandler: func(e *vecty.Event) {
							c.dialogs["mdc-dialog-with-list"].Open = true
							vecty.Rerender(c.dialogs["mdc-dialog-with-list"])
						},
					},
				),
				formfield.New(nil,
					&formfield.State{
						Label: "Toggle RTL",
						Input: checkbox.New(
							&base.Props{
								ID: "toggle-rtl",
								Markup: []vecty.Applyer{
									event.Change(func(e *vecty.Event) {
										checked := e.Target.Get("checked").Bool()
										for _, v := range c.dialogs {
											el := v.Props.Element.Node()
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
							nil,
						),
					},
				),
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

func (c *dialogDemoView) newDemoDialog(p *base.Props, s *dialog.State) *dialog.D {
	if p.ID == "" {
		panic("newDemoDialog got a Props with empty ID.")
	}
	c.dialogs[p.ID] = dialog.New(p, s)
	return c.dialogs[p.ID]
}
