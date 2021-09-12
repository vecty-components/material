package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/lithammer/dedent"
	"github.com/vecty-material/material"
	"github.com/vecty-material/material/base/applyer"
	"github.com/vecty-material/material/button"
	"github.com/vecty-material/material/checkbox"
	"github.com/vecty-material/material/demos/common"
	"github.com/vecty-material/material/dialog"
	"github.com/vecty-material/material/formfield"
)

// dialogDemo is our main page component.
type dialogDemoView struct {
	vecty.Core
	dialogs map[string]*dialog.D
}

func main() {
	vecty.SetTitle("Dialog - Material Components Catalog")
	vecty.AddStylesheet("https://material-components-web.appspot.com/assets/dialog.css")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto+Mono")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:300,400,500")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")

	material.AddIcon("https://material-components-web.appspot.com/images/logo_components_color_2x_web_48dp.png")
	material.AddScript("https://material-components-web.appspot.com/assets/material-components-web.js")
	material.AddCSS(dedent.Dedent(`
		.demo-body {
			padding: 24px;
			margin: 0;
			box-sizing: border-box;
		}
		.demo-content > button {
			margin-bottom: 6px;
		}

		section.demo-content {
			padding: 24px;
		}

		.catalog-dialog-demo {
			position: relative;
			width: 320px;
			z-index: auto;
		}

		.dialog-container {
			display: flex;
			align-items: center;
			justify-content: center;
		}
	`))

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
				c.newDemoDialog("mdc-dialog-hero",
					&dialog.D{
						Root: vecty.Markup(
							vecty.Class("catalog-dialog-demo"),
							prop.ID("mdc-dialog-hero"),
							applyer.CSSOnly(),
						),
						Header: "Are you happy?",
						Body: vecty.Text("Please check the left and right side " +
							"of this element for fun."),
						Open:       true,
						NoBackdrop: true,
					},
				),
			),
			elem.Div(
				vecty.Markup(vecty.Class("demo-body")),
				c.newDemoDialog("mdc-dialog-default",
					&dialog.D{
						Root: vecty.Markup(
							prop.ID("mdc-dialog-default"),
						),
						Header: "Use Google's location service?",
						Body: vecty.Text("Let Google help apps determine " +
							"location. This means sending anonymous location " +
							"data to Google, even when no apps are running."),
						Role: "alertdialog",
					},
				),
				c.newDemoDialog("mdc-dialog-colored-footer-buttons",
					&dialog.D{
						Root: vecty.Markup(
							prop.ID("mdc-dialog-colored-footer-buttons"),
						),
						Header: "Use Google's location service?",
						Body: vecty.Text("Let Google help apps determine " +
							"location. This means sending anonymous location " +
							"data to Google, even when no apps are running."),
						Role: "alertdialog",
					},
				),
				c.newDemoDialog("mdc-dialog-with-list",
					&dialog.D{
						Root: vecty.Markup(
							prop.ID("mdc-dialog-with-list"),
						),
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
				&button.B{
					Root: vecty.Markup(
						prop.ID("default-dialog-activation"),
					),
					Label:  vecty.Text("Show Dialog"),
					Raised: true,
					OnClick: func(thisB *button.B, e *vecty.Event) {
						class := "mdc-dialog-default"
						if _, ok := c.dialogs[class]; ok {
							c.dialogs[class].Open = true
							vecty.Rerender(c.dialogs[class])
						}
					},
				},
				&button.B{
					Root: vecty.Markup(
						prop.ID("colored-footer-button-dialog-activation"),
					),
					Label:  vecty.Text("Show Colored Footer Button Dialog"),
					Raised: true,
					OnClick: func(thisB *button.B, e *vecty.Event) {
						class := "mdc-dialog-colored-footer-buttons"
						if _, ok := c.dialogs[class]; ok {
							c.dialogs[class].Open = true
							vecty.Rerender(c.dialogs[class])
						}
					},
				},
				&button.B{
					Root: vecty.Markup(
						prop.ID("dialog-with-list-activation"),
					),
					Label:  vecty.Text("Show Scrolling Dialog"),
					Raised: true,
					OnClick: func(thisB *button.B, e *vecty.Event) {
						class := "mdc-dialog-with-list"
						if _, ok := c.dialogs[class]; ok {
							c.dialogs[class].Open = true
							vecty.Rerender(c.dialogs[class])
						}
					},
				},
				&formfield.FF{
					Label: "Toggle RTL",
					Input: &checkbox.CB{
						Root: vecty.Markup(
							prop.ID("toggle-rtl"),
						),
						OnChange: func(thisCB *checkbox.CB, e *vecty.Event) {
							for _, v := range c.dialogs {
								el := v.MDC.RootElement.Node()
								if thisCB.Checked {
									el.Call("setAttribute",
										"dir", "rtl")
									return
								}
								el.Call("removeAttribute", "dir")
							}
						},
					},
				},
			),
		),
	)
}

// TODO: Use ul (list) componenet
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

func (c *dialogDemoView) newDemoDialog(id string, d *dialog.D) *dialog.D {
	c.dialogs[id] = d
	return d
}
