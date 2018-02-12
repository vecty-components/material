package dialog_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component"
	"agamigo.io/material/dialog"
	"agamigo.io/material/mdctest"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material dialog component.
	c := &dialog.D{}
	printStatus(c)

	// Set up a DOM HTMLElement suitable for a dialog.
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
	err = c.Open()
	if err != nil {
		log.Fatalf("Unable to open dialog: %v", err)
	}
	printState(c)
	err = c.Close()
	if err != nil {
		log.Fatalf("Unable to close dialog: %v", err)
	}
	printState(c)

	// Output:
	// MDCDialog: uninitialized
	// MDCDialog: running
	//
	// [Go] IsOpen: false
	// [JS] IsOpen: false
	//
	// [Go] IsOpen: true
	// [JS] IsOpen: true
	//
	// [Go] IsOpen: false
	// [JS] IsOpen: false
}

func printStatus(c *dialog.D) {
	fmt.Printf("%s\n", c)
}

func printState(c *dialog.D) {
	fmt.Println()
	mdcObj := c.GetObject()
	fmt.Printf("[Go] IsOpen: %v\n", c.IsOpen)
	fmt.Printf("[JS] IsOpen: %v\n", mdcObj.Get("open"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
