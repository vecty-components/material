package checkbox_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/checkbox"
	"agamigo.io/material/mdctest"
)

func Example() {
	// Create a new instance of a material checkbox component.
	c, err := checkbox.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up the DOM with an HTMLElement suitable for a checkbox.
	mdctest.Dom.SetHTML("<html><body>" + c.HTML() +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement in the DOM.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Output:
	// {"component":"MDCCheckbox","status":"stopped"}
	// {"component":"MDCCheckbox","status":"running"}
}
