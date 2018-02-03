package mdctest

//go:generate yarn install

import (
	"log"

	"agamigo.io/gojs"
	"agamigo.io/gojs/jsdom"
	"github.com/gopherjs/gopherjs/js"
)

func init() {
	err := EmulateDOM()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

const (
	MCW_NODE_MODULE = "material-components-web/dist/material-components-web"
)

// LoadMDCModule is a shortcut to
func LoadMDCModule() (err error) {
	gojs.CatchException(&err)
	js.Global.Set("mdc", js.Global.Call("require", MCW_NODE_MODULE))
	return err
}

// EmulateDOM sets up a fake DOM in Node for "gopherjs test" We emulate a
// browser dom since tests run in Node, and MDC components need a dom element to
// attach to.  This is not needed when running in a browser.
func EmulateDOM() (err error) {
	html := `<div class="mdc-checkbox">
				<input class="mdc-checkbox__native-control" type="checkbox">
			</div>`
	dom, err := jsdom.New()
	if err != nil {
		return err
	}
	dom.SetHTML(`<html><body>` + html + `</body></html>`)
	js.Global.Set("window", dom.Window())
	js.Global.Set("HTMLElement", dom.Window().Get("HTMLElement"))
	return nil
}
