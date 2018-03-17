package main

import (
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/formfield"
	"agamigo.io/vecty-material/radio"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// radioDemoView is our demo page component.
type radioDemoView struct {
	vecty.Core
}

func main() {
	vecty.RenderBody(&radioDemoView{})
}

// Render implements the vecty.Component interface.
func (c *radioDemoView) Render() vecty.ComponentOrHTML {
	basicRadios := []*radio.R{
		radio.New(
			&base.Props{ID: "ex1-default-radio1"},
			&radio.State{Name: "ex1-default", Checked: true}),
		radio.New(
			&base.Props{ID: "ex1-default-radio2"},
			&radio.State{Name: "ex1-default"}),
		radio.New(
			&base.Props{
				ID:     "ex1-custom-radio1",
				Markup: []vecty.Applyer{vecty.Class("demo-radio--custom")}},
			&radio.State{Name: "ex1-custom", Checked: true}),
		radio.New(
			&base.Props{
				ID:     "ex1-custom-radio2",
				Markup: []vecty.Applyer{vecty.Class("demo-radio--custom")}},
			&radio.State{Name: "ex1-custom"}),
	}
	for _, r := range basicRadios {
		r.Component().MDCState.Basic = true
	}

	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		&common.ToolbarHeader{
			Title:      "Radio",
			Navigation: common.NavBack,
		},
		elem.Main(
			elem.Div(vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust"))),
			elem.Section(
				vecty.Markup(vecty.Class("hero")),
				radio.New(nil, &radio.State{Name: "hero"}),
				radio.New(nil, &radio.State{Name: "hero", Checked: true}),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				elem.Heading2(vecty.Text("With JavaScript")),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					formfield.New(nil, &formfield.State{
						Label: "Default Radio 1",
						Input: radio.New(&base.Props{ID: "ex0-default-radio1"},
							&radio.State{
								Name:    "ex0-default",
								Checked: true,
							},
						),
					}),
					formfield.New(nil, &formfield.State{
						Label: "Default Radio 2",
						Input: radio.New(&base.Props{ID: "ex0-default-radio2"},
							&radio.State{Name: "ex0-default"},
						),
					}),
				),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					formfield.New(nil, &formfield.State{
						Label: "Custom Radio 1",
						Input: radio.New(
							&base.Props{
								ID: "ex0-custom-radio1",
								Markup: []vecty.Applyer{
									vecty.Class("demo-radio--custom")},
							},
							&radio.State{
								Name:    "ex0-custom",
								Checked: true,
							},
						),
					}),
					formfield.New(nil, &formfield.State{
						Label: "Custom Radio 2",
						Input: radio.New(
							&base.Props{
								ID: "ex0-custom-radio2",
								Markup: []vecty.Applyer{
									vecty.Class("demo-radio--custom"),
								},
							},
							&radio.State{Name: "ex0-custom"},
						),
					}),
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				elem.Heading2(vecty.Text("CSS Only")),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					formfield.New(nil, &formfield.State{
						Label: "Default Radio 1",
						Input: basicRadios[0],
					}),
					formfield.New(nil, &formfield.State{
						Label: "Default Radio 2",
						Input: basicRadios[1],
					}),
				),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					formfield.New(nil, &formfield.State{
						Label: "Custom Radio 1",
						Input: basicRadios[2],
					}),
					formfield.New(nil, &formfield.State{
						Label: "Custom Radio 2",
						Input: basicRadios[3],
					}),
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				elem.Heading2(vecty.Text("Disabled")),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					formfield.New(nil, &formfield.State{
						Label: "Disabled Radio 1",
						Input: radio.New(&base.Props{ID: "ex4a-radio1"},
							&radio.State{
								Name:     "ex4a",
								Checked:  true,
								Disabled: true,
							},
						),
					}),
					formfield.New(nil, &formfield.State{
						Label: "Disabled Radio 2",
						Input: radio.New(&base.Props{ID: "ex4a-radio2"},
							&radio.State{
								Name:     "ex4a",
								Disabled: true,
							},
						),
					}),
				),
			),
		),
	)
}
