package dialog_test

import (
	"fmt"
	"log"

	"syscall/js"

	"github.com/vecty-material/material/material/dialog"
	"github.com/vecty-material/material/material/internal/mdctest"
)

func Example() {
	// Create a new instance of a material dialog component.
	c := dialog.New()
	printName(c)
	printState(c)
	c.Open = true
	printState(c)

	// Set up a DOM HTMLElement suitable for a dialog.
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
	c.Open = false
	printState(c)
	c.Open = true
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	printState(c)

	// Output:
	// MDCDialog
	//
	// [Go] Open: false
	// [JS] Open: false
	//
	// [Go] Open: true
	// [JS] Open: true
	//
	// [Go] Open: true
	// [JS] Open: true
	//
	// [Go] Open: false
	// [JS] Open: false
	//
	// [Go] Open: true
	// [JS] Open: true
	//
	// [Go] Open: true
	// [JS] Open: true
}

func printName(c *dialog.D) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *dialog.D) {
	fmt.Println()
	mdcObj := c.Component()
	fmt.Printf("[Go] Open: %v\n", c.Open)
	fmt.Printf("[JS] Open: %v\n", mdcObj.Get("open"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
