package linearprogress_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/linearprogress"
	"agamigo.io/material/mdctest"
)

func Example() {
	// Create a new instance of a material linearprogress component.
	c, err := linearprogress.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for a linearprogress.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)
	fmt.Printf("Determinate: %v, Progress: %v, Buffer: %v, Reverse: %v\n",
		c.Determinate, c.Progress, c.Buffer, c.Reverse)

	err = c.Open()
	if err != nil {
		log.Fatalf("Unable to Open component %s: %v\n", c, err.Error())
	}

	c.Determinate = true
	c.Progress = 54
	c.Buffer = 100
	c.Reverse = true

	err = c.Close()
	if err != nil {
		log.Fatalf("Unable to Close component %s: %v\n", c, err.Error())
	}

	fmt.Printf("Determinate: %v, Progress: %v, Buffer: %v, Reverse: %v\n",
		c.Determinate, c.Progress, c.Buffer, c.Reverse)

	// Output:
	// {"component":"MDCLinearProgress","status":"stopped"}
	// {"component":"MDCLinearProgress","status":"running"}
	// Determinate: false, Progress: 0, Buffer: 0, Reverse: false
	// Determinate: true, Progress: 54, Buffer: 100, Reverse: true
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
