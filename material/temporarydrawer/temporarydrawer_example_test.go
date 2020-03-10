package temporarydrawer_test

import (
	"fmt"
	"log"

	"syscall/js"

	"github.com/vecty-material/material/material/internal/mdctest"
	"github.com/vecty-material/material/material/temporarydrawer"
)

func Example() {
	// Create a new instance of a material temporarydrawer component.
	c := temporarydrawer.New()
	printName(c)
	printState(c)
	c.Open = true
	printState(c)

	// Set up a DOM HTMLElement suitable for a temporarydrawer.
	js.Global().Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global().Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}

	printState(c)
	c.Open = false
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	printState(c)

	// Output:
	// MDCTemporaryDrawer
	//
	// MDC Open: false
	//
	// MDC Open: true
	//
	// MDC Open: true
	//
	// MDC Open: false
	//
	// MDC Open: false
}

func printName(c *temporarydrawer.TD) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *temporarydrawer.TD) {
	fmt.Println()
	fmt.Printf("MDC Open: %v\n", c.Component().Get("open"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
