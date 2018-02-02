package mdctest

import (
	"agamigo.io/gojs"
	"agamigo.io/gojs/jsdom"
	"github.com/gopherjs/gopherjs/js"
)

const (
	MCW_NODE_MODULE = "material-components-web/dist/material-components-web"
)

func LoadMDCModule() (err error) {
	gojs.CatchException(&err)
	js.Global.Set("mdc", js.Global.Call("require", MCW_NODE_MODULE))
	return err
}

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
