package checkbox

import (
	"agamigo.io/material/checkbox"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type CB struct {
	*checkbox.CB
	vecty.Core
	id      string
	classes vecty.ClassMap
	basic   bool
	started bool
	element *vecty.HTML
}

type State struct {
	Checked       bool
	Indeterminate bool
	Disabled      bool
	Value         string
}

func New() *CB {
	c := &CB{}
	return c.WithState(&State{})
}

func (c *CB) WithBasic() *CB {
	if c.started {
		err := c.Stop()
		if err != nil {
			print(err)
		}
	}
	c.basic = true
	return c
}

func (c *CB) WithID(id string) *CB {
	c.id = id
	return c
}

func (c *CB) ID() string {
	return c.id
}

func (c *CB) WithClass(class string) *CB {
	if c.classes == nil {
		c.classes = make(vecty.ClassMap, 1)
	}
	c.classes[class] = true
	return c
}

func (c *CB) WithState(s *State) *CB {
	c.CB = checkbox.New()
	c.Checked = s.Checked
	c.Indeterminate = s.Indeterminate
	c.Disabled = s.Disabled
	c.Value = s.Value
	return c
}

func (c *CB) Render() vecty.ComponentOrHTML {
	c.element = elem.Div(
		vecty.Markup(
			vecty.Class("mdc-checkbox"),
			vecty.MarkupIf(c.Disabled,
				vecty.Class("mdc-checkbox--disabled"),
			),
			c.getClasses(),
		),
		elem.Input(
			vecty.Markup(
				vecty.Class("mdc-checkbox__native-control"),
				vecty.MarkupIf(c.ID() != "",
					prop.ID(c.ID()),
				),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(c.CB.Checked),
				vecty.MarkupIf(c.Value != "",
					prop.Value(c.Value),
				),
				vecty.Property("disabled", c.Disabled),
				vecty.Property("indeterminate", c.Indeterminate),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-checkbox__background"),
				vecty.UnsafeHTML(
					`<svg class="mdc-checkbox__checkmark"
							viewBox="0 0 24 24">
							<path class="mdc-checkbox__checkmark-path"
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
	return c.element
}

func (c *CB) getClasses() vecty.ClassMap {
	return c.classes
}

func (c *CB) Mount() {
	if c.CB == nil {
		c.CB = checkbox.New()
	}
	if c.basic {
		return
	}
	if c.element == nil {
		panic("Element is nil while mounting upgradedCB.")
	}
	e := c.element.Node()
	if e == nil || e == js.Undefined {
		panic("Element is nil while mounting upgradedCB.")
	}
	err := c.Start(e)
	if err != nil {
		panic(err)
	}
	c.started = true
}

func (c *CB) Unmount() {
	err := c.Stop()
	if err != nil {
		panic(err)
	}
	c.started = false
}
