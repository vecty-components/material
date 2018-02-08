package mdctest // import "agamigo.io/material/mdctest"

import (
	"agamigo.io/gojs"
	"agamigo.io/gojs/jsdom"
	"github.com/gopherjs/gopherjs/js"
)

const (
	MCW_MODULE = "material-components-web/dist/material-components-web"
)

var (
	Dom jsdom.JSDOM
)

func Init() error {
	err := LoadMDCModule()
	if err != nil {
		return err
	}

	Dom, err = EmulateDOM()
	if err != nil {
		return err
	}

	return nil
}

// For some reason the material-components-web node module does not come with
// MDCMenu, it only comes with an undocumented MDCSimpleMenu.
func InitMenu() (err error) {
	gojs.CatchException(&err)
	mdc := js.Global.Get("Object").New()
	mdc.Set("menu", js.Global.Call("require", "@material/menu/dist/mdc.menu"))
	js.Global.Set("mdc", mdc)

	Dom, err = EmulateDOM()
	if err != nil {
		return err
	}

	return err
}

// LoadMDCModule is a shortcut to
func LoadMDCModule() (err error) {
	gojs.CatchException(&err)
	js.Global.Set("mdc", js.Global.Call("require", MCW_MODULE))
	return err
}

// EmulateDOM sets up a fake DOM in Node for "gopherjs test" We emulate a
// browser dom since tests run in Node, and MDC components need a dom element to
// attach to.  This is not needed when running in a browser.
func EmulateDOM() (dom jsdom.JSDOM, err error) {
	dom, err = jsdom.New(``, &js.M{"pretendToBeVisual": true})
	if err != nil {
		return nil, err
	}
	js.Debugger()
	dom.SetHTML(`<html><body></body></html>`)
	js.Global.Set("window", dom.Window())
	js.Global.Set("document", dom.Window().Get("document"))
	js.Global.Set("HTMLElement", dom.Window().Get("HTMLElement"))
	raf := dom.Window().Get("requestAnimationFrame")
	js.Global.Set("requestAnimationFrame", raf)
	return dom, err
}
