package snackbar_test

import (
	"fmt"
	"log"

	"agamigo.io/material/internal/mdctest"
	"agamigo.io/material/snackbar"
	"github.com/gopherjs/gopherwasm/js"
)

func Example() {
	// Create a new instance of a material snackbar component.
	c := snackbar.New()
	printName(c)
	printState(c)
	c.Message = "snackbar message. before Start()"
	c.ActionHandler = func() { fmt.Println("Action Handled! before Start()") }
	c.ActionText = "Action Button Text. before Start()"
	c.Timeout = 1000 // 1 second
	c.MultiLine = true
	c.ActionOnBottom = true
	c.DismissOnAction = true
	printState(c)

	// Set up a DOM HTMLElement suitable for a snackbar.
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
	c.Message = "snackbar message. after Start()"
	c.ActionHandler = func() { fmt.Println("Action Handled! after Start()") }
	c.ActionText = "Action Button Text. after Start()"
	c.Timeout = 2000 // 2 second
	c.MultiLine = false
	c.ActionOnBottom = false
	c.DismissOnAction = false
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
	printState(c)

	// Output:
	// MDCSnackbar
	// DismissOnAction: false
	// [Go] Message: , Timeout: 2750, ActionHandler Exists: false, ActionText: Multiline: false, ActionOnBottom: false
	// DismissOnAction: true
	// [Go] Message: snackbar message. before Start(), Timeout: 1000, ActionHandler Exists: true, ActionText: Action Button Text. before Start()Multiline: true, ActionOnBottom: true
	// DismissOnAction: true
	// [Go] Message: snackbar message. before Start(), Timeout: 1000, ActionHandler Exists: true, ActionText: Action Button Text. before Start()Multiline: false, ActionOnBottom: true
	// Snackbar has not been shown yet.
	// DismissOnAction: false
	// [Go] Message: snackbar message. after Start(), Timeout: 2000, ActionHandler Exists: true, ActionText: Action Button Text. after Start()Multiline: false, ActionOnBottom: false
	// [Js] Message: snackbar message. after Start(), Timeout: 2000, ActionHandler: 0x1, ActionText: Action Button Text. after Start()Multiline: false, ActionOnBottom: false
	// DismissOnAction: false
	// [Go] Message: snackbar message. after Start(), Timeout: 2000, ActionHandler Exists: true, ActionText: Action Button Text. after Start()Multiline: false, ActionOnBottom: false
	// [Js] Message: snackbar message. after Start(), Timeout: 2000, ActionHandler: 0x1, ActionText: Action Button Text. after Start()Multiline: false, ActionOnBottom: false
}

func printName(c *snackbar.S) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *snackbar.S) {
	fmt.Printf("DismissOnAction: %v\n",
		c.Component().Get("dismissOnAction"))

	fmt.Printf("[Go] Message: %v, Timeout: %v, ActionHandler Exists: %v,"+
		" ActionText: %v", c.Message, c.Timeout,
		c.Component().Get("actionHandler") != nil, c.ActionText)
	fmt.Printf("Multiline: %v, ActionOnBottom: %v\n",
		c.MultiLine, c.ActionOnBottom)
	if c.Component().Get("foundation_") != js.Undefined() {
		o := c.Component().Get("foundation_").Get("snackbarData_")
		if o == nil {
			fmt.Println("Snackbar has not been shown yet.")
			return
		}
		fmt.Printf("[Js] Message: %v, Timeout: %v, ActionHandler: %v, ActionText: %v",
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
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
