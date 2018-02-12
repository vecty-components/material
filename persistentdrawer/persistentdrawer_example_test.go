package persistentdrawer_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/persistentdrawer"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material persistentdrawer component.
	c := &persistentdrawer.PD{}
	printStatus(c)

	// Set up a DOM HTMLElement suitable for a persistentdrawer.
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
	c.Open = true
	printState(c)

	// Output:
	// MDCPersistentDrawer: uninitialized
	// MDCPersistentDrawer: running
	//
	// MDC Open: false
	//
	// MDC Open: true
}

func printStatus(c *persistentdrawer.PD) {
	fmt.Printf("%s\n", c)
}

func printState(c *persistentdrawer.PD) {
	fmt.Println()
	fmt.Printf("MDC Open: %v\n", c.GetObject().Get("open"))
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
