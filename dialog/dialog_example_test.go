package dialog_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/dialog"
	"agamigo.io/material/mdctest"
)

func Example() {
	// Create a new instance of a material dialog component.
	c, err := dialog.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up the DOM with an HTMLElement suitable for a dialog.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement in the DOM.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)
	fmt.Printf("IsOpen: %v\n", c.IsOpen())

	err = c.Open()
	if err != nil {
		log.Fatalf("Unable to open dialog: %v", err)
	}
	fmt.Printf("IsOpen: %v\n", c.IsOpen())

	err = c.Close()
	if err != nil {
		log.Fatalf("Unable to close dialog: %v", err)
	}
	fmt.Printf("IsOpen: %v\n", c.IsOpen())

	// Output:
	// {"component":"MDCDialog","status":"stopped"}
	// {"component":"MDCDialog","status":"running"}
	// IsOpen: false
	// IsOpen: true
	// IsOpen: false
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
