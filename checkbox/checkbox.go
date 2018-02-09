package checkbox // import "agamigo.io/vecty-material/checkbox"

import (
	mdccheckbox "agamigo.io/material/checkbox"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type CB struct {
	vecty.Core
	*mdccheckbox.CB
}

func New() (checkbox *CB, err error) {
	c, err := mdccheckbox.New()
	if err != nil {
		return nil, err
	}

	return &CB{CB: c}, nil
}

func (c *CB) Render() vecty.ComponentOrHTML {
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

func (c *CB) Mount() {
	c.Start()
}

func (c *CB) Unmount() {
	c.Stop()
}
