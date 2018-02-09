package snackbar_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/snackbar"
	"github.com/davecgh/go-spew/spew"
)

func Example() {
	// Create a new instance of a material snackbar component.
	c, err := snackbar.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an snackbar.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}

	fmt.Printf("%s\n", c)

	spew.Dump(c.Data)
	c.Message = "snackbar message here"
	c.ActionHandler = func() {
		fmt.Println("Action Handled!")
	}
	c.ActionText = "Action Button Text"
	c.Timeout = 1000 // 1 second
	c.MultiLine = true
	c.ActionOnBottom = true
	c.DismissOnAction = true
	spew.Dump(c.Data)

	// Output:
	// {"component":"MDCSnackbar","status":"stopped"}
	// {"component":"MDCSnackbar","status":"running"}
	// (*snackbar.Data)(<nil>)({
	//  object: (*js.Object)(<nil>)(<already shown>),
	//  Message: (string) (len=9) "undefined",
	//  Timeout: (int) 2750,
	//  ActionHandler: (func()) 0x1,
	//  ActionText: (string) (len=9) "undefined",
	//  MultiLine: (bool) false,
	//  ActionOnBottom: (bool) false
	// })
	// (*snackbar.Data)(<nil>)({
	//  object: (*js.Object)(<nil>)(<already shown>),
	//  Message: (string) (len=21) "snackbar message here",
	//  Timeout: (int) 1000,
	//  ActionHandler: (func()) 0x1,
	//  ActionText: (string) (len=18) "Action Button Text",
	//  MultiLine: (bool) true,
	//  ActionOnBottom: (bool) true
	// })
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
