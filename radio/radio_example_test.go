package radio_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/radio"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material radio component.
	c := &radio.R{}
	printStatus(c)

	// Set up a DOM HTMLElement suitable for a radio.
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
	c.Checked = false
	c.Disabled = true
	c.Value = "new value"
	printState(c)

	// Output:
	// MDCRadio: uninitialized
	// MDCRadio: running
	//
	// Checked: true, Disabled: false, Value: on
	//
	// Checked: false, Disabled: true, Value: new value
}

func printStatus(c *radio.R) {
	fmt.Printf("%s\n", c)
}

func printState(c *radio.R) {
	fmt.Println()
	fmt.Printf("Checked: %v, Disabled: %v, Value: %v\n",
		c.GetObject().Get("checked"), c.GetObject().Get("disabled"),
		c.GetObject().Get("value"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
