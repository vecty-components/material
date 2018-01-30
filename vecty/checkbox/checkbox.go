package checkbox

import (
	mdccheckbox "agamigo.io/material/component/checkbox"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type C interface {
	vecty.Component
	mdccheckbox.C
}

type checkbox struct {
	vecty.Core
	mdccheckbox.C
}

func New() C {
	return &checkbox{
		C: mdccheckbox.New(),
	}
}

func (c *checkbox) Render() vecty.ComponentOrHTML {
	e := elem.Div(
		vecty.Markup(
			vecty.Class("mdc-checkbox"),
		),
		elem.Input(
			vecty.Markup(
				vecty.Class("mdc-checkbox__native-control"),
				prop.Type(prop.TypeCheckbox),
				prop.ID("native-js-checkbox"),
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
	return e
}

func (c *checkbox) Mount() {
	c.Start()
}

func (c *checkbox) Unmount() {
	c.Stop()
}
