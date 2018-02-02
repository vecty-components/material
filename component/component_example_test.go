package component_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component"
	"agamigo.io/material/mdctest"
)

func Example() {
	// We emulate a browser dom since this example/test is run in Node, and
	// c.Start() needs a dom element to attach to.  This is not needed when
	// running in a browser.
	err := mdctest.EmulateDOM()
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Bring in the material-components-web JS module.
	// Note: This could be done using an external JS loader, inc.js file etc..
	err = mdctest.LoadMDCModule()
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
