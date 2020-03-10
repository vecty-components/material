package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/vecty-material/material/base/applyer"
	"github.com/vecty-material/material/demos/common"
	"github.com/vecty-material/material/formfield"
	"github.com/vecty-material/material/radio"
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
		&radio.R{
			Root: vecty.Markup(
				applyer.CSSOnly(),
			),
			Input: vecty.Markup(
				prop.ID("ex1-default-radio1"),
			),
			Name:    "ex1-default",
			Checked: true,
		},
		&radio.R{
			Root: vecty.Markup(
				applyer.CSSOnly(),
			),
			Input: vecty.Markup(
				prop.ID("ex1-default-radio2"),
			),
			Name: "ex1-default",
		},
		&radio.R{
			Root: vecty.Markup(
				vecty.Class("demo-radio--custom"),
				applyer.CSSOnly(),
			),
			Input: vecty.Markup(
				prop.ID("ex1-custom-radio1"),
			),
			Name:    "ex1-custom",
			Checked: true,
		},
		&radio.R{
			Root: vecty.Markup(
				vecty.Class("demo-radio--custom"),
				applyer.CSSOnly(),
			),
			Input: vecty.Markup(
				prop.ID("ex1-custom-radio2"),
			),
			Name: "ex1-custom",
		},
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
				&radio.R{Name: "hero"},
				&radio.R{Name: "hero", Checked: true},
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				elem.Heading2(vecty.Text("With JavaScript")),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					&formfield.FF{
						Label: "Default Radio 1",
						Input: &radio.R{
							Input: vecty.Markup(
								prop.ID("ex0-default-radio1"),
							),
							Name:    "ex0-default",
							Checked: true,
						},
					},
					&formfield.FF{
						Label: "Default Radio 2",
						Input: &radio.R{
							Input: vecty.Markup(
								prop.ID("ex0-default-radio2"),
							),
							Name: "ex0-default"},
					},
				),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					&formfield.FF{
						Label: "Custom Radio 1",
						Input: &radio.R{
							Root: vecty.Markup(
								vecty.Class("demo-radio--custom"),
							),
							Input: vecty.Markup(
								prop.ID("ex0-custom-radio1"),
							),
							Name:    "ex0-custom",
							Checked: true,
						},
					},
					&formfield.FF{
						Label: "Custom Radio 2",
						Input: &radio.R{
							Root: vecty.Markup(
								vecty.Class("demo-radio--custom"),
							),
							Input: vecty.Markup(
								prop.ID("ex0-custom-radio2"),
							),
							Name: "ex0-custom",
						},
					},
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				elem.Heading2(vecty.Text("CSS Only")),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					&formfield.FF{
						Label: "Default Radio 1",
						Input: basicRadios[0],
					},
					&formfield.FF{
						Label: "Default Radio 2",
						Input: basicRadios[1],
					},
				),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					&formfield.FF{
						Label: "Custom Radio 1",
						Input: basicRadios[2],
					},
					&formfield.FF{
						Label: "Custom Radio 2",
						Input: basicRadios[3],
					},
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				elem.Heading2(vecty.Text("Disabled")),
				elem.Div(vecty.Markup(vecty.Class("demo-radio-row")),
					&formfield.FF{
						Label: "Disabled Radio 1",
						Input: &radio.R{
							Input: vecty.Markup(
								prop.ID("ex4a-radio1"),
							),
							Name:     "ex4a",
							Checked:  true,
							Disabled: true,
						},
					},
					&formfield.FF{
						Label: "Disabled Radio 2",
						Input: &radio.R{
							Input: vecty.Markup(
								prop.ID("ex4a-radio2"),
							),
							Name:     "ex4a",
							Disabled: true,
						},
					},
				),
			),
		),
	)
}
