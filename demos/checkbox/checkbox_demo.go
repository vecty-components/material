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
	testFF     *formfield.FF
}

func main() {
	cdv := &checkboxDemoView{
		checkboxes: map[string]*checkbox.CB{
			HERO_ID:                checkbox.New().WithID(HERO_ID),
			BASIC_ID:               checkbox.New().WithID(BASIC_ID),
			BASIC_DISABLED_ID:      checkbox.New().WithID(BASIC_DISABLED_ID),
			BASIC_INDETERMINATE_ID: checkbox.New().WithID(BASIC_INDETERMINATE_ID),
			BASIC_CUSTOM_ALL_ID:    checkbox.New().WithID(BASIC_CUSTOM_ALL_ID),
			BASIC_CUSTOM_SOME_ID:   checkbox.New().WithID(BASIC_CUSTOM_SOME_ID),
			JS_ID:                  checkbox.New().WithID(JS_ID),
			JS_INDETERMINATE_ID:    checkbox.New().WithID(JS_INDETERMINATE_ID),
			JS_CUSTOM_ALL_ID:       checkbox.New().WithID(JS_CUSTOM_ALL_ID),
			JS_CUSTOM_SOME_ID:      checkbox.New().WithID(JS_CUSTOM_SOME_ID),
		},
	}
	for id, cb := range cdv.checkboxes {
		switch id {
		case BASIC_ID, BASIC_DISABLED_ID, BASIC_INDETERMINATE_ID,
			BASIC_CUSTOM_ALL_ID, BASIC_CUSTOM_SOME_ID:
			cb.WithBasic()
		}
		switch id {
		case BASIC_DISABLED_ID:
			cb.Disabled = true
		case BASIC_INDETERMINATE_ID, JS_INDETERMINATE_ID:
			cb.Checked = true
			cb.Indeterminate = true
		case BASIC_CUSTOM_ALL_ID, JS_CUSTOM_ALL_ID:
			cb.WithClass("demo-checkbox--custom-all")
		case BASIC_CUSTOM_SOME_ID, JS_CUSTOM_SOME_ID:
			cb.WithClass("demo-checkbox--custom-stroke-and-fill")
		}
	}
	cdv.defaultFF = formfield.New().WithConfig(
		&formfield.Config{
			Label: "Default checkbox",
			Input: cdv.checkboxes[BASIC_ID],
		},
	)
	vecty.RenderBody(cdv)
}

// Render implements the vecty.Component interface.
func (c *checkboxDemoView) Render() vecty.ComponentOrHTML {
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
								c.defaultFF.AlignEnd = !c.defaultFF.AlignEnd
								vecty.Rerender(c.defaultFF)
							},
						},
					),
				),
				elem.Div(
					formfield.New().WithConfig(
						&formfield.Config{
							Label: "Disabled checkbox",
							Input: c.checkboxes[BASIC_DISABLED_ID],
						},
					),
					elem.Div(
						formfield.New().WithConfig(
							&formfield.Config{
								Label: "Indeterminate checkbox",
								Input: c.checkboxes[BASIC_INDETERMINATE_ID],
							},
						),
					),
					elem.Div(
						formfield.New().WithConfig(
							&formfield.Config{
								Label: "Custom colored checkbox (stroke, fill, ripple, and focus)",
								Input: c.checkboxes[BASIC_CUSTOM_ALL_ID],
							},
						),
					),
					elem.Div(
						formfield.New().WithConfig(
							&formfield.Config{
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
					formfield.New().WithConfig(
						&formfield.Config{
							Label: "Default checkbox",
							Input: c.checkboxes[JS_ID],
						},
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(c.checkboxes[JS_ID]),
						makeDisabledButton(c.checkboxes[JS_ID]),
					),
				),
				elem.Div(
					formfield.New().WithConfig(
						&formfield.Config{
							Label: "Indeterminate checkbox",
							Input: c.checkboxes[JS_INDETERMINATE_ID],
						},
					),
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
					formfield.New().WithConfig(
						&formfield.Config{
							Label: "Custom colored checkbox (stroke, fill, ripple, and focus)",
							Input: c.checkboxes[JS_CUSTOM_ALL_ID],
						},
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("demo-toggle-group"),
						),
						makeIndeterminateButton(c.checkboxes[JS_CUSTOM_ALL_ID]),
						makeDisabledButton(c.checkboxes[JS_CUSTOM_ALL_ID]),
					),
				),
				elem.Div(
					formfield.New().WithConfig(
						&formfield.Config{
							Label: "Custom colored checkbox (stroke and fill only)",
							Input: c.checkboxes[JS_CUSTOM_SOME_ID],
						},
					),
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

func makeButton(h func(*vecty.Event), l vecty.List, class string) *button.B {
	return &button.B{
		Label:        l,
		ClickHandler: h,
		Stroked:      true,
		Compact:      true,
		Classes:      vecty.ClassMap{class: true},
	}
}

func makeIndeterminateButton(cb *checkbox.CB) *button.B {
	return makeButton(
		func(e *vecty.Event) {
			cb.Indeterminate = !cb.Indeterminate
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
		func(e *vecty.Event) {
			cb.Disabled = !cb.Disabled
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

func newCBsWithIDs(ids ...string) map[string]*checkbox.CB {
	cbs := make(map[string]*checkbox.CB, len(ids))
	for _, id := range ids {
		cbs[id] = checkbox.New().WithID(id)
	}
	return cbs
}
