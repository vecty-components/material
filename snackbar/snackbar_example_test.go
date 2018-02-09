package snackbar_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/snackbar"
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

	fmt.Printf("%+v\n", c.Data)
	c.Message = "snackbar message here"
	c.ActionHandler = func() {
		fmt.Println("Action Handled!")
	}
	c.ActionText = "Action Button Text"
	c.Timeout = 1000 // 1 second
	c.MultiLine = true
	c.ActionOnBottom = true
	c.DismissOnAction = true
	fmt.Printf("%+v\n", c.Data)

	// Output:
	// {"component":"MDCSnackbar","status":"stopped"}
	// {"component":"MDCSnackbar","status":"running"}
	// &{object:0x1 Message:undefined Timeout:2750 ActionHandler:0x1 ActionText:undefined MultiLine:false ActionOnBottom:false}
	// &{object:0x1 Message:snackbar message here Timeout:1000 ActionHandler:0x1 ActionText:Action Button Text MultiLine:true ActionOnBottom:true}

}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
