package toolbar_test

import (
	"fmt"
	"log"

	"agamigo.io/material/mdctest"
	"agamigo.io/material/toolbar"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material toolbar component.
	c := &toolbar.T{}

	// Set up a DOM HTMLElement suitable for a toolbar.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}

	printStatus(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCTextField
}

func printStatus(c *toolbar.T) {
	fmt.Printf("%s\n", c.Component().Type)
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
