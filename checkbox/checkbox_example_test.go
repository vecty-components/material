package checkbox_test

import (
	"fmt"
	"log"

	"agamigo.io/material/checkbox"
	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
)

func Example() {
	// Create a new instance of a material checkbox component.
	c, err := checkbox.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for a checkbox.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)
	fmt.Printf("Checked: %v, Indeterminate: %v, Disabled: %v, Value: %v\n",
		c.Checked, c.Indeterminate, c.Disabled, c.Value)

	c.Checked = true
	fmt.Printf("Checked: %v, Indeterminate: %v, Disabled: %v, Value: %v\n",
		c.Checked, c.Indeterminate, c.Disabled, c.Value)

	c.Disabled = true
	fmt.Printf("Checked: %v, Indeterminate: %v, Disabled: %v, Value: %v\n",
		c.Checked, c.Indeterminate, c.Disabled, c.Value)

	c.Indeterminate = true
	c.Value = "new value"
	fmt.Printf("Checked: %v, Indeterminate: %v, Disabled: %v, Value: %v\n",
		c.Checked, c.Indeterminate, c.Disabled, c.Value)

	// Output:
	// {"component":"MDCCheckbox","status":"stopped"}
	// {"component":"MDCCheckbox","status":"running"}
	// Checked: false, Indeterminate: false, Disabled: false, Value: on
	// Checked: true, Indeterminate: false, Disabled: false, Value: on
	// Checked: true, Indeterminate: false, Disabled: true, Value: on
	// Checked: true, Indeterminate: true, Disabled: true, Value: new value
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
