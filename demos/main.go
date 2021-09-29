package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/vecty-material/material/base"

	router "marwan.io/vecty-router"
)

func main() {
	base.SetViewport()
	base.AddResources()

	vecty.AddStylesheet("https://material-components-web.appspot.com/assets/index.css")

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
			router.NewRoute("/", &CatalogView{}, router.NewRouteOpts{ExactMatch: true}),
			router.NotFoundHandler(&NotFoundView{}),
		),
	)
}
