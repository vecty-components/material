package snackbar_test

import (
	"fmt"
	"log"

	"agamigo.io/material/mdctest"
	"agamigo.io/material/snackbar"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material snackbar component.
	c := &snackbar.S{}

	// Set up a DOM HTMLElement suitable for a snackbar.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}

	printStatus(c)
	printState(c)
	c.Message = "snackbar message here"
	c.ActionHandler = func() { fmt.Println("Action Handled!") }
	c.ActionText = "Action Button Text"
	c.Timeout = 1000 // 1 second
	c.MultiLine = true
	c.ActionOnBottom = true
	c.DismissOnAction = true
	err = c.Show()
	if err != nil {
		log.Fatalf("Unable to show snackbar: %v", err)
	}
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCSnackbar

	// DismissOnAction: undefined
	// Snackbar has not been shown yet.

	// DismissOnAction: true
	// Message: snackbar message here
	// Timeout: 1000
	// ActionHandler: 0x1, ActionText: Action Button Text
	// Multiline: true, ActionOnBottom: true
}

func printStatus(c *snackbar.S) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *snackbar.S) {
	fmt.Println()
	fmt.Printf("DismissOnAction: %v\n",
		c.Component().Get("dismissOnAction"))

	o := c.Component().Get("foundation_").Get("snackbarData_")
	if o == nil {
		fmt.Println("Snackbar has not been shown yet.")
		return
	}
	fmt.Printf("Message: %v\nTimeout: %v\nActionHandler: %v, ActionText: %v\n",
		o.Get("message"),
		o.Get("timeout"),
		o.Get("actionHandler").Interface(),
		o.Get("actionText"),
	)
	fmt.Printf("Multiline: %v, ActionOnBottom: %v\n",
		o.Get("multiline"),
		o.Get("actionOnBottom"),
	)
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
