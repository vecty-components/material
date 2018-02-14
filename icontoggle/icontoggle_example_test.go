package icontoggle_test

import (
	"fmt"
	"log"

	"agamigo.io/material/icontoggle"
	"agamigo.io/material/internal/mdctest"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material icontoggle component.
	c := &icontoggle.IT{}

	// Set up a DOM HTMLElement suitable for an icontoggle.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err.Error())
	}

	printStatus(c)
	printState(c)
	c.Disabled = true
	c.On = true
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCIconToggle
	//
	// On: false, Disabled: false
	//
	// On: true, Disabled: true
}

func printStatus(c *icontoggle.IT) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *icontoggle.IT) {
	fmt.Println()
	mdcObj := c.Component()
	fmt.Printf("On: %v, Disabled: %v\n",
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
