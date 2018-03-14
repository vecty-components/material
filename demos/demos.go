package main

import (
	"path"

	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/ul"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// demosCatalogView is our main page component.
type demosCatalogView struct {
	vecty.Core
}

type listItem struct {
	vecty.Core
	cssName     string
	imgPath     string
	description string
}

func main() {
	dcv := &demosCatalogView{}
	vecty.RenderBody(dcv)
}

// Render implements the vecty.Component interface.
func (c *demosCatalogView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		&common.ToolbarHeader{
			Title:      "Material Components Catalog",
			Navigation: common.NavRoot,
		},
		elem.Main(elem.Navigation(
			vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust")),
			ul.New(
				&base.Props{
					Markup: []vecty.Applyer{
						vecty.Class("demo-catalog-list"),
					},
				},
				&ul.State{Items: []*ul.Item{
					ul.NewItem(nil,
						&ul.ItemState{
							Primary:   "Button",
							Secondary: "Raised and flat buttons",
							Href:      makeHref("button"),
							Graphic:   renderGraphic("ic_button_24px.svg"),
							GraphicMarkup: []vecty.Applyer{
								vecty.Class("demo-catalog-list-icon"),
							},
						},
					),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Card",
					// 		Secondary: "Various card layout styles",
					// 		Href:      makeHref("card"),
					// 		Graphic:   renderGraphic(""),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					ul.NewItem(nil,
						&ul.ItemState{
							Primary:   "Checkbox",
							Secondary: "Multi-selection controls",
							Href:      makeHref("checkbox"),
							Graphic: renderGraphic(
								"ic_selection_control_24px.svg"),
							GraphicMarkup: []vecty.Applyer{
								vecty.Class("demo-catalog-list-icon"),
							},
						},
					),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Chips",
					// 		Secondary: "Chips for actions, selection, and input ",
					// 		Href:      makeHref("chips"),
					// 		Graphic:   renderGraphic("ic_chips_24dp.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					ul.NewItem(nil,
						&ul.ItemState{
							Primary:   "Dialog",
							Secondary: "Secondary text",
							Href:      makeHref("dialog"),
							Graphic:   renderGraphic("ic_dialog_24px.svg"),
							GraphicMarkup: []vecty.Applyer{
								vecty.Class("demo-catalog-list-icon"),
							},
						},
					),
					ul.NewItem(nil,
						&ul.ItemState{
							Primary:   "Drawer",
							Secondary: "Various drawer styles",
							Href:      makeHref("drawer"),
							Graphic: renderGraphic(
								"ic_side_navigation_24px.svg"),
							GraphicMarkup: []vecty.Applyer{
								vecty.Class("demo-catalog-list-icon"),
							},
						},
					),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Elevation",
					// 		Secondary: "Shadow for different elevations",
					// 		Href:      makeHref("elevation"),
					// 		Graphic:   renderGraphic("ic_shadow_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Floating action button",
					// 		Secondary: "The primary action in an application",
					// 		Href:      makeHref("fab"),
					// 		Graphic:   renderGraphic("ic_button_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Grid list",
					// 		Secondary: "2D grid layouts",
					// 		Href:      makeHref("grid"),
					// 		Graphic:   renderGraphic("ic_card_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Icon toggle",
					// 		Secondary: "Toggling icon states",
					// 		Href:      makeHref("icon"),
					// 		Graphic:   renderGraphic("ic_component_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Layout grid",
					// 		Secondary: "Grid and gutter support",
					// 		Href:      makeHref("layout"),
					// 		Graphic:   renderGraphic("ic_card_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Linear Progress",
					// 		Secondary: "Fills from 0% to 100%, represented by bars",
					// 		Href:      makeHref("linear"),
					// 		Graphic:   renderGraphic("ic_progress_activity.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "List",
					// 		Secondary: "Item layouts in lists",
					// 		Href:      makeHref("list"),
					// 		Graphic:   renderGraphic("ic_list_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Menu",
					// 		Secondary: "Pop over menus",
					// 		Href:      makeHref("menu"),
					// 		Graphic:   renderGraphic("ic_menu_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Radio buttons",
					// 		Secondary: "Single selection controls",
					// 		Href:      makeHref("radio"),
					// 		Graphic:   renderGraphic("ic_radio_button_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Ripple",
					// 		Secondary: "Touch ripple",
					// 		Href:      makeHref("ripple"),
					// 		Graphic:   renderGraphic("ic_ripple_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Select",
					// 		Secondary: "Popover selection menus",
					// 		Href:      makeHref("select"),
					// 		Graphic:   renderGraphic("ic_menu_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Slider",
					// 		Secondary: "Range Controls",
					// 		Href:      makeHref("slider"),
					// 		Graphic:   renderGraphic("slider.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Snackbar",
					// 		Secondary: "Transient messages",
					// 		Href:      makeHref("snackbar"),
					// 		Graphic:   renderGraphic("ic_toast_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Switch",
					// 		Secondary: "On off switches",
					// 		Href:      makeHref("switch"),
					// 		Graphic:   renderGraphic("ic_switch_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Tabs",
					// 		Secondary: "Tabs with support for icon and text labels",
					// 		Href:      makeHref("tabs"),
					// 		Graphic:   renderGraphic("ic_tabs_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Text field",
					// 		Secondary: "Single and multiline text fields",
					// 		Href:      makeHref("text"),
					// 		Graphic:   renderGraphic("ic_text_field_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Theme",
					// 		Secondary: "Using primary and secondary colors",
					// 		Href:      makeHref("theme"),
					// 		Graphic:   renderGraphic("ic_theme_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Toolbar",
					// 		Secondary: "Header and footers",
					// 		Href:      makeHref("toolbar"),
					// 		Graphic:   renderGraphic("ic_toolbar_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   "Typography",
					// 		Secondary: "Type hierarchy",
					// 		Href:      makeHref("typography"),
					// 		Graphic:   renderGraphic("ic_typography_24px.svg"),
					// 		GraphicMarkup: []vecty.Applyer{
					// 			vecty.Class("demo-catalog-list-icon"),
					// 		},
					// 	},
					// ),
				}},
			))),
	)
}

func renderGraphic(
	filename string) vecty.ComponentOrHTML {
	return elem.Image(
		vecty.Markup(
			prop.Src("https://material-components-web.appspot.com/images/" +
				filename),
		),
	)
}

func makeHref(cName string) string {
	pathname := js.Global.Get("window").Get("location").Get("pathname").String()
	return path.Clean(pathname + "/" + cName)
}
