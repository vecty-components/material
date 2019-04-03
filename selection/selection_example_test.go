package selection_test

import (
	"fmt"
	"log"

	"agamigo.io/material/internal/mdctest"
	"agamigo.io/material/selection"
	"github.com/gopherjs/gopherwasm/js"
)

func Example() {
	// Create a new instance of a material selection component.
	c := selection.New()
	printName(c)
	printState(c)
	c.SelectedIndex = 1
	c.Disabled = true
	printState(c)

	// Set up a DOM HTMLElement suitable for a selection.
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
	c.SelectedIndex = 0
	c.Disabled = false
	l := c.Component().Get("root_").Call("querySelector", ".mdc-list")
	l.Call("removeChild", c.Options().Index(1))
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}
	printState(c)

	// Output:
	// MDCSelect
	//
	// SelectedIndex: 0, SelectedString: , Disabled: false
	// SelectedElem: undefined
	// Options: undefined
	//
	// SelectedIndex: 1, SelectedString: , Disabled: true
	// SelectedElem: undefined
	// Options: undefined
	//
	// SelectedIndex: 1, SelectedString: Option #2, Disabled: true
	// SelectedElem: [object NodeList]
	// Options: [object HTMLLIElement],[object HTMLLIElement]
	//
	// SelectedIndex: 0, SelectedString: Option #1, Disabled: false
	// SelectedElem: [object NodeList]
	// Options: [object HTMLLIElement]
	//
	// SelectedIndex: 0, SelectedString: Option #1, Disabled: false
	// SelectedElem: [object NodeList]
	// Options: [object HTMLLIElement]
}

func printName(c *selection.S) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *selection.S) {
	fmt.Println()
	fmt.Printf("SelectedIndex: %v, SelectedString: %v, Disabled: %v\n",
		c.Component().Get("selectedIndex"), c.SelectedString(),
		c.Component().Get("disabled"))
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
