package toolbar_test

import (
	"fmt"
	"log"

	"github.com/vecty-material/material/internal/mdctest"
	"github.com/vecty-material/material/toolbar"
	"github.com/gopherjs/gopherwasm/js"
)

func Example() {
	// Create a new instance of a material toolbar component.
	c := toolbar.New()
	printName(c)

	// Set up a DOM HTMLElement suitable for a toolbar.
	js.Global().Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global().Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}
	printName(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	printName(c)

	// Output:
	// MDCToolbar
	// MDCToolbar
	// MDCToolbar
}

func printName(c *toolbar.T) {
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
