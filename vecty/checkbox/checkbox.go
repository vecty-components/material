package checkbox

import (
	"time"

	mdccheckbox "agamigo.io/material-components-go/mdc/checkbox"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type MDCCheckboxState int

const (
	CHECKED MDCCheckboxState = iota + 1
	UNCHECKED
	INDETERMINATE
)

type C struct {
	*mdccheckbox.C
	vecty.Core
	state   MDCCheckboxState `js:"state"`
	enabled bool             `js:"enabled"`
}

func New() *C {
	c := &C{}
	c.C = mdccheckbox.New()
	c.state = CHECKED
	c.enabled = true
	return c
}

func (c *C) SetState(s MDCCheckboxState, enabled bool) {
	c.state = s
	c.enabled = enabled
	vecty.Rerender(c)
}

func (c *C) applyState() vecty.Applyer {
	switch c.state {
	case UNCHECKED:
		return prop.Checked(false)
	case CHECKED:
		return prop.Checked(true)
	case INDETERMINATE:
		return vecty.Property("indeterminate", true)
	}
	return nil
}

func (c *C) handleChange(e *vecty.Event) {
	state := c.C.Object.Get("foundation_").Get("currentCheckState_").String()
	switch state {
	case "checked":
		c.state = CHECKED
	case "unchecked":
		c.state = UNCHECKED
	case "indeterminate":
		c.state = INDETERMINATE
	}
}

func (c *C) Render() vecty.ComponentOrHTML {
	println("Render checkbox called")
	e := elem.Div(
		vecty.Markup(
			// TODO: Report bug?
			// vecty.ClassMap(c.C.Classes),
			vecty.Class("mdc-checkbox"),
		),
		elem.Input(
			vecty.Markup(
				prop.Type(prop.TypeCheckbox),
				vecty.Class("mdc-checkbox__native-control"),
				prop.ID("native-js-checkbox"),
				c.applyState(),
				vecty.MarkupIf(
					c.state != INDETERMINATE,
					vecty.Property("indeterminate", false),
				),
				vecty.MarkupIf(
					!c.enabled,
					vecty.Attribute("disabled", ""),
				),
				event.Change(c.handleChange),
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

func (c *C) SkipRender(prev vecty.Component) bool {
	p, ok := prev.(*C)
	if !ok {
		return false
	}

	if c.state != p.state || c.enabled != p.enabled {
		return false
	}

	return true
}

func (c *C) Mount() {
	println("Mount checkbox called")
	c.C.Start()
	go c.testCB()
}

func (c *C) Unmount() {
	println("Unmount checkbox called")
	c.C.Stop()
}

func (c *C) testCB() {
	for _ = range time.Tick(5 * time.Second) {
		vecty.Rerender(c)
		println(c.state)
		println(c.enabled)
		// switch c.state {
		// case UNCHECKED:
		// 	c.SetState(CHECKED, false)
		// case INDETERMINATE:
		// 	c.SetState(UNCHECKED, true)
		// case CHECKED:
		// 	c.SetState(INDETERMINATE, false)
		// }
	}
}
