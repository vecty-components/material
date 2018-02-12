package checkbox_test

import (
	"fmt"
	"log"

	"agamigo.io/material/checkbox"
	"agamigo.io/material/component"
	"agamigo.io/material/mdctest"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material checkbox component.
	c := &checkbox.CB{}
	printStatus(c)

	// Set up a DOM HTMLElement suitable for a checkbox.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.ComponentType().MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := component.Start(c, rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	printStatus(c)

	printState(c)
	c.Checked = true
	c.Disabled = true
	c.Indeterminate = true
	c.Value = "new value"
	printState(c)

	// Output:
	// MDCCheckbox: uninitialized
	// MDCCheckbox: running
	//
	// Checked: false, Indeterminate: false, Disabled: false, Value: on
	//
	// Checked: true, Indeterminate: true, Disabled: true, Value: new value
}

func printStatus(c *checkbox.CB) {
	fmt.Printf("%s\n", c)
}

func printState(c *checkbox.CB) {
	fmt.Println()
	mdcObj := c.GetObject()
	fmt.Printf("Checked: %v, Indeterminate: %v, Disabled: %v, Value: %v\n",
		mdcObj.Get("checked"), mdcObj.Get("indeterminate"),
		mdcObj.Get("disabled"), mdcObj.Get("value"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
