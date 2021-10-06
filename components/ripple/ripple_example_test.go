package ripple_test

import (
	"fmt"
	"log"

	"syscall/js"

	"github.com/vecty-components/material/components/internal/mdctest"
	"github.com/vecty-components/material/components/ripple"
)

func Example() {
	// Create a new instance of a material ripple component.
	c := ripple.New()
	printName(c)
	printState(c)
	c.Unbounded = true
	c.Disabled = true
	printState(c)

	// Set up a DOM HTMLElement suitable for a ripple.
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
	c.Unbounded = false
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
	c.Disabled = false
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	printState(c)

	// Output:
	// MDCRipple
	//
	// Unbounded: false, Disabled: false
	//
	// Unbounded: true, Disabled: true
	//
	// Unbounded: true, Disabled: true
	//
	// Unbounded: false, Disabled: false
	//
	// Unbounded: false, Disabled: false
}

func printName(c *ripple.R) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *ripple.R) {
	fmt.Println()
	fmt.Printf("Unbounded: %v, Disabled: %v\n",
		c.Component().Get("unbounded"), c.Component().Get("disabled"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
