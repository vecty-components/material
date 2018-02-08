package icontoggle_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/icontoggle"
	"agamigo.io/material/mdctest"
)

func Example() {
	// Create a new instance of a material icontoggle component.
	c, err := icontoggle.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an icontoggle.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)
	fmt.Printf("MDC On: %v, MDC Disabled: %v\n",
		c.GetObject().Get("on"), c.GetObject().Get("disabled"))

	c.Disabled = true
	c.On = true
	fmt.Printf("MDC On: %v, MDC Disabled: %v\n",
		c.GetObject().Get("on"), c.GetObject().Get("disabled"))

	// Output:
	// {"component":"MDCIconToggle","status":"stopped"}
	// {"component":"MDCIconToggle","status":"running"}
	// MDC On: false, MDC Disabled: false
	// MDC On: true, MDC Disabled: true
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
