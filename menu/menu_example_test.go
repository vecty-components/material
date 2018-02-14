package menu_test

import (
	"fmt"
	"log"

	"agamigo.io/material/internal/mdctest"
	"agamigo.io/material/menu"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material menu component.
	c := &menu.M{}

	// Set up a DOM HTMLElement suitable for a checkbox.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}

	printStatus(c)
	printState(c)
	c.OpenFocus(2)
	c.QuickOpen = true
	c.SetAnchorCorner(menu.BOTTOM_END)
	ms := c.AnchorMargins()
	ms.Left = ms.Left + 50
	ms.Right = ms.Right + 100
	ms.Top = ms.Top + 150
	ms.Bottom = ms.Bottom + 200
	c.SetAnchorMargins(ms)
	printState(c)
	c.Open = false
	c.ItemsContainer().Call("removeChild", c.Items()[0])
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCMenu
	//
	// Open: false, QuickOpen, false, Items: 12
	// AnchorCorner: 8, ItemsContainer: [object HTMLUListElement]
	// AnchorMargins
	// [Go] Left: 0, Right: 0, Top: 0, Bottom: 0
	// [JS] Left: 0, Right: 0, Top: 0, Bottom: 0
	//
	// Open: true, QuickOpen, true, Items: 12
	// AnchorCorner: 13, ItemsContainer: [object HTMLUListElement]
	// AnchorMargins
	// [Go] Left: 50, Right: 100, Top: 150, Bottom: 200
	// [JS] Left: 50, Right: 100, Top: 150, Bottom: 200
	//
	// Open: false, QuickOpen, true, Items: 11
	// AnchorCorner: 13, ItemsContainer: [object HTMLUListElement]
	// AnchorMargins
	// [Go] Left: 50, Right: 100, Top: 150, Bottom: 200
	// [JS] Left: 50, Right: 100, Top: 150, Bottom: 200
}

func printStatus(c *menu.M) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *menu.M) {
	fmt.Println()
	fmt.Printf("Open: %v, QuickOpen, %v, Items: %v\n",
		c.Open, c.QuickOpen, len(c.Items()))
	fmt.Printf("AnchorCorner: %v, ItemsContainer: %v\n",
		c.AnchorCorner(), c.ItemsContainer())
	jsMargins := c.Component().Get("foundation_").Get("anchorMargin_")
	fmt.Println("AnchorMargins")
	fmt.Printf("[Go] Left: %v, Right: %v, Top: %v, Bottom: %v\n",
		c.AnchorMargins().Left,
		c.AnchorMargins().Right,
		c.AnchorMargins().Top,
		c.AnchorMargins().Bottom,
	)
	fmt.Printf("[JS] Left: %v, Right: %v, Top: %v, Bottom: %v\n",
		jsMargins.Get("left"),
		jsMargins.Get("right"),
		jsMargins.Get("top"),
		jsMargins.Get("bottom"),
	)
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.InitMenu()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
