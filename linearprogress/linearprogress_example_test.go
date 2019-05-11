package linearprogress_test

import (
	"fmt"
	"log"

	"github.com/vecty-material/material/internal/mdctest"
	"github.com/vecty-material/material/linearprogress"
	"github.com/gopherjs/gopherwasm/js"
)

func Example() {
	// Create a new instance of a material linearprogress component.
	c := linearprogress.New()
	printName(c)
	printState(c)
	c.Determinate = false
	c.Progress = .54
	c.Buffer = 1.00
	c.Reverse = true
	printState(c)

	// Set up a DOM HTMLElement suitable for a checkbox.
	js.Global().Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global().Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err.Error())
	}

	printState(c)
	err = c.Open()
	if err != nil {
		log.Fatalf("Unable to Open component %s: %v\n", c.Component().Type,
			err.Error())
	}
	c.Determinate = true
	c.Progress = .33
	c.Buffer = .98
	c.Reverse = false
	err = c.Close()
	if err != nil {
		log.Fatalf("Unable to Close component %s: %v\n", c.Component().Type,
			err.Error())
	}
	printState(c)
	jsTests(c)
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCLinearProgress
	//
	// [Go] Determinate: false, Progress: 0, Buffer: 0, Reverse: false
	//
	// [Go] Determinate: false, Progress: 0.54, Buffer: 1, Reverse: true
	//
	// [Go] Determinate: false, Progress: 0.54, Buffer: 1, Reverse: true
	// [JS] Determinate: false, Progress: 0.54, Buffer: 1, Reverse: true
	//
	// [Go] Determinate: true, Progress: 0.33, Buffer: 0.98, Reverse: false
	// [JS] Determinate: true, Progress: 0.33, Buffer: 0.98, Reverse: false
	//
	// [Go] Determinate: true, Progress: 0.45, Buffer: 0.4, Reverse: false
	// [JS] Determinate: true, Progress: 0.45, Buffer: 0.4, Reverse: false
}

func printName(c *linearprogress.LP) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *linearprogress.LP) {
	fmt.Println()
	fmt.Printf("[Go] Determinate: %v, Progress: %v, Buffer: %v, Reverse: %v\n",
		c.Determinate, c.Progress, c.Buffer, c.Reverse)
	mdcObj := c.Component().Get("foundation_")
	if mdcObj != js.Undefined() {
		fmt.Printf("[JS] Determinate: %v, Progress: %v, Buffer: %v, Reverse: %v\n",
			mdcObj.Get("determinate_"),
			mdcObj.Get("progress_"),
			c.Component().Get("buffer"),
			mdcObj.Get("reverse_"),
		)
	}
}

func jsTests(c *linearprogress.LP) {
	o := c.Component()
	o.Set("determinate", true)
	o.Set("progress", .45)
	o.Set("buffer", .40)
	o.Set("reverse", false)
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
