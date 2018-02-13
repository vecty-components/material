package ripple_test

import (
	"fmt"
	"log"

	"agamigo.io/material"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/ripple"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material ripple component.
	c := &ripple.R{}
	printStatus(c)

	// Set up a DOM HTMLElement suitable for a ripple.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.ComponentType().MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := material.Start(c, rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	printStatus(c)

	printState(c)
	c.Unbounded = true
	err = c.Activate()
	if err != nil {
		fmt.Printf("Unable to active ripple: %v", err)
	}
	err = c.Deactivate()
	if err != nil {
		fmt.Printf("Unable to deactive ripple: %v", err)
	}
	err = c.Layout()
	if err != nil {
		fmt.Printf("Unable to recompute ripple layout: %v", err)
	}
	c.Disabled = true
	printState(c)

	// Output:
	// MDCRipple: uninitialized
	// MDCRipple: running
	//
	// Unbounded: false, Disabled: false
	//
	// Unbounded: true, Disabled: true
}

func printStatus(c *ripple.R) {
	fmt.Printf("%s\n", c)
}

func printState(c *ripple.R) {
	fmt.Println()
	fmt.Printf("Unbounded: %v, Disabled: %v\n",
		c.GetObject().Get("unbounded"), c.GetObject().Get("disabled"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
