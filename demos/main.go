package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	router "marwan.io/vecty-router"

	"github.com/vecty-material/material/base"
)

func main() {
	base.SetViewport()
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
	vecty.AddStylesheet("/assets/styles/App.css")

	return elem.Body(&CatalogPage{})
}

func RichLink(route string, elements []vecty.ComponentOrHTML, opts router.LinkOptions) *vecty.HTML {
	children := make([]vecty.MarkupOrChild, len(elements))
	for i, element := range elements {
		children[i] = element
	}

	return elem.Anchor(append([]vecty.MarkupOrChild{
		vecty.Markup(
			prop.Href(route),
			vecty.MarkupIf(opts.ID != "", prop.ID(opts.ID)),
			vecty.MarkupIf(opts.Class != "", vecty.Class(opts.Class)),
			event.Click(onClick(route)).PreventDefault(),
		),
	}, children...)...)
}

func onClick(route string) router.EventCallback {
	return func(e *vecty.Event) {
		router.Redirect(route)
	}
}
