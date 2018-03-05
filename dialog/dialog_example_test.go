package dialog_test

import (
	"fmt"
	"log"

	"agamigo.io/material/dialog"
	"agamigo.io/material/internal/mdctest"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material dialog component.
	c := dialog.New()
	printName(c)
	printState(c)

	// Set up a DOM HTMLElement suitable for a dialog.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err.Error())
	}

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

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCDialog
	//
	// [Go] IsOpen: false
	// [JS] IsOpen: undefined
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

func printName(c *dialog.D) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *dialog.D) {
	fmt.Println()
	mdcObj := c.Component()
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
