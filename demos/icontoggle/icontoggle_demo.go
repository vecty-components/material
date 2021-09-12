package main

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/demos/common"
	"github.com/vecty-material/material/icon"
	"github.com/vecty-material/material/icontoggle"
)

// icontoggleDemoView is our demo page component.
type icontoggleDemoView struct {
	vecty.Core
	favStatus bool `vecty:"prop"`
}

type favorite struct {
	vecty.Core
	status string
}

func main() {
	base.SetViewport()

	vecty.SetTitle("Icon Toggle - Material Components Demo")
	vecty.AddStylesheet("https://material-components-web.appspot.com/assets/icon-toggle.css")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto+Mono")
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:300,400,500")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.6.3/css/font-awesome.css")

	base.AddIcon("https://material-components-web.appspot.com/images/logo_components_color_2x_web_48dp.png")

	time.Sleep(1 * time.Second)

	vecty.RenderBody(&icontoggleDemoView{})
}

// Render implements the vecty.Component interface.
func (c *icontoggleDemoView) Render() vecty.ComponentOrHTML {
	favStatus := &favorite{status: "no"}
	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		&common.ToolbarHeader{
			Title:      "Icon Toggle",
			Navigation: common.NavBack,
		},
		elem.Main(
			elem.Div(vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust"))),
			elem.Section(
				vecty.Markup(vecty.Class("hero")),
				elem.Div(
					vecty.Markup(vecty.Class("demo-wrapper")),
					&icontoggle.IT{
						OffLabel: "Add to Favorites",
						OffIcon: &icon.I{
							Name: "favorite_border",
						},
						OnLabel: "Remove From Favorites",
						OnIcon: &icon.I{
							Name: "favorite",
						},
					},
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				elem.Div(
					vecty.Markup(vecty.Class("toggle-example")),
					elem.Heading2(vecty.Text("Using Material Icons")),
					elem.Div(
						vecty.Markup(vecty.Class("demo-wrapper")),
						&icontoggle.IT{
							OffLabel: "Add to Favorites",
							OffIcon: &icon.I{
								Name: "favorite_border",
							},
							OnLabel: "Remove From Favorites",
							OnIcon: &icon.I{
								Name: "favorite",
							},
							ChangeHandler: func(it *icontoggle.IT,
								e *vecty.Event) {
								if it.On {
									favStatus.status = "yes"
								} else {
									favStatus.status = "no"
								}
								vecty.Rerender(favStatus)
							},
						},
					),
					favStatus,
				),
				elem.Div(
					vecty.Markup(vecty.Class("toggle-example")),
					elem.Heading2(vecty.Text("Using Font Awesome")),
					elem.Div(
						vecty.Markup(vecty.Class("demo-wrapper")),
						&icontoggle.IT{
							On:       true,
							OffLabel: "Star this item",
							OffIcon: &icon.I{
								ClassOverride: []string{"fa", "fa-star-o"},
							},
							OnLabel: "Unstar this item",
							OnIcon: &icon.I{
								ClassOverride: []string{"fa", "fa-star"},
							},
						},
					),
				),
				elem.Div(
					vecty.Markup(vecty.Class("toggle-example")),
					elem.Heading2(vecty.Text("Disabled Icons")),
					elem.Div(
						vecty.Markup(vecty.Class("demo-wrapper")),
						&icontoggle.IT{
							Disabled: true,
							OffLabel: "Add to Favorites",
							OffIcon: &icon.I{
								Name: "favorite_border",
							},
							OnLabel: "Remove From Favorites",
							OnIcon: &icon.I{
								Name: "favorite",
							},
						},
					),
				),
				elem.Div(
					vecty.Markup(vecty.Class("toggle-example")),
					elem.Heading2(vecty.Text("Additional Color Combinations")),
					elem.Div(vecty.Markup(vecty.Class("demo-color-combos")),
						elem.Div(
							vecty.Markup(
								prop.ID("light-on-bg"),
								vecty.Class("demo-color-combo"),
							),
							elem.Div(
								&icontoggle.IT{
									OffLabel: "Add to Favorites",
									OffIcon: &icon.I{
										Name: "favorite_border",
									},
									OnLabel: "Remove From Favorites",
									OnIcon: &icon.I{
										Name: "favorite",
									},
								},
							),
							elem.Div(vecty.Markup(
								vecty.Class(
									"mdc-theme--text-primary-on-primary")),
								vecty.Text("Light icon on background"),
							),
						),
						elem.Div(
							vecty.Markup(
								prop.ID("dark-on-bg"),
								vecty.Class("demo-color-combo"),
							),
							elem.Div(
								vecty.Markup(
									vecty.Class("mdc-theme--primary"),
								),
								&icontoggle.IT{
									OffLabel: "Add to Favorites",
									OffIcon: &icon.I{
										Name: "favorite_border",
									},
									OnLabel: "Remove From Favorites",
									OnIcon: &icon.I{
										Name: "favorite",
									},
								},
							),
							elem.Div(
								vecty.Text("Dark icon on background"),
							),
						),
						elem.Div(
							vecty.Markup(
								prop.ID("custom-color-combo"),
								vecty.Class("demo-color-combo"),
							),
							elem.Div(
								vecty.Markup(
									vecty.Class("mdc-theme--primary"),
								),
								&icontoggle.IT{
									OffLabel: "Add to Favorites",
									OffIcon: &icon.I{
										Name: "favorite_border",
									},
									OnLabel: "Remove From Favorites",
									OnIcon: &icon.I{
										Name: "favorite",
									},
								},
							),
							elem.Div(
								vecty.Text("Custom color"),
							),
						),
					),
				),
			),
		),
	)
}

func (c *favorite) Render() vecty.ComponentOrHTML {
	return elem.Paragraph(
		vecty.Text("Favorited? "),
		elem.Span(
			vecty.Text(c.status),
		),
	)
}
