package radio_test

import (
	"fmt"
	"log"

	"syscall/js"

	"github.com/vecty-material/material/internal/mdctest"
	"github.com/vecty-material/material/radio"
)

func Example() {
	// Create a new instance of a material radio component.
	c := radio.New()
	printName(c)
	printState(c)
	c.Checked = true
	c.Disabled = true
	c.Value = "before Start()"
	printState(c)

	// Set up a DOM HTMLElement suitable for a radio.
	js.Global().Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global().Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}

	c.Value = "after Start() 1"
	printState(c)
	c.Checked = false
	c.Disabled = false
	c.Value = "after Start() 2"

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	c.Value = "after Stop()"
	printState(c)

	// Output:
	// MDCRadio
	// [Go] Checked: false, Disabled: false, Value:
	// [JS] Checked: false, Disabled: false, Value:
	// [Go] Checked: true, Disabled: true, Value:before Start()
	// [JS] Checked: true, Disabled: true, Value:before Start()
	// [Go] Checked: true, Disabled: true, Value:after Start() 1
	// [JS] Checked: true, Disabled: true, Value:after Start() 1
	// [Go] Checked: false, Disabled: false, Value:after Stop()
	// [JS] Checked: false, Disabled: false, Value:after Stop()
}

func printName(c *radio.R) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *radio.R) {
	fmt.Printf("[Go] Checked: %v, Disabled: %v, Value:%v\n",
		c.Checked, c.Disabled, c.Value)
	fmt.Printf("[JS] Checked: %v, Disabled: %v, Value:%v\n",
		c.Component().Get("checked"), c.Component().Get("disabled"),
		c.Component().Get("value"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
