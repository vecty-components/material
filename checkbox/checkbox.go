package checkbox // import "agamigo.io/vecty-material/checkbox"

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

type CB interface {
	vecty.Component
	vecty.Mounter
	vecty.Unmounter
	Checked() bool
	SetChecked(v bool)
	Disabled() bool
	SetDisabled(v bool)
	Indeterminate() bool
	SetIndeterminate(v bool)
	Value() string
	SetValue(v string)
	ID() string
	Element() *js.Object
	AddClass(c string)
	DelClass(c string)
	getClasses() vecty.ClassMap
}

func render(c CB) vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("mdc-checkbox"),
			vecty.MarkupIf(c.Disabled(),
				vecty.Class("mdc-checkbox--disabled"),
			),
			c.getClasses(),
		),
		elem.Input(
			vecty.Markup(
				vecty.Class("mdc-checkbox__native-control"),
				prop.ID(c.ID()),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(c.Checked()),
				vecty.MarkupIf(c.Value() != "",
					prop.Value(c.Value())),
				vecty.MarkupIf(c.Disabled(),
					vecty.Property("disabled", true),
				),
				vecty.MarkupIf(c.Indeterminate(),
					vecty.Property("indeterminate", true)),
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
}
