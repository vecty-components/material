package base

import (
	"syscall/js"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	router "marwan.io/vecty-router"
)

type StaticComponent struct {
	vecty.Core
	Child vecty.ComponentOrHTML
}

// RenderStoredChild is a helper which provides a Component which wraps the
// provided ComponentOrHTML. It exists as a workaround to a vecty issue.
//
// See: https://github.com/hexops/vecty/issues/191
func RenderStoredChild(child vecty.ComponentOrHTML) *StaticComponent {
	return &StaticComponent{Child: child}
}

func (c *StaticComponent) Render() vecty.ComponentOrHTML {
	switch t := c.Child.(type) {
	case vecty.List:
		return elem.Div(t)
	}
	return c.Child
}

func (c *StaticComponent) SkipRender(prev vecty.Component) bool {
	switch prev.(type) {
	case *StaticComponent:
		return true
	}
	return false
}

func MarkupIfNotNil(rootMarkup *vecty.MarkupList) vecty.Applyer {
	if rootMarkup != nil {
		return vecty.MarkupIf(rootMarkup != nil, *rootMarkup)
	} else {
		return nil
	}
}

func SetViewport() {
	meta := js.Global().Get("document").Call("createElement", "meta")
	meta.Set("name", "viewport")
	meta.Set("content", "width=device-width, initial-scale=1")
	js.Global().Get("document").Get("head").Call("appendChild", meta)
}

func AddIcon(url string) {
	link := js.Global().Get("document").Call("createElement", "link")
	link.Set("rel", "icon")
	link.Set("type", "image/png")
	link.Set("href", url)
	js.Global().Get("document").Get("head").Call("appendChild", link)
}

func AddScript(url string) {
	script := js.Global().Get("document").Call("createElement", "script")
	script.Set("src", url)
	js.Global().Get("document").Get("head").Call("appendChild", script)
}

func AddCSS(css string) {
	style := js.Global().Get("document").Call("createElement", "style")
	style.Set("innerHTML", css)
	js.Global().Get("document").Get("head").Call("appendChild", style)
}

func ClearCSS() {
	styles := js.Global().Get("document").Call("getElementsByTagName", "style")
	for i := 0; i < styles.Length(); i++ {
		styles.Index(i).Call("remove")
	}
}

func AddStyles() {
	vecty.AddStylesheet("https://unpkg.com/material-components-web@" + MDC_VERSION + "/dist/material-components-web.min.css")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto+Mono")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:300,400,500")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
}

func Boot() {
	AddScript("https://unpkg.com/material-components-web@" + MDC_VERSION + "/dist/material-components-web.min.js")

	for {
		time.Sleep(25 * time.Millisecond)
		if mdcObject := js.Global().Get("mdc"); !mdcObject.IsUndefined() {
			break
		}
	}
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
