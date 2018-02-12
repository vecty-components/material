package temporarydrawer_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component"
	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/temporarydrawer"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material temporarydrawer component.
	c := &temporarydrawer.TD{}
	printStatus(c)

	// Set up a DOM HTMLElement suitable for a temporarydrawer.
	js.Global.Get("document").Get("body").Set("innerHTML",
		componenthtml.HTML(c.MDCType()))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := component.Start(c, rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	printStatus(c)

	printState(c)
	c.Open = true
	printState(c)

	// Output:
	// MDCTemporaryDrawer: uninitialized
	// MDCTemporaryDrawer: running
	//
	// MDC Open: false
	//
	// MDC Open: true
}

func printStatus(c *temporarydrawer.TD) {
	fmt.Printf("%s\n", c)
}

func printState(c *temporarydrawer.TD) {
	fmt.Println()
	fmt.Printf("MDC Open: %v\n", c.GetObject().Get("open"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
