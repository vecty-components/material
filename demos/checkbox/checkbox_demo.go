package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/lithammer/dedent"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/base/applyer"
	"github.com/vecty-material/material/button"
	"github.com/vecty-material/material/checkbox"
	"github.com/vecty-material/material/demos/common"
	"github.com/vecty-material/material/formfield"
)

const (
	HERO_ID                = "hero-checkbox"
	BASIC_ID               = "basic-checkbox"
	BASIC_DISABLED_ID      = "basic-disabled-checkbox"
	BASIC_INDETERMINATE_ID = "basic-indeterminate-checkbox"
	BASIC_CUSTOM_ALL_ID    = "basic-custom-checkbox-all"
	BASIC_CUSTOM_SOME_ID   = "basic-custom-checkbox-stroke-and-fill"
	JS_ID                  = "native-js-checkbox"
	JS_INDETERMINATE_ID    = "native-js-checkbox-indeterminate"
	JS_CUSTOM_ALL_ID       = "native-js-checkbox-custom-all"
	JS_CUSTOM_SOME_ID      = "native-js-checkbox-custom-stroke-and-fill"
)

// checkboxDemoView is our demo page component.
type checkboxDemoView struct {
	vecty.Core
	checkboxes map[string]*checkbox.CB
	defaultFF  *formfield.FF
}

func main() {
	base.SetViewport()

	vecty.SetTitle("Checkbox - Material Components Catalog")
	vecty.AddStylesheet("https://material-components-web.appspot.com/assets/checkbox.css")

	base.AddIcon("https://material-components-web.appspot.com/images/logo_components_color_2x_web_48dp.png")
	base.AddCSS(dedent.Dedent(`
		.example--with-js .mdc-form-field {
			min-width: 450px;
		}

		.demo-toggle-group {
			display: inline-block;
		}

		@media (max-width: 600px) {
			.mdc-button {
				margin-bottom: 4px;
			}
		}

		.mdc-button code {
			text-transform: none;
		}
	`))

	base.AddResources()

	cdv := &checkboxDemoView{
		checkboxes: map[string]*checkbox.CB{
			HERO_ID:                nil,
			BASIC_ID:               nil,
			BASIC_DISABLED_ID:      nil,
			BASIC_INDETERMINATE_ID: nil,
			BASIC_CUSTOM_ALL_ID:    nil,
			BASIC_CUSTOM_SOME_ID:   nil,
			JS_ID:                  nil,
			JS_INDETERMINATE_ID:    nil,
			JS_CUSTOM_ALL_ID:       nil,
			JS_CUSTOM_SOME_ID:      nil,
		},
	}
	for id, _ := range cdv.checkboxes {
		cdv.checkboxes[id] = &checkbox.CB{
			Input: vecty.Markup(prop.ID(id)),
		}
		c := cdv.checkboxes[id]
		var applyers []vecty.Applyer
		switch id {
		case BASIC_ID, BASIC_DISABLED_ID, BASIC_INDETERMINATE_ID,
			BASIC_CUSTOM_ALL_ID, BASIC_CUSTOM_SOME_ID:
			applyers = append(applyers, applyer.CSSOnly())
		}
		switch id {
		case BASIC_DISABLED_ID:
			c.Disabled = true
		case BASIC_INDETERMINATE_ID, JS_INDETERMINATE_ID:
			c.Checked = true
			c.Indeterminate = true
		case BASIC_CUSTOM_ALL_ID, JS_CUSTOM_ALL_ID:
			applyers = append(applyers,
				vecty.Class("demo-checkbox--custom-all"))

		case BASIC_CUSTOM_SOME_ID, JS_CUSTOM_SOME_ID:
			applyers = append(applyers,
				vecty.Class("demo-checkbox--custom-stroke-and-fill"))
		}
		c.Root = vecty.Markup(applyers...)
	}
	cdv.defaultFF = &formfield.FF{
		Label: "Default checkbox",
		Input: cdv.checkboxes[BASIC_ID],
	}
	vecty.RenderBody(cdv)
}

