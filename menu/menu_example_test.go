package menu_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/menu"
)

func Example() {
	// Create a new instance of a material menu component.
	c, err := menu.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an menu.
	mdctest.Dom.SetHTML(`<html><body>` + componenthtml.HTML(c.CType()) +
		`</body></html>`)

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)
	fmt.Printf("Open: %v, QuickOpen, %v, Items: %v\n",
		c.Open, c.QuickOpen, c.Items.Length())
	fmt.Printf("AnchorCorner: %v, ItemsContainer: %v\n",
		c.AnchorCorner(), c.ItemsContainer)
	fmt.Printf("AnchorMargins\nLeft: %v, Right %v, Top %v, Bottom %v\n",
		c.LeftMargin, c.RightMargin, c.TopMargin, c.BottomMargin)

	c.OpenFocus(2)
	c.QuickOpen = true
	c.SetAnchorCorner(menu.BOTTOM_END)
	c.LeftMargin = 50
	c.RightMargin = 100
	c.TopMargin = 150
	c.BottomMargin = 200

	fmt.Printf("Open: %v, QuickOpen, %v, Items: %v\n",
		c.Open, c.QuickOpen, c.Items.Length())
	fmt.Printf("AnchorCorner: %v, ItemsContainer: %v\n",
		c.AnchorCorner(), c.ItemsContainer)
	fmt.Printf("AnchorMargins\nLeft: %v, Right %v, Top %v, Bottom %v\n",
		c.LeftMargin, c.RightMargin, c.TopMargin, c.BottomMargin)

	c.Open = false
	c.ItemsContainer.Call("removeChild", c.Items.Index(0))

	fmt.Printf("Open: %v, QuickOpen, %v, Items: %v\n",
		c.Open, c.QuickOpen, c.Items.Length())
	fmt.Printf("AnchorCorner: %v, ItemsContainer: %v\n",
		c.AnchorCorner(), c.ItemsContainer)
	fmt.Printf("AnchorMargins\nLeft: %v, Right %v, Top %v, Bottom %v\n",
		c.LeftMargin, c.RightMargin, c.TopMargin, c.BottomMargin)

	// Output:
	// {"component":"MDCMenu","status":"stopped"}
	// {"component":"MDCMenu","status":"running"}
	// Open: false, QuickOpen, false, Items: 12
	// AnchorCorner: 8, ItemsContainer: [object HTMLUListElement]
	// AnchorMargins
	// Left: 0, Right 0, Top 0, Bottom 0
	// Open: true, QuickOpen, true, Items: 12
	// AnchorCorner: 13, ItemsContainer: [object HTMLUListElement]
	// AnchorMargins
	// Left: 50, Right 100, Top 150, Bottom 200
	// Open: false, QuickOpen, true, Items: 11
	// AnchorCorner: 13, ItemsContainer: [object HTMLUListElement]
	// AnchorMargins
	// Left: 50, Right 100, Top 150, Bottom 200
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.InitMenu()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
