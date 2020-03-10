// jsdom provides GopherJS wrappers for the jsdom Node module. jsdom is useful
// for simulating a browser environment for basic testing purposes, without
// requiring a full WebDriver setup.
//
// Currently the jsdom package most easily used by assigning objects like
// "window" to Node's global environment. This is not recommended by the jsdom
// project, and we hope to develop bindings that aid in using jsdom's features
// for safely running code in the emulated DOM without stuffing things into
// Node's global environment. See:
// https://github.com/jsdom/jsdom/wiki/Don't-stuff-jsdom-globals-onto-the-Node-global

package jsdom // import "github.com/vecty-material/material/gojs/jsdom"

import (
	"syscall/js"

	"github.com/vecty-material/material/gojs"
)

type M map[string]interface{}

type JSDOM interface {
	DOM() js.Value
	Window() js.Value
	Document() js.Value
	SetHTML(html string)
	PopulateBody(html string) js.Value
	QueryElement(querySelector string) (e js.Value, err error)
	RootElement() js.Value
}

type jsdom struct {
	js.Value
	options M
}

func New(html string, options *M) (JSDOM, error) {
	c, err := jsdomClass()
	if err != nil {
		return nil, err
	}

	j, err := newJSDOM(c, html, options)
	if err != nil {
		return nil, err
	}

	return j, err
}

func jsdomClass() (c js.Value, err error) {
	defer gojs.CatchException(&err)
	c = js.Global().Call("require", "jsdom")
	return c, err
}

func newJSDOM(c js.Value, html string, options *M) (j jsdom, err error) {
	defer gojs.CatchException(&err)
	j.options = *options
	j.Value = c.Get("JSDOM").New(html, options)
	return j, err
}

func (j jsdom) DOM() js.Value {
	return j.Value
}

func (j jsdom) Window() js.Value {
	return j.Get("window")
}

func (j jsdom) Document() js.Value {
	return j.Window().Get("document")
}

func (j jsdom) SetHTML(html string) {
	j.Document().Get("documentElement").Set("innerHTML", html)
}

// PopulateBody resets documentElement with html inside a valid html/body DOM
// and returns the HTMLElement of html.
func (j jsdom) PopulateBody(html string) js.Value {
	j.SetHTML("<html><body>" + html +
		"</body></html>")
	return js.Global().Get("document").Get("body").Get("firstElementChild")
}

func (j jsdom) QueryElement(querySelector string) (e js.Value, err error) {
	defer gojs.CatchException(&err)
	e = j.Document().Call("querySelector", querySelector)
	return e, err
}

func (j jsdom) RootElement() js.Value {
	return j.Document().Get("documentElement")
}
