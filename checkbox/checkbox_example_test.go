package checkbox_test

import (
	"fmt"
	"log"

	"agamigo.io/material/checkbox"
	"agamigo.io/material/internal/mdctest"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material checkbox component.
	c := checkbox.New()
	printName(c)
	printState(c)
	c.Checked = true
	c.Disabled = true
	c.Indeterminate = true
	c.Value = "new value"
	printState(c)

	// Set up a DOM HTMLElement suitable for a checkbox.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err.Error())
	}

	c.Value = "after start"
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	printState(c)

	// Output:
	// MDCCheckbox
	// Checked: false, Indeterminate: false, Disabled: false, Value: v
	// Checked: true, Indeterminate: true, Disabled: true, Value: v
	// Checked: true, Indeterminate: true, Disabled: true, Value: v
	// Checked: true, Indeterminate: true, Disabled: true, Value: v
}

func printName(c *checkbox.CB) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *checkbox.CB) {
	mdcObj := c.Component()
	fmt.Printf("Checked: %v, Indeterminate: %v, Disabled: %v, Value: v",
		mdcObj.Get("checked"), mdcObj.Get("indeterminate"),
		mdcObj.Get("disabled")) //, mdcObj.Get("value"))
	fmt.Println()
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
