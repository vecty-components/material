package main

import (
	"agamigo.io/vecty-material/checkbox"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// PageView is our main page component.
type PageView struct {
	vecty.Core
}

// form-field + checkbox component
type FFCB struct {
	vecty.Core
	Checkbox checkbox.CB
	label    string
	alignEnd bool
}

func main() {
	vecty.SetTitle("Checkbox - Material Components Catalog")
	pv := &PageView{}
	vecty.RenderBody(pv)
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	heroCB := checkbox.NewBasic("hero-checkbox")
	defaultFFCB := &FFCB{
		label:    "Default checkbox",
		Checkbox: checkbox.NewBasic("basic-checkbox"),
	}
	defaultFFCBU := &FFCB{
		label:    "Default checkbox",
		Checkbox: checkbox.NewUpgraded("native-js-checkbox"),
	}
	disabledFFCB := &FFCB{
		label:    "Disabled checkbox",
		Checkbox: checkbox.NewBasic("basic-disabled-checkbox"),
	}
	disabledFFCB.Checkbox.SetDisabled(true)
	indeterminateFFCB := &FFCB{
		label:    "Indeterminate checkbox",
		Checkbox: checkbox.NewBasic("basic-indeterminate-checkbox"),
	}
	indeterminateFFCBU := &FFCB{
		label:    "Indeterminate checkbox",
		Checkbox: checkbox.NewUpgraded("native-js-checkbox-indeterminate"),
	}
	indeterminateFFCB.Checkbox.SetIndeterminate(true)
	indeterminateFFCB.Checkbox.SetChecked(true)
	indeterminateFFCBU.Checkbox.SetIndeterminate(true)
	indeterminateFFCBU.Checkbox.SetChecked(true)
	customAllFFCB := &FFCB{
		label:    "Custom colored checkbox (stroke, fill, ripple, and focus)",
		Checkbox: checkbox.NewBasic("basic-custom-checkbox-all"),
	}
	customAllFFCBU := &FFCB{
		label:    "Custom colored checkbox (stroke, fill, ripple, and focus)",
		Checkbox: checkbox.NewUpgraded("native-js-checkbox-custom-all"),
	}
	customAllFFCB.Checkbox.AddClass("demo-checkbox--custom-all")
	customAllFFCBU.Checkbox.AddClass("demo-checkbox--custom-all")
	customSomeFFCB := &FFCB{
		label:    "Custom colored checkbox (stroke and fill only)",
		Checkbox: checkbox.NewBasic("basic-custom-checkbox-stroke-and-fill"),
	}
	customSomeFFCBU := &FFCB{
		label:    "Custom colored checkbox (stroke and fill only)",
		Checkbox: checkbox.NewUpgraded("native-js-checkbox-custom-stroke-and-fill"),
	}
	customSomeFFCB.Checkbox.AddClass("demo-checkbox--custom-stroke-and-fill")
	customSomeFFCBU.Checkbox.AddClass("demo-checkbox--custom-stroke-and-fill")

	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		renderHeader(),
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
					defaultFFCB,
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						renderButton(
							func(e *vecty.Event) {
								ff := e.Target.Get("parentElement")
								ff = ff.Get("parentElement")
								dir := ff.Call("hasAttribute", "dir").Bool()
								if dir {
									ff.Call("removeAttribute", "dir")
									return
								}
								ff.Call("setAttribute", "dir", "rtl")
							},
							vecty.Text("Toggle RTL"),
						),
						renderButton(
							func(e *vecty.Event) {
								defaultFFCB.alignEnd = !defaultFFCB.alignEnd
								vecty.Rerender(defaultFFCB)
							},
							vecty.Text("Toggle "),
							elem.Code(
								vecty.Text("--align-end"),
							),
						),
					),
				),
				elem.Div(
					disabledFFCB,
					elem.Div(
						indeterminateFFCB,
					),
					elem.Div(
						customAllFFCB,
					),
					elem.Div(
						customSomeFFCB,
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
					defaultFFCBU,
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						renderIndeterminateButton(defaultFFCBU.Checkbox),
						renderDisabledButton(defaultFFCBU.Checkbox),
					),
				),
				elem.Div(
					indeterminateFFCBU,
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						renderIndeterminateButton(indeterminateFFCBU.Checkbox),
						renderDisabledButton(indeterminateFFCBU.Checkbox),
					),
				),
				elem.Div(
					customAllFFCBU,
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						renderIndeterminateButton(customAllFFCBU.Checkbox),
						renderDisabledButton(customAllFFCBU.Checkbox),
					),
				),
				elem.Div(
					customSomeFFCBU,
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						renderIndeterminateButton(customSomeFFCBU.Checkbox),
						renderDisabledButton(customSomeFFCBU.Checkbox),
					),
				),
			),
		),
	)
}

func renderHeader() *vecty.HTML {
	return elem.Header(
		vecty.Markup(
			vecty.Class("mdc-toolbar"),
			vecty.Class("mdc-toolbar--fixed"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("mdc-toolbar__row"),
			),
			elem.Section(
				vecty.Markup(
					vecty.Class("mdc-toolbar__section"),
					vecty.Class("mdc-toolbar__section--align-start"),
				),
				elem.Anchor(
					vecty.Markup(
						prop.Href("/"),
						vecty.Class("catalog-back"),
						vecty.Class("mdc-toolbar__menu-icon"),
					),
					elem.Italic(
						vecty.Markup(
							vecty.Class("material-icons"),
						),
						// vecty.Text("&#xE5C4;"),
						vecty.Text("arrow_back"),
					),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("mdc-toolbar__title"),
						vecty.Class("catalog-title"),
					),
					vecty.Text("Checkbox"),
				),
			),
		),
	)
}

func (c *FFCB) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("mdc-form-field"),
			vecty.MarkupIf(c.alignEnd,
				vecty.Class("mdc-form-field--align-end"),
			),
		),
		c.Checkbox,
		elem.Label(
			vecty.Markup(
				prop.For(c.Checkbox.ID()),
			),
			vecty.Text(c.label),
		),
	)
}

func (c *FFCB) Mount() {
	c.Checkbox.Mount()
}

func (c *FFCB) Unmount() {
	c.Checkbox.Unmount()
}

func renderButton(handler func(*vecty.Event), label ...*vecty.HTML) *vecty.HTML {
	list := make(vecty.List, len(label))
	for i, v := range label {
		list[i] = v
	}
	return elem.Button(
		vecty.Markup(
			vecty.Class("mdc-button"),
			vecty.Class("mdc-button--stroked"),
			vecty.Class("mdc-button--compact"),
			prop.Type(prop.TypeButton),
			event.Click(handler),
		),
		list,
	)
}

func renderIndeterminateButton(cb checkbox.CB) *vecty.HTML {
	return renderButton(
		func(e *vecty.Event) {
			cb.SetIndeterminate(
				!cb.Indeterminate())
		},
		vecty.Text("Toggle "),
		elem.Code(
			vecty.Text("indeterminate"),
		),
	)
}

func renderDisabledButton(cb checkbox.CB) *vecty.HTML {
	return renderButton(
		func(e *vecty.Event) {
			cb.SetDisabled(
				!cb.Disabled())
		},
		vecty.Text("Toggle "),
		elem.Code(
			vecty.Text("disabled"),
		),
	)
}
