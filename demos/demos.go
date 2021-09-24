package main

import (
	"path"

	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/demos/common"
	"github.com/vecty-material/material/ul"

	router "marwan.io/vecty-router"
)

// DemosCatalogView is our main page component.
type DemosCatalogView struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (c *DemosCatalogView) Render() vecty.ComponentOrHTML {
	vecty.SetTitle("Material Components Catalog")

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
			&ul.L{
				Root: vecty.Markup(
					vecty.Class("demo-catalog-list"),
				),
				Items: []vecty.ComponentOrHTML{
					&ul.Item{
						Primary:   vecty.Text("Button"),
						Secondary: vecty.Text("Raised and flat buttons"),
						OnClick:   makeRedirect("button"),
						Graphic:   renderGraphic("ic_button_24px.svg"),
					},
					// &ul.Item{
					// 	Primary:   vecty.Text("Card"),
					// 	Secondary: vecty.Text("Various card layout styles"),
					// 	Href:      makeHref("card"),
					// 	Graphic:   renderGraphic("ic_card_24px.svg"),
					// },
					&ul.Item{
						Primary:   vecty.Text("Checkbox"),
						Secondary: vecty.Text("Multi-selection controls"),
						OnClick:   makeRedirect("checkbox"),
						Graphic: renderGraphic(
							"ic_selection_control_24px.svg"),
					},
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Chips"),
					// 		Secondary: vecty.Text("Chips for actions, selection, and input "),
					// 		Href:      makeHref("chips"),
					// 		Graphic:   renderGraphic("ic_chips_24dp.svg"),
					// 	},
					// ),
					&ul.Item{
						Primary:   vecty.Text("Dialog"),
						Secondary: vecty.Text("Secondary text"),
						OnClick:   makeRedirect("dialog"),
						Graphic:   renderGraphic("ic_dialog_24px.svg"),
					},
					&ul.Item{
						Primary:   vecty.Text("Drawer"),
						Secondary: vecty.Text("Various drawer styles"),
						OnClick:   makeRedirect("drawer"),
						Graphic: renderGraphic(
							"ic_side_navigation_24px.svg"),
					},
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Elevation"),
					// 		Secondary: vecty.Text("Shadow for different elevations"),
					// 		Href:      makeHref("elevation"),
					// 		Graphic:   renderGraphic("ic_shadow_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Floating action button"),
					// 		Secondary: vecty.Text("The primary action in an application"),
					// 		Href:      makeHref("fab"),
					// 		Graphic:   renderGraphic("ic_button_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Grid list"),
					// 		Secondary: vecty.Text("2D grid layouts"),
					// 		Href:      makeHref("grid"),
					// 		Graphic:   renderGraphic("ic_card_24px.svg"),
					// 	},
					// ),
					&ul.Item{
						Primary:   vecty.Text("Icon toggle"),
						Secondary: vecty.Text("Toggling icon states"),
						OnClick:   makeRedirect("icontoggle"),
						Graphic:   renderGraphic("ic_component_24px.svg"),
					},
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Layout grid"),
					// 		Secondary: vecty.Text("Grid and gutter support"),
					// 		Href:      makeHref("layout"),
					// 		Graphic:   renderGraphic("ic_card_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Linear Progress"),
					// 		Secondary: vecty.Text("Fills from 0% to 100%, represented by bars"),
					// 		Href:      makeHref("linear"),
					// 		Graphic:   renderGraphic("ic_progress_activity.svg"),
					// 	},
					// ),
					&ul.Item{
						Primary:   vecty.Text("List"),
						Secondary: vecty.Text("Item layouts in lists"),
						OnClick:   makeRedirect("list"),
						Graphic:   renderGraphic("ic_list_24px.svg"),
					},
					&ul.Item{
						Primary:   vecty.Text("Menu"),
						Secondary: vecty.Text("Pop over menus"),
						OnClick:   makeRedirect("menu"),
						Graphic:   renderGraphic("ic_menu_24px.svg"),
					},
					&ul.Item{
						Primary:   vecty.Text("Radio buttons"),
						Secondary: vecty.Text("Single selection controls"),
						OnClick:   makeRedirect("radio"),
						Graphic:   renderGraphic("ic_radio_button_24px.svg"),
					},
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Ripple"),
					// 		Secondary: vecty.Text("Touch ripple"),
					// 		Href:      makeHref("ripple"),
					// 		Graphic:   renderGraphic("ic_ripple_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Select"),
					// 		Secondary: vecty.Text("Popover selection menus"),
					// 		Href:      makeHref("select"),
					// 		Graphic:   renderGraphic("ic_menu_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Slider"),
					// 		Secondary: vecty.Text("Range Controls"),
					// 		Href:      makeHref("slider"),
					// 		Graphic:   renderGraphic("slider.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Snackbar"),
					// 		Secondary: vecty.Text("Transient messages"),
					// 		Href:      makeHref("snackbar"),
					// 		Graphic:   renderGraphic("ic_toast_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Switch"),
					// 		Secondary: vecty.Text("On off switches"),
					// 		Href:      makeHref("switch"),
					// 		Graphic:   renderGraphic("ic_switch_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Tabs"),
					// 		Secondary: vecty.Text("Tabs with support for icon and text labels"),
					// 		Href:      makeHref("tabs"),
					// 		Graphic:   renderGraphic("ic_tabs_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Text field"),
					// 		Secondary: vecty.Text("Single and multiline text fields"),
					// 		Href:      makeHref("text"),
					// 		Graphic:   renderGraphic("ic_text_field_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Theme"),
					// 		Secondary: vecty.Text("Using primary and secondary colors"),
					// 		Href:      makeHref("theme"),
					// 		Graphic:   renderGraphic("ic_theme_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Toolbar"),
					// 		Secondary: vecty.Text("Header and footers"),
					// 		Href:      makeHref("toolbar"),
					// 		Graphic:   renderGraphic("ic_toolbar_24px.svg"),
					// 	},
					// ),
					// ul.NewItem(nil,
					// 	&ul.ItemState{
					// 		Primary:   vecty.Text("Typography"),
					// 		Secondary: vecty.Text("Type hierarchy"),
					// 		Href:      makeHref("typography"),
					// 		Graphic:   renderGraphic("ic_typography_24px.svg"),
					// 	},
					// ),
				}},
		)),
	)
}

func renderGraphic(filename string) vecty.ComponentOrHTML {
	return elem.Image(
		vecty.Markup(
			vecty.Class("demo-catalog-list-icon"),
			prop.Src("https://material-components-web.appspot.com/images/"+
				filename),
		),
	)
}

func makeHref(cName string) string {
	pathname := js.Global().Get("window").Get("location").Get("pathname").String()
	return path.Clean(pathname + "/" + cName)
}

func makeRedirect(cName string) func(i *ul.Item, e *vecty.Event) {
	return func(i *ul.Item, e *vecty.Event) {
		router.Redirect("/" + cName)
	}
}
