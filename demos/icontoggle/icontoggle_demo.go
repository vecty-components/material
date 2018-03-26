package main

import (
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/icon"
	"agamigo.io/vecty-material/icontoggle"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// icontoggleDemoView is our demo page component.
type icontoggleDemoView struct {
	vecty.Core
	favStatus bool `vecty:"prop"`
}

func main() {
	vecty.RenderBody(&icontoggleDemoView{})
}

// Render implements the vecty.Component interface.
func (c *icontoggleDemoView) Render() vecty.ComponentOrHTML {
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
					icontoggle.New(nil, &icontoggle.State{
						OffLabel: "Add to Favorites",
						OffIcon: icon.New(nil, &icon.State{
							Name: "favorite_border",
						}),
						OnLabel: "Remove From Favorites",
						OnIcon: icon.New(nil, &icon.State{
							Name: "favorite",
						}),
					}),
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("example")),
				elem.Div(
					vecty.Markup(vecty.Class("toggle-example")),
					elem.Heading2(vecty.Text("Using Material Icons")),
					elem.Div(
						vecty.Markup(vecty.Class("demo-wrapper")),
						icontoggle.New(nil, &icontoggle.State{
							OffLabel: "Add to Favorites",
							OffIcon: icon.New(nil, &icon.State{
								Name: "favorite_border",
							}),
							OnLabel: "Remove From Favorites",
							OnIcon: icon.New(nil, &icon.State{
								Name: "favorite",
							}),
							ChangeHandler: func(it *icontoggle.IT,
								e *vecty.Event) {
								c.favStatus = it.On
								vecty.Rerender(c)
							},
						}),
					),
					elem.Paragraph(
						vecty.Text("Favorited? "),
						elem.Span(
							vecty.If(c.favStatus, vecty.Text("yes")),
							vecty.If(!c.favStatus, vecty.Text("no")),
						),
					),
				),
				elem.Div(
					vecty.Markup(vecty.Class("toggle-example")),
					elem.Heading2(vecty.Text("Using Font Awesome")),
					elem.Div(
						vecty.Markup(vecty.Class("demo-wrapper")),
						icontoggle.New(nil, &icontoggle.State{
							On:       true,
							OffLabel: "Star this item",
							OffIcon: icon.New(nil, &icon.State{
								ClassOverride: []string{"fa", "fa-star-o"},
							}),
							OnLabel: "Unstar this item",
							OnIcon: icon.New(nil, &icon.State{
								ClassOverride: []string{"fa", "fa-star"},
							}),
						}),
					),
				),
				elem.Div(
					vecty.Markup(vecty.Class("toggle-example")),
					elem.Heading2(vecty.Text("Disabled Icons")),
					elem.Div(
						vecty.Markup(vecty.Class("demo-wrapper")),
						icontoggle.New(nil, &icontoggle.State{
							Disabled: true,
							OffLabel: "Add to Favorites",
							OffIcon: icon.New(nil, &icon.State{
								Name: "favorite_border",
							}),
							OnLabel: "Remove From Favorites",
							OnIcon: icon.New(nil, &icon.State{
								Name: "favorite",
							}),
						}),
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
								icontoggle.New(nil, &icontoggle.State{
									OffLabel: "Add to Favorites",
									OffIcon: icon.New(nil, &icon.State{
										Name: "favorite_border",
									}),
									OnLabel: "Remove From Favorites",
									OnIcon: icon.New(nil, &icon.State{
										Name: "favorite",
									}),
								}),
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
								icontoggle.New(nil, &icontoggle.State{
									OffLabel: "Add to Favorites",
									OffIcon: icon.New(nil, &icon.State{
										Name: "favorite_border",
									}),
									OnLabel: "Remove From Favorites",
									OnIcon: icon.New(nil, &icon.State{
										Name: "favorite",
									}),
								}),
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
								icontoggle.New(nil, &icontoggle.State{
									OffLabel: "Add to Favorites",
									OffIcon: icon.New(nil, &icon.State{
										Name: "favorite_border",
									}),
									OnLabel: "Remove From Favorites",
									OnIcon: icon.New(nil, &icon.State{
										Name: "favorite",
									}),
								}),
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
