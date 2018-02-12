package toolbar_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/toolbar"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material toolbar component.
	c := &toolbar.T{}
	printStatus(c)

	// Set up a DOM HTMLElement suitable for a toolbar.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.ComponentType().MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := component.Start(c, rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	printStatus(c)

	// Output:
	// MDCTextField: uninitialized
	// MDCTextField: running
}

func printStatus(c *toolbar.T) {
	fmt.Printf("%s\n", c)
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
