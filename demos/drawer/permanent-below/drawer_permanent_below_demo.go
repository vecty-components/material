package main

import (
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/drawer"
	"agamigo.io/vecty-material/formfield"
	"agamigo.io/vecty-material/radio"
	"agamigo.io/vecty-material/ul"
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
			drawer.New(
				&base.Props{
					ID: "demo-drawer",
					Markup: []vecty.Applyer{
						vecty.Class("demo-drawer"),
					},
				},
				&drawer.State{
					Type: drawer.Permanent,
					Content: ul.NewGroup(nil,
						&ul.GroupState{Lists: []*ul.L{
							ul.New(nil,
								&ul.State{Items: []*ul.Item{
									iconListItem("inbox", "Inbox"),
									iconListItem("star", "Star"),
									iconListItem("send", "Sent Mail"),
									iconListItem("drafts", "Drafts")},
									ClickHandler: func(l *ul.L, i *ul.Item,
										e *vecty.Event) {
										for _, v := range l.Items {
											switch {
											case i == v:
												i.Activated = true
												vecty.Rerender(i)
											default:
												v.Activated = false
												vecty.Rerender(v)
											}
										}
									},
								},
							),
							ul.ListDivider(),
							ul.New(nil,
								&ul.State{Items: []*ul.Item{
									iconListItem("email", "All Mail"),
									iconListItem("delete", "Trash"),
									iconListItem("report", "Spam"),
								}},
							),
						}},
					),
				},
			),
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
								ChangeHandler: func(e *vecty.Event) {
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
								ChangeHandler: func(e *vecty.Event) {
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
								ChangeHandler: func(e *vecty.Event) {
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
							ClickHandler: func(e *vecty.Event) {
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
							ClickHandler: func(e *vecty.Event) {
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
							ClickHandler: func(e *vecty.Event) {
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

func iconListItem(icon, text string) *ul.Item {
	var selected bool
	if icon == "inbox" {
		selected = true
	}
	return ul.NewItem(
		&base.Props{
			Markup: []vecty.Applyer{
				vecty.Class("demo-drawer-list-item"),
			},
		},
		&ul.ItemState{
			Selected: selected,
			Graphic: vecty.List{
				elem.Italic(
					vecty.Markup(
						vecty.Class("material-icons"),
						vecty.Attribute("aria-hidden", "true"),
					),
					vecty.Text(icon),
				),
			},
			Primary: text,
		},
	)
}
