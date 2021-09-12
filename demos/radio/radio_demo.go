package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/lithammer/dedent"
	"github.com/vecty-material/material"
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
	material.SetViewport()

	vecty.SetTitle("Radio Button - Material Components Catalog")
	vecty.AddStylesheet("https://material-components-web.appspot.com/assets/radio.css")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto+Mono")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:300,400,500")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.6.3/css/font-awesome.css")

	material.AddIcon("https://material-components-web.appspot.com/images/logo_components_color_2x_web_48dp.png")
	material.AddCSS(dedent.Dedent(`
		.example {
			margin: 24px;
			padding: 24px;
		}
	`))

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
