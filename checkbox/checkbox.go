package checkbox // import "agamigo.io/vecty-material/checkbox"

import (
	"math/rand"

	mdccb "agamigo.io/material/checkbox"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type CB struct {
	*mdccb.CB
	vecty.Core
	id string
}

func New() *CB {
	cb := &CB{}
	cb.CB = &mdccb.CB{}
	return cb
}

func (c *CB) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("mdc-checkbox"),
			prop.ID(c.String()),
		),
		elem.Input(
			vecty.Markup(
				vecty.Class("mdc-checkbox__native-control"),
				prop.Type(prop.TypeCheckbox),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-checkbox__background"),
				vecty.UnsafeHTML(
					`<svg class="mdc-checkbox__checkmark"
						viewBox="0 0 24 24">
					<path class="mdc-checkbox__checkmark__path"
						fill="none"
						stroke="white"
						d="M1.73,12.91 8.1,19.28 22.79,4.59"/>
					</svg>`,
				),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-checkbox__mixedmark"),
				),
			),
		),
	)
}

func (c *CB) Mount() {
	e := js.Global.Get("document").Call("getElementById", c.String())
	err := c.Start(e)
	if err != nil {
		panic(err)
	}
}

func (c *CB) Unmount() {
	err := c.Stop()
	if err != nil {
		panic(err)
	}
}

func (c *CB) String() string {
	if c.id == "" {
		c.id = "MDCCheckbox-" + string(rand.Intn(1000))
	}
	return c.id
}
