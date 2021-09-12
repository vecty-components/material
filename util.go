package material

import "syscall/js"

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
