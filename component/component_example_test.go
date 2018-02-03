package component_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/checkbox"
)

func Example() {
	// Create a new instance of a checkbox component.
	c, err := checkbox.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	dom.SetHTML("<html><body>" + c.HTML() +
		"</body></html>")

	// Start the component after its HTMLElement is instantiated.
	// In this case it was instantiated in emulateDOM()
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Output:
	// {"component":"MDCCheckbox","status":"stopped"}
	// {"component":"MDCCheckbox","status":"running"}
}