// Render implements the vecty.Component interface.
func (c *checkboxDemoView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		&common.ToolbarHeader{
			Title:      "Checkbox",
			Navigation: common.NavBack,
		},
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
				c.checkboxes[HERO_ID],
				elem.Label(
					vecty.Markup(
						prop.ID("hero-checkbox-label"),
						prop.For(HERO_ID),
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
					c.defaultFF,
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						&button.B{
							Label:    vecty.Text("Toggle RTL"),
							Outlined: true,
							OnClick: func(thisB *button.B, e *vecty.Event) {
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
							Outlined: true,
							OnClick: func(thisB *button.B,
								e *vecty.Event) {
								c.defaultFF.AlignEnd = !c.defaultFF.AlignEnd
								vecty.Rerender(c.defaultFF)
							},
						},
					),
					elem.Div(
						&formfield.FF{
							Label: "Disabled checkbox",
							Input: c.checkboxes[BASIC_DISABLED_ID],
						},
						elem.Div(
							&formfield.FF{
								Label: "Indeterminate checkbox",
								Input: c.checkboxes[BASIC_INDETERMINATE_ID],
							},
						),
						elem.Div(
							&formfield.FF{
								Label: "Custom colored checkbox (stroke, fill, ripple, and focus)",
								Input: c.checkboxes[BASIC_CUSTOM_ALL_ID],
							},
						),
						elem.Div(
							&formfield.FF{
								Label: "Custom colored checkbox (stroke and fill only)",
								Input: c.checkboxes[BASIC_CUSTOM_SOME_ID],
							},
						),
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
						Input: c.checkboxes[JS_ID],
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(c.checkboxes[JS_ID]),
						makeDisabledButton(c.checkboxes[JS_ID]),
					),
				),
				elem.Div(
					&formfield.FF{
						Label: "Indeterminate checkbox",
						Input: c.checkboxes[JS_INDETERMINATE_ID],
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(
							c.checkboxes[JS_INDETERMINATE_ID]),
						makeDisabledButton(c.checkboxes[JS_INDETERMINATE_ID]),
					),
				),
				elem.Div(
					&formfield.FF{
						Label: "Custom colored checkbox (stroke, fill, ripple, and focus)",
						Input: c.checkboxes[JS_CUSTOM_ALL_ID],
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(c.checkboxes[JS_CUSTOM_ALL_ID]),
						makeDisabledButton(c.checkboxes[JS_CUSTOM_ALL_ID]),
					),
				),
				elem.Div(
					&formfield.FF{
						Label: "Custom colored checkbox (stroke and fill only)",
						Input: c.checkboxes[JS_CUSTOM_SOME_ID],
					},
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(c.checkboxes[JS_CUSTOM_SOME_ID]),
						makeDisabledButton(c.checkboxes[JS_CUSTOM_SOME_ID]),
					),
				),
			),
		),
	)
}

func makeButton(h func(*button.B, *vecty.Event), l vecty.List, class string) *button.B {
	var applyer vecty.Applyer
	if class != "" {
		applyer = vecty.Class(class)
	}
	return &button.B{
		Root:     vecty.Markup(applyer),
		Label:    l,
		OnClick:  h,
		Outlined: true,
	}
}

func makeIndeterminateButton(cb *checkbox.CB) *button.B {
	return makeButton(
		func(thisB *button.B, e *vecty.Event) {
			cb.Indeterminate = !cb.Indeterminate
			vecty.Rerender(cb)
		},
		vecty.List{vecty.Text("Toggle "),
			elem.Code(
				vecty.Text("indeterminate"),
			),
		},
		"toggle-indeterminate",
	)
}

func makeDisabledButton(cb *checkbox.CB) *button.B {
	return makeButton(
		func(thisB *button.B, e *vecty.Event) {
			cb.Disabled = !cb.Disabled
			vecty.Rerender(cb)
		},
		vecty.List{
			vecty.Text("Toggle "),
			elem.Code(
				vecty.Text("disabled"),
			),
		},
		"toggle-disabled",
	)
}
