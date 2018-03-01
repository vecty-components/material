package main

import (
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/formfield"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// PageView is our main page component.
type PageView struct {
	vecty.Core
}

func main() {
	pv := &PageView{}
	vecty.RenderBody(pv)
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	heroCB := checkbox.NewBasic("hero-checkbox")
	defaultCB := checkbox.NewBasic("basic-checkbox")
	defaultFF := &formfield.FF{
		Label: "Default checkbox",
		Input: defaultCB,
	}
	defaultCBU := checkbox.NewUpgraded("native-js-checkbox")
	disabledCB := checkbox.NewBasic("basic-disabled-checkbox")
	disabledCB.SetDisabled(true)
	indeterminateCB := checkbox.NewBasic("basic-indeterminate-checkbox")
	indeterminateCB.SetIndeterminate(true)
	indeterminateCB.SetChecked(true)
	indeterminateCBU := checkbox.NewUpgraded("native-js-checkbox-indeterminate")
	indeterminateCBU.SetIndeterminate(true)
	indeterminateCBU.SetChecked(true)
	customAllCB := checkbox.NewBasic("basic-custom-checkbox-all")
	customAllCB.AddClass("demo-checkbox--custom-all")
	customAllCBU := checkbox.NewUpgraded("native-js-checkbox-custom-all")
	customAllCBU.AddClass("demo-checkbox--custom-all")
	customSomeCB := checkbox.NewBasic("basic-custom-checkbox-stroke-and-fill")
	customSomeCB.AddClass("demo-checkbox--custom-stroke-and-fill")
	customSomeCBU := checkbox.NewUpgraded("native-js-checkbox-custom-stroke-and-fill")
	customSomeCBU.AddClass("demo-checkbox--custom-stroke-and-fill")

	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		&common.ToolbarHeader{Title: "Checkbox"},
		elem.Main(
			elem.Div(
				vecty.Markup(
					vecty.Class("mdc-toolbar-fixed-adjust"),
				),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class("hero"),
				),
				heroCB,
				elem.Label(
					vecty.Markup(
						prop.ID("hero-checkbox-label"),
						prop.For("hero-checkbox"),
					),
					vecty.Text("Checkbox"),
				),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class("example"),
				),
				elem.Heading2(
					vecty.Text("CSS Only"),
				),
				elem.Div(
					defaultFF,
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						&button.B{
							Label:   vecty.Text("Toggle RTL"),
							Stroked: true,
							Compact: true,
							ClickHandler: func(e *vecty.Event) {
								ff := e.Target.Get("parentElement")
								ff = ff.Get("parentElement")
								dir := ff.Call("hasAttribute", "dir").Bool()
								if dir {
									ff.Call("removeAttribute", "dir")
									return
								}
								ff.Call("setAttribute", "dir", "rtl")
							},
						},
						&button.B{
							Label: vecty.List{
								vecty.Text("Toggle "),
								elem.Code(
									vecty.Text("--align-end"),
								),
							},
							Stroked: true,
							Compact: true,
							ClickHandler: func(e *vecty.Event) {
								defaultFF.AlignEnd = !defaultFF.AlignEnd
								vecty.Rerender(defaultFF)
							},
						},
					),
				),
				elem.Div(
					&formfield.FF{
						Label: "Disabled checkbox",
						Input: disabledCB,
					},
					elem.Div(
						&formfield.FF{
							Label: "Indeterminate checkbox",
							Input: indeterminateCB,
						},
					),
					elem.Div(
						&formfield.FF{
							Label: "Custom colored checkbox (stroke, fill, ripple, and focus)",
							Input: customAllCB,
						},
					),
					elem.Div(
						&formfield.FF{
							Label: "Custom colored checkbox (stroke and fill only)",
							Input: customSomeCB,
						},
					),
				),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class("example"),
					vecty.Class("example--with-js"),
				),
				elem.Heading2(
					vecty.Text("With JavaScript"),
				),
				elem.Div(
					&formfield.FF{
						Label: "Default checkbox",
						Input: defaultCBU,
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(defaultCBU),
						makeDisabledButton(defaultCBU),
					),
				),
				elem.Div(
					&formfield.FF{
						Label: "Indeterminate checkbox",
						Input: indeterminateCBU,
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(indeterminateCB),
						makeDisabledButton(indeterminateCB),
					),
				),
				elem.Div(
					&formfield.FF{
						Label: "Custom colored checkbox (stroke, fill, ripple, and focus)",
						Input: customAllCBU,
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(customAllCB),
						makeDisabledButton(customAllCB),
					),
				),
				elem.Div(
					&formfield.FF{
						Label: "Custom colored checkbox (stroke and fill only)",
						Input: customSomeCBU,
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(customSomeCB),
						makeDisabledButton(customSomeCB),
					),
				),
			),
		),
	)
}

func makeButton(h func(*vecty.Event), l vecty.List) *button.B {
	return &button.B{
		Label:        l,
		ClickHandler: h,
		Stroked:      true,
		Compact:      true,
	}
}

func makeIndeterminateButton(cb checkbox.CBInterface) *button.B {
	return makeButton(
		func(e *vecty.Event) {
			cb.SetIndeterminate(
				!cb.Indeterminate())
		},
		vecty.List{vecty.Text("Toggle "),
			elem.Code(
				vecty.Text("indeterminate"),
			),
		},
	)
}

func makeDisabledButton(cb checkbox.CBInterface) *button.B {
	return makeButton(
		func(e *vecty.Event) {
			cb.SetDisabled(
				!cb.Disabled())
		},
		vecty.List{
			vecty.Text("Toggle "),
			elem.Code(
				vecty.Text("disabled"),
			),
		},
	)
}
