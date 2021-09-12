package formfield_test

import (
	"fmt"
	"log"

	"syscall/js"

	"github.com/vecty-material/material/components/checkbox"
	"github.com/vecty-material/material/components/formfield"
	"github.com/vecty-material/material/components/internal/mdctest"
)

func Example() {
	// Create a new instance of a material formfield component and its
	// childElement.
	child := checkbox.New()
	c := formfield.New()

	// Set up a DOM HTMLElement suitable for a formfield.
	js.Global().Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global().Get("document").Get("body").Get("firstElementChild")
	childElem := rootElem.Get("firstElementChild")

	// Start the child element.
	err := child.Start(childElem)
	if err != nil {
		log.Fatalf("Unable to start child component %s: %v\n",
			child.Component().Type, err.Error())
	}

	// Start the parent formfield component, which associates it with an
	// HTMLElement.
	err = c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err.Error())
	}

	printName(c)
	printState(c)
	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	c.Input = child.Component().Value
	err = c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err.Error())
	}
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCFormField
	//
	// Child Element: [object Object]
	//
	// Child Element: [object Object]
}

func printName(c *formfield.FF) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *formfield.FF) {
	fmt.Println()
	mdcObj := c.Component()
	fmt.Printf("Child Element: %v\n",
		mdcObj.Get("input"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
