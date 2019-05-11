package main

import (
	"github.com/vecty-material/vecty-material/button"
	"github.com/vecty-material/vecty-material/demos/common"
	dcommon "github.com/vecty-material/vecty-material/demos/drawer/common"
	"github.com/vecty-material/vecty-material/drawer"
	"github.com/vecty-material/vecty-material/formfield"
	"github.com/vecty-material/vecty-material/radio"
	"github.com/gopherjs/gopherwasm/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type drawerDemoView struct {
	vecty.Core
	body *vecty.HTML
}

func main() {
	vecty.RenderBody(&drawerDemoView{})
}

func (c *drawerDemoView) Render() vecty.ComponentOrHTML {
	ewc := elem.Div(
		vecty.Markup(
			prop.ID("extra-wide-content"),
			vecty.Style("display", "none"),
			vecty.Class("mdc-elevation--z2"),
			vecty.UnsafeHTML("&nbsp;"),
		))
	etc := elem.Div(
		vecty.Markup(
			prop.ID("extra-tall-content"),
			vecty.Style("display", "none"),
			vecty.Class("mdc-elevation--z2"),
			vecty.UnsafeHTML("&nbsp;"),
		))
	c.body = elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
			vecty.Class("demo-body"),
		),
		&common.ToolbarHeader{
			Title:      "Permanent Drawer Below Toolbar",
			Navigation: common.NavNone,
		},
		elem.Div(
			vecty.Markup(
				vecty.Class("demo-content"),
				vecty.Class("mdc-toolbar-fixed-adjust"),
			),
			dcommon.NewDemoDrawer(drawer.Permanent),
			elem.Main(vecty.Markup(vecty.Class("demo-main")),
				elem.Heading1(
					vecty.Markup(vecty.Class("mdc-typography--display1")),
					vecty.Text("Permanent Drawer")),
				elem.Paragraph(
					vecty.Markup(vecty.Class("mdc-typography--body1")),
					vecty.Text("It sits to the left of this content."),
				),
				elem.Div(vecty.Markup(prop.ID("demo-radio-buttons")),
					&formfield.FF{
						Label: "Default",
						Input: vecty.List{
							&radio.R{
								Root: vecty.Markup(
									prop.ID("theme-radio-default"),
								),
								Name:    "theme",
								Checked: true,
								OnChange: func(thisR *radio.R,
									e *vecty.Event) {
									d := js.Global().Get("window").Get("document")
									dd := d.Call("querySelector", ".demo-drawer")
									dd.Get("classList").Call("remove",
										"demo-drawer--accessible")
									dd.Get("classList").Call("remove",
										"demo-drawer--custom")
								},
							},
						},
					},
					&formfield.FF{
						Label: "Custom Theme",
						Input: vecty.List{
							&radio.R{
								Root: vecty.Markup(
									prop.ID("theme-radio-custom"),
								),
								Name: "theme",
								OnChange: func(thisR *radio.R,
									e *vecty.Event) {
									d := js.Global().Get("window").Get("document")
									dd := d.Call("querySelector", ".demo-drawer")
									dd.Get("classList").Call("remove",
										"demo-drawer--accessible")
									dd.Get("classList").Call("add",
										"demo-drawer--custom")
								},
							},
						},
					},
					&formfield.FF{
						Label: "Accessible Theme",
						Input: vecty.List{
							&radio.R{
								Root: vecty.Markup(
									prop.ID("theme-radio-accessible"),
								),
								Name: "theme",
								OnChange: func(thisR *radio.R,
									e *vecty.Event) {
									d := js.Global().Get("window").Get("document")
									dd := d.Call("querySelector", ".demo-drawer")
									dd.Get("classList").Call("remove",
										"demo-drawer--custom")
									dd.Get("classList").Call("add",
										"demo-drawer--accessible")
								},
							},
						},
					},
				),
				elem.Div(vecty.Markup(vecty.Class("extra-content-wrapper")),
					&button.B{
						Root: vecty.Markup(vecty.Class(
							"demo-toolbar-example-heading__rtl-toggle-button"),
						),
						Label:    vecty.Text("Toggle RTL"),
						Outlined: true,
						Dense:    true,
						OnClick: func(thisB *button.B,
							e *vecty.Event) {
							b := c.body.Node()
							if b.Call("getAttribute",
								"dir").String() == "rtl" {
								b.Call("setAttribute", "dir", "ltr")
								return
							}
							b.Call("setAttribute", "dir", "rtl")
						},
					},
				),
				elem.Div(vecty.Markup(vecty.Class("extra-content-wrapper")),
					&button.B{
						Root: vecty.Markup(
							prop.ID("toggle-wide"),
						),
						Label:    vecty.Text("Toggle extra-wide content"),
						Outlined: true,
						Dense:    true,
						OnClick: func(thisB *button.B,
							e *vecty.Event) {
							s := ewc.Node().Get("style")
							if s.Get("display").String() == "none" {
								s.Set("display", "")
								return
							}
							s.Set("display", "none")
						},
					},
					ewc,
				),
				elem.Div(vecty.Markup(vecty.Class("extra-content-wrapper")),
					&button.B{
						Root: vecty.Markup(
							prop.ID("toggle-tall"),
						),
						Label:    vecty.Text("Toggle extra-tall content"),
						Outlined: true,
						Dense:    true,
						OnClick: func(thisB *button.B,
							e *vecty.Event) {
							s := etc.Node().Get("style")
							if s.Get("display").String() == "none" {
								s.Set("display", "")
								return
							}
							s.Set("display", "none")
						},
					},
					etc,
				),
			),
		),
	)
	return c.body
}
