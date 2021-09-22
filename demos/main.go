package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"

	"github.com/vecty-material/material/base"
	dbutton "github.com/vecty-material/material/demos/button"
	dcheckbox "github.com/vecty-material/material/demos/checkbox"
	ddialog "github.com/vecty-material/material/demos/dialog"
	ddrawer "github.com/vecty-material/material/demos/drawer"
	dpadrawer "github.com/vecty-material/material/demos/drawer/permanent-above"
	dpbdrawer "github.com/vecty-material/material/demos/drawer/permanent-below"
	dpersistent "github.com/vecty-material/material/demos/drawer/persistent"
	dtemporary "github.com/vecty-material/material/demos/drawer/temporary"
	dicontoggle "github.com/vecty-material/material/demos/icontoggle"
	dlist "github.com/vecty-material/material/demos/list"
	dmenu "github.com/vecty-material/material/demos/menu"
	dradio "github.com/vecty-material/material/demos/radio"

	router "marwan.io/vecty-router"
)

func main() {
	base.SetViewport()

	base.AddIcon("https://material-components-web.appspot.com/images/logo_components_color_2x_web_48dp.png")
	base.AddResources()

	body := &Body{}
	vecty.RenderBody(body)
}

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			router.NewRoute("/", &DemosCatalogView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/button", &dbutton.ButtonDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/checkbox", dcheckbox.NewCheckboxDemoView(), router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/dialog", ddialog.NewDialogDemoView(), router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/drawer", &ddrawer.DrawerDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/drawer/permanent-above", &dpadrawer.DrawerDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/drawer/permanent-below", &dpbdrawer.DrawerDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/drawer/persistent", &dpersistent.DrawerDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/drawer/temporary", &dtemporary.DrawerDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/icontoggle", &dicontoggle.IcontoggleDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/radio", &dradio.RadioDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/list", &dlist.ListDemoView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NewRoute("/menu", dmenu.NewMenuDemoView(), router.NewRouteOpts{ExactMatch: true}),

			router.NotFoundHandler(&notFound{}),
		),
	)
}

type notFound struct {
	vecty.Core
}

func (nf *notFound) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("home-view")),
		elem.Div(
			vecty.Markup(prop.ID("home-top")),
			elem.Heading1(
				vecty.Text("page not found"),
			),
		),
	)
}
