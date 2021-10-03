package iconbutton_test

import (
	"fmt"
	"log"

	"syscall/js"

	"github.com/vecty-material/material/components/iconbutton"
	"github.com/vecty-material/material/components/internal/mdctest"
)

func Example() {
	// Create a new instance of a material iconbutton component.
	c := iconbutton.New()
	printName(c)
	printState(c)
	c.On = true

	// Set up a DOM HTMLElement suitable for an iconbutton.
	js.Global().Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global().Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err.Error())
	}

	printState(c)
	c.On = false
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCIconbutton
	//
	// [Go] On: false, Disabled: false
	// [JS] On: false, Disabled: false
	//
	// [Go] On: true, Disabled: true
	// [JS] On: true, Disabled: true
	//
	// [Go] On: false, Disabled: false
	// [JS] On: false, Disabled: false
}

func printName(c *iconbutton.IB) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *iconbutton.IB) {
	fmt.Println()
	mdcObj := c.Component()
	fmt.Printf("[Go] On: %v, Disabled: %v\n",
		c.On, true)
	fmt.Printf("[JS] On: %v, Disabled: %v\n",
		mdcObj.Get("on"), mdcObj.Get("disabled"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
