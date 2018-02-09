package ripple_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/ripple"
)

func Example() {
	// Create a new instance of a material ripple component.
	c, err := ripple.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an ripple.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}

	fmt.Printf("%s\n", c)
	fmt.Printf("Unbounded: %v, Disabled: %v\n",
		c.GetObject().Get("unbounded"), c.GetObject().Get("disabled"))

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

	fmt.Printf("Unbounded: %v, Disabled: %v\n",
		c.GetObject().Get("unbounded"), c.GetObject().Get("disabled"))

	// Output:
	// {"component":"MDCRipple","status":"stopped"}
	// {"component":"MDCRipple","status":"running"}
	// Unbounded: false, Disabled: false
	// Unbounded: true, Disabled: true
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
