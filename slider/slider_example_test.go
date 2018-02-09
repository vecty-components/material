package slider_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/slider"
)

func Example() {
	// Create a new instance of a material slider component.
	c, err := slider.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an slider.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}

	fmt.Printf("%s\n", c)
	fmt.Printf("Value: %v, Min: %v, Max %v, Step %v, Disabled: %v\n",
		c.GetObject().Get("foundation_").Get("value_"),
		c.GetObject().Get("foundation_").Get("min_"),
		c.GetObject().Get("foundation_").Get("max_"),
		c.GetObject().Get("foundation_").Get("step_"),
		c.GetObject().Get("foundation_").Get("disabled_"))

	c.Value = 10.0
	c.Min = 5.5
	c.Max = 50.0
	c.Step = 5.0
	c.Disabled = true

	fmt.Printf("Value: %v, Min: %v, Max %v, Step %v, Disabled: %v\n",
		c.GetObject().Get("foundation_").Get("value_"),
		c.GetObject().Get("foundation_").Get("min_"),
		c.GetObject().Get("foundation_").Get("max_"),
		c.GetObject().Get("foundation_").Get("step_"),
		c.GetObject().Get("foundation_").Get("disabled_"))

	// Output:
	// {"component":"MDCSlider","status":"stopped"}
	// {"component":"MDCSlider","status":"running"}
	// Value: 0, Min: 0, Max 100, Step 0, Disabled: false
	// Value: 10, Min: 5.5, Max 50, Step 5, Disabled: true
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
