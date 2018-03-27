package main

import (
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/demos/common"
	dcommon "agamigo.io/vecty-material/demos/drawer/common"
	"agamigo.io/vecty-material/drawer"
	"agamigo.io/vecty-material/formfield"
	"agamigo.io/vecty-material/radio"
	"github.com/gopherjs/gopherjs/js"
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
					formfield.New(nil, &formfield.State{
						Label: "Default",
						Input: vecty.List{radio.New(
							&base.Props{
								ID: "theme-radio-default",
							},
							&radio.State{
								Name:    "theme",
								Checked: true,
								ChangeHandler: func(thisR *radio.R,
									e *vecty.Event) {
									d := js.Global.Get("window").Get("document")
									dd := d.Call("querySelector", ".demo-drawer")
									dd.Get("classList").Call("remove",
										"demo-drawer--accessible")
									dd.Get("classList").Call("remove",
										"demo-drawer--custom")
								},
							},
						)}},
					),
					formfield.New(nil, &formfield.State{
						Label: "Custom Theme",
						Input: vecty.List{radio.New(
							&base.Props{
								ID: "theme-radio-custom",
							},
							&radio.State{
								Name: "theme",
								ChangeHandler: func(thisR *radio.R,
									e *vecty.Event) {
									d := js.Global.Get("window").Get("document")
									dd := d.Call("querySelector", ".demo-drawer")
									dd.Get("classList").Call("remove",
										"demo-drawer--accessible")
									dd.Get("classList").Call("add",
										"demo-drawer--custom")
								},
							},
						)}},
					),
					formfield.New(nil, &formfield.State{
						Label: "Accessible Theme",
						Input: vecty.List{radio.New(
							&base.Props{
								ID: "theme-radio-accessible",
							},
							&radio.State{
								Name: "theme",
								ChangeHandler: func(thisR *radio.R,
									e *vecty.Event) {
									d := js.Global.Get("window").Get("document")
									dd := d.Call("querySelector", ".demo-drawer")
									dd.Get("classList").Call("remove",
										"demo-drawer--custom")
									dd.Get("classList").Call("add",
										"demo-drawer--accessible")
								},
							},
						)}},
					),
				),
				elem.Div(vecty.Markup(vecty.Class("extra-content-wrapper")),
					button.New(
						&base.Props{
							Markup: []vecty.Applyer{vecty.Class(
								"demo-toolbar-example-heading__rtl-toggle-button"),
							},
						},
						&button.State{
							Label:   vecty.Text("Toggle RTL"),
							Stroked: true,
							Dense:   true,
							ClickHandler: func(thisB *button.B,
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
				),
				elem.Div(vecty.Markup(vecty.Class("extra-content-wrapper")),
					button.New(
						&base.Props{ID: "toggle-wide"},
						&button.State{
							Label:   vecty.Text("Toggle extra-wide content"),
							Stroked: true,
							Dense:   true,
							ClickHandler: func(thisB *button.B,
								e *vecty.Event) {
								s := ewc.Node().Get("style")
								if s.Get("display").String() == "none" {
									s.Set("display", "")
									return
								}
								s.Set("display", "none")
							}}),
					ewc,
				),
				elem.Div(vecty.Markup(vecty.Class("extra-content-wrapper")),
					button.New(
						&base.Props{ID: "toggle-tall"},
						&button.State{
							Label:   vecty.Text("Toggle extra-tall content"),
							Stroked: true,
							Dense:   true,
							ClickHandler: func(thisB *button.B,
								e *vecty.Event) {
								s := etc.Node().Get("style")
								if s.Get("display").String() == "none" {
									s.Set("display", "")
									return
								}
								s.Set("display", "none")
							}}),
					etc,
				),
			),
		),
	)
	return c.body
}
