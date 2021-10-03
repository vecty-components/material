package temporary

import (
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/lithammer/dedent"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/button"
	"github.com/vecty-material/material/demos/common"
	dcommon "github.com/vecty-material/material/demos/drawer/common"
	"github.com/vecty-material/material/drawer"
	"github.com/vecty-material/material/formfield"
	"github.com/vecty-material/material/radio"
)

type DrawerDemoView struct {
	vecty.Core
	body   *vecty.HTML
	drawer *drawer.D
}

func (c *DrawerDemoView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Drawer (Temporary) - Material Components Catalog")
	vecty.AddStylesheet("https://material-components-web.appspot.com/assets/radio.css")
	vecty.AddStylesheet("https://material-components-web.appspot.com/assets/drawer/drawer.css")

	base.AddCSS(dedent.Dedent(`
		.demo-body {
			padding: 0;
			margin: 0;
			box-sizing: border-box;
		}

		.demo-main {
			padding-left: 16px;
			padding-right: 16px;
			padding-bottom: 16px;
			overflow: auto;
		}
	`))

	c.drawer = dcommon.NewDemoDrawer(drawer.Temporary)
	c.body = elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
			vecty.Class("demo-body"),
		),
		c.drawer,
		elem.Div(
			vecty.Markup(
				vecty.Class("demo-content"),
			),
			&common.ToolbarHeader{
				Title:      "Temporary Drawer",
				Navigation: common.NavMenu,
				MenuHandler: func(e *vecty.Event) {
					c.drawer.Open = true
					vecty.Rerender(c.drawer)
				},
				NoFixed: true,
			},
			elem.Main(vecty.Markup(vecty.Class("demo-main")),
				elem.Heading1(
					vecty.Markup(vecty.Class("mdc-typography--display1")),
					vecty.Text("Temporary Drawer")),
				elem.Paragraph(
					vecty.Markup(vecty.Class("mdc-typography--body1")),
					vecty.Text("Click the menu icon above to open"),
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
						Label: elem.Anchor(
							vecty.Markup(
								event.Click(func(e *vecty.Event) {
									b := c.body.Node()
									if b.Call("getAttribute",
										"dir").String() == "rtl" {
										b.Call("setAttribute", "dir", "ltr")
										return
									}
									b.Call("setAttribute", "dir", "rtl")
								}),
							),
							vecty.Text("Toggle RTL"),
						),
						Outlined: true,
						Dense:    true,
					},
				),
			),
		),
	)
	return c.body
}
