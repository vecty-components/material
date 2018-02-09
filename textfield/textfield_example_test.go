package textfield_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/textfield"
)

func Example() {
	// Create a new instance of a material textfield component.
	c, err := textfield.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an textfield.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}

	fmt.Printf("%s\n", c)

	fmt.Printf("Disabled: %v, Valid: %v, Required: %v\n",
		c.Disabled, c.Valid, c.Required)
	fmt.Printf("Value: %v, HelperText: %v\n", c.Value, c.HelperText)

	c.Required = false
	c.HelperText = "Must be at least 8 characters."
	c.Value = "longerpassword"
	c.Disabled = true

	fmt.Printf("Disabled: %v, Valid: %v, Required: %v\n",
		c.Disabled, c.Valid, c.Required)
	fmt.Printf("Value: %v, HelperText: %v\n", c.Value, c.HelperText)

	// Output:
	// {"component":"MDCTextField","status":"stopped"}
	// {"component":"MDCTextField","status":"running"}
	// Disabled: false, Valid: false, Required: true
	// Value: , HelperText: undefined
	// Disabled: true, Valid: true, Required: false
	// Value: longerpassword, HelperText: Must be at least 8 characters.
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
