package radio_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/radio"
)

func Example() {
	// Create a new instance of a material radio component.
	c, err := radio.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an radio.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}

	fmt.Printf("%s\n", c)
	fmt.Printf("Checked: %v, Disabled: %v, Value: %v\n",
		c.GetObject().Get("checked"), c.GetObject().Get("disabled"),
		c.GetObject().Get("value"))

	c.Checked = false
	c.Disabled = true
	c.Value = "new value"

	fmt.Printf("Checked: %v, Disabled: %v, Value: %v\n",
		c.GetObject().Get("checked"), c.GetObject().Get("disabled"),
		c.GetObject().Get("value"))

	// Output:
	// {"component":"MDCRadio","status":"stopped"}
	// {"component":"MDCRadio","status":"running"}
	// Checked: true, Disabled: false, Value: on
	// Checked: false, Disabled: true, Value: new value
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
