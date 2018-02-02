package component_test

import (
	"fmt"
	"log"

	"agamigo.io/gojs"
	"agamigo.io/gojs/jsdom"
	"agamigo.io/material/component"
	"github.com/gopherjs/gopherjs/js"
)

const (
	MCW_NODE_MODULE = "material-components-web/dist/material-components-web"
)

func Example() {
	// We emulate a browser dom since this example/test is run in Node, and
	// c.Start() needs a dom element to attach to.  This is not needed when
	// running in a browser.
	err := emulateDOM()
	if err != nil {
		log.Fatalf("%v", err)
	}
	// Bring in the material-components-web JS module.
	// Note: This could be done using an external JS loader, inc.js file etc..
	err = loadMDCModule()
	if err != nil {
		log.Fatalf("Unable to load MDC JS module: %v", err)
	}

	// Create a new instance of a checkbox component.
	c, err := component.New(component.Checkbox)
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Start the component after its HTMLElement is instantiated.
	// In this case it was instantiated in emulateDOM()
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Output:
	// {"component":"MDCCheckbox","status":"stopped"}
	// {"component":"MDCCheckbox","status":"running"}
}

func loadMDCModule() (err error) {
	gojs.CatchException(&err)
	js.Global.Set("mdc", js.Global.Call("require", MCW_NODE_MODULE))
	return err
}

func emulateDOM() (err error) {
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
