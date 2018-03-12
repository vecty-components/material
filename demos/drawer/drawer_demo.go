package main

import (
	"path"

	"agamigo.io/vecty-material/demos/common"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type drawerDemoView struct {
	vecty.Core
}

func main() {
	vecty.RenderBody(&drawerDemoView{})
}

func (c *drawerDemoView) Render() vecty.ComponentOrHTML {
	pathname := js.Global.Get("window").Get("location").Get("pathname").String()
	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
			vecty.Class("demo-body"),
		),
		&common.ToolbarHeader{
			Title:      "Drawer",
			Navigation: common.NavBack,
		},
		elem.Main(
			vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust")),
			elem.Div(vecty.Markup(vecty.Class("intro")),
				elem.Paragraph(vecty.Text("To best show the functionality "+
					"of drawers, we put all demos in iframes. Click the links "+
					"above the iframe to view the demo in a full browser window.",
				)),
				elem.Paragraph(vecty.Text("It's worth noting that we also "+
					"use icons in this demo, which aren't loaded by default. "+
					"In order to properly include icons in your own project, "+
					"you'll need to load in the Material Icons stylesheet:",
				)),
				elem.Preformatted(vecty.Markup(
					vecty.UnsafeHTML(`&lt;link rel="stylesheet" `+
						`href="https://fonts.googleapis.com`+
						`/icon?family=Material+Icons"&gt;`,
					))),
			),
			elem.Section(vecty.Markup(vecty.Class("examples")),
				elem.Div(vecty.Markup(vecty.Class("example")),
					elem.Heading2(
						vecty.Markup(vecty.Class("drawer-example-heading")),
						elem.Span(vecty.Text("Temporary Drawer")),
					),
					elem.Paragraph(elem.Anchor(
						vecty.Markup(
							prop.Href(path.Clean(pathname+"/temporary/")),
							vecty.Property("target", "_blank"),
						),
						vecty.Text("View in separate window"),
					)),
					elem.InlineFrame(vecty.Markup(
						prop.Src(path.Clean(pathname+"/temporary/")))),
				),
				elem.Div(vecty.Markup(vecty.Class("example")),
					elem.Heading2(
						vecty.Markup(vecty.Class("drawer-example-heading")),
						elem.Span(vecty.Text("Persistent Drawer")),
					),
					elem.Paragraph(elem.Anchor(
						vecty.Markup(
							prop.Href(path.Clean(pathname+"/persistent/")),
							vecty.Property("target", "_blank"),
						),
						vecty.Text("View in separate window"),
					)),
					elem.InlineFrame(vecty.Markup(
						prop.Src(path.Clean(pathname+"/persistent/")))),
				),
				elem.Div(vecty.Markup(vecty.Class("example")),
					elem.Heading2(
						vecty.Markup(vecty.Class("drawer-example-heading")),
						elem.Span(vecty.Text("Permanent drawer above toolbar")),
					),
					elem.Paragraph(elem.Anchor(
						vecty.Markup(
							prop.Href(path.Clean(pathname+"/permanent-above/")),
							vecty.Property("target", "_blank"),
						),
						vecty.Text("View in separate window"),
					)),
					elem.InlineFrame(vecty.Markup(
						prop.Src(path.Clean(pathname+"/permanent-above/")))),
				),
				elem.Div(vecty.Markup(vecty.Class("example")),
					elem.Heading2(
						vecty.Markup(vecty.Class("drawer-example-heading")),
						elem.Span(vecty.Text("Permanent drawer below toolbar")),
					),
					elem.Paragraph(elem.Anchor(
						vecty.Markup(
							prop.Href(path.Clean(pathname+"/permanent-below/")),
							vecty.Property("target", "_blank"),
						),
						vecty.Text("View in separate window"),
					)),
					elem.InlineFrame(vecty.Markup(
						prop.Src(path.Clean(pathname+"/permanent-below/")))),
				),
			),
		),
	)
}
