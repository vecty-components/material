package slider_test

import (
	"fmt"
	"log"

	"syscall/js"

	"github.com/vecty-material/material/internal/mdctest"
	"github.com/vecty-material/material/slider"
)

func Example() {
	// Create a new instance of a material slider component.
	c := slider.New()
	printName(c)
	printState(c)
	c.Value = 10.0
	c.Min = 5.5
	c.Max = 50.0
	c.Step = 5.0
	c.Disabled = true
	printState(c)

	// Set up a DOM HTMLElement suitable for a slider.
	js.Global().Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global().Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}

	printState(c)
	c.Value = c.Value + 5
	c.Min = c.Min + 5
	c.Max = c.Max + 5
	c.Step = c.Step + 5
	c.Disabled = false
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	printState(c)

	// Output:
	// MDCSlider
	//
	// [Go] Value: 0, Min: 0, Max 0, Step 0, Disabled: false
	//
	// [Go] Value: 10, Min: 5.5, Max 50, Step 5, Disabled: true
	//
	// [Go] Value: 10, Min: 5.5, Max 50, Step 5, Disabled: true
	// [JS] Value: 10, Min: 5.5, Max 50, Step 5, Disabled: true
	//
	// [Go] Value: 20, Min: 10.5, Max 55, Step 10, Disabled: false
	// [JS] Value: 20, Min: 10.5, Max 55, Step 10, Disabled: false
	//
	// [Go] Value: 20, Min: 10.5, Max 55, Step 10, Disabled: false
	// [JS] Value: 20, Min: 10.5, Max 55, Step 10, Disabled: false
}

func printName(c *slider.S) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *slider.S) {
	fmt.Println()
	fmt.Printf("[Go] Value: %v, Min: %v, Max %v, Step %v, Disabled: %v\n",
		c.Value, c.Min, c.Max, c.Step, c.Disabled)
	if !c.Component().Get("foundation_").IsUndefined() {
		fmt.Printf("[JS] Value: %v, Min: %v, Max %v, Step %v, Disabled: %v\n",
			c.Component().Get("foundation_").Get("value_"),
			c.Component().Get("foundation_").Get("min_"),
			c.Component().Get("foundation_").Get("max_"),
			c.Component().Get("foundation_").Get("step_"),
			c.Component().Get("foundation_").Get("disabled_"))
	}
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
