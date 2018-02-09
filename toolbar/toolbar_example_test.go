package toolbar_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/toolbar"
)

func Example() {
	// Create a new instance of a material toolbar component.
	c, err := toolbar.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an toolbar.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}

	fmt.Printf("%s\n", c)

	// Output:
	// {"component":"MDCToolbar","status":"stopped"}
	// {"component":"MDCToolbar","status":"running"}
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
