package slider_test

import (
	"fmt"
	"log"

	"agamigo.io/material/internal/mdctest"
	"agamigo.io/material/slider"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material slider component.
	c := &slider.S{}

	// Set up a DOM HTMLElement suitable for a slider.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}

	printStatus(c)
	printState(c)
	c.Value = 10.0
	c.Min = 5.5
	c.Max = 50.0
	c.Step = 5.0
	c.Disabled = true
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCSlider
	//
	// Value: 0, Min: 0, Max 100, Step 0, Disabled: false
	//
	// Value: 10, Min: 5.5, Max 50, Step 5, Disabled: true
}

func printStatus(c *slider.S) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *slider.S) {
	fmt.Println()
	fmt.Printf("Value: %v, Min: %v, Max %v, Step %v, Disabled: %v\n",
		c.Component().Get("foundation_").Get("value_"),
		c.Component().Get("foundation_").Get("min_"),
		c.Component().Get("foundation_").Get("max_"),
		c.Component().Get("foundation_").Get("step_"),
		c.Component().Get("foundation_").Get("disabled_"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
