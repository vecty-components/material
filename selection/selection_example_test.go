package selection_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/selection"
)

func Example() {
	// Create a new instance of a material selection component.
	c, err := selection.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an selection.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}

	fmt.Printf("%s\n", c)
	fmt.Printf("SelectedIndex: %v, SelectedString: %v, Disabled: %v\n",
		c.GetObject().Get("selectedIndex"), c.SelectedString(),
		c.GetObject().Get("disabled"))
	fmt.Printf("SelectedElem: %v\nOptions: %v\n",
		c.SelectedElem(), c.Options())

	c.SelectedIndex = 0
	c.Disabled = true

	l := c.GetObject().Get("root_").Call("querySelector", ".mdc-list")
	l.Call("removeChild", c.Options().Index(1))

	fmt.Printf("SelectedIndex: %v, SelectedString: %v, Disabled: %v\n",
		c.GetObject().Get("selectedIndex"), c.SelectedString(),
		c.GetObject().Get("disabled"))
	fmt.Printf("SelectedElem: %v\nOptions: %v\n",
		c.SelectedElem(), c.Options())

	// Output:
	// {"component":"MDCSelect","status":"stopped"}
	// {"component":"MDCSelect","status":"running"}
	// SelectedIndex: -1, SelectedString: , Disabled: false
	// SelectedElem: [object NodeList]
	// Options: [object HTMLLIElement],[object HTMLLIElement]
	// SelectedIndex: 0, SelectedString: Option #1, Disabled: true
	// SelectedElem: [object NodeList]
	// Options: [object HTMLLIElement]
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
