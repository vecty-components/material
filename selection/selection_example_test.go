package selection_test

import (
	"fmt"
	"log"

	"agamigo.io/material"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/selection"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material selection component.
	c := &selection.S{}
	printStatus(c)

	// Set up a DOM HTMLElement suitable for a selection.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.ComponentType().MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := material.Start(c, rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	printStatus(c)

	printState(c)
	c.SelectedIndex = 0
	c.Disabled = true
	l := c.GetObject().Get("root_").Call("querySelector", ".mdc-list")
	l.Call("removeChild", c.Options().Index(1))
	printState(c)

	// Output:
	// MDCSelect: uninitialized
	// MDCSelect: running
	//
	// SelectedIndex: -1, SelectedString: , Disabled: false
	// SelectedElem: [object NodeList]
	// Options: [object HTMLLIElement],[object HTMLLIElement]
	//
	// SelectedIndex: 0, SelectedString: Option #1, Disabled: true
	// SelectedElem: [object NodeList]
	// Options: [object HTMLLIElement]
}

func printStatus(c *selection.S) {
	fmt.Printf("%s\n", c)
}

func printState(c *selection.S) {
	fmt.Println()
	fmt.Printf("SelectedIndex: %v, SelectedString: %v, Disabled: %v\n",
		c.GetObject().Get("selectedIndex"), c.SelectedString(),
		c.GetObject().Get("disabled"))
	fmt.Printf("SelectedElem: %v\nOptions: %v\n",
		c.SelectedElem(), c.Options())
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
