package main

import (
	"path"
	"strings"

	"agamigo.io/vecty-material/demos/common"
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
		elem.Main(
			elem.Navigation(
				vecty.Markup(
					vecty.Class("mdc-toolbar-fixed-adjust"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Attribute("role", "list"),
						vecty.Class("demo-catalog-list"),
						vecty.Class("mdc-list"),
						vecty.Class("mdc-list--two-line"),
					),
				),
				&listItem{
					cssName:     "button",
					description: "Raised and flat buttons",
					imgPath: "https://material-components-web.appspot.com/" +
						"images/ic_button_24px.svg",
				},
				// &listItem{
				// 	cssName:     "card",
				// 	description: "Various card layout styles",
				// },
				&listItem{
					cssName:     "checkbox",
					description: "Multi-selection controls",
					imgPath: "https://material-components-web.appspot.com/" +
						"images/ic_selection_control_24px.svg",
				},
				// &listItem{
				// 	cssName:     "chips",
				// 	imgPath:     "/images/ic_chips_24dp.svg",
				// 	description: "Chips for actions, selection, and input ",
				// },
				&listItem{
					cssName:     "dialog",
					description: "Secondary text",
					imgPath: "https://material-components-web.appspot.com/" +
						"images/ic_dialog_24px.svg",
				},
				// &listItem{
				// 	cssName:     "drawer",
				// 	imgPath:     "/images/ic_side_navigation_24px.svg",
				// 	description: "Various drawer styles",
				// },
				// &listItem{
				// 	cssName:     "elevation",
				// 	imgPath:     "/images/ic_shadow_24px.svg",
				// 	description: "Shadow for different elevations",
				// },
				// &listItem{
				// 	cssName:     "fab",
				// 	imgPath:     "/images/ic_button_24px.svg",
				// 	description: "The primary action in an application",
				// },
				// &listItem{
				// 	cssName:     "grid-list",
				// 	imgPath:     "/images/ic_card_24px.svg",
				// 	description: "2D grid layouts",
				// },
				// &listItem{
				// 	cssName:     "icon-toggle",
				// 	imgPath:     "/images/ic_component_24px.svg",
				// 	description: "Toggling icon states",
				// },
				// &listItem{
				// 	cssName:     "layout-grid",
				// 	imgPath:     "/images/ic_card_24px.svg",
				// 	description: "Grid and gutter support",
				// },
				// &listItem{
				// 	cssName:     "linear-progress",
				// 	imgPath:     "/images/ic_progress_activity.svg",
				// 	description: "Fills from 0% to 100%, represented by bars",
				// },
				// &listItem{
				// 	cssName:     "list",
				// 	description: "Item layouts in lists",
				// },
				// &listItem{
				// 	cssName:     "menu",
				// 	description: "Pop over menus",
				// },
				// &listItem{
				// 	cssName:     "radio",
				// 	imgPath:     "/images/ic_radio_button_24px.svg",
				// 	description: "Single selection controls",
				// },
				// &listItem{
				// 	cssName:     "ripple",
				// 	description: "Touch ripple",
				// },
				// &listItem{
				// 	cssName:     "select",
				// 	imgPath:     "/images/ic_menu_24px.svg",
				// 	description: "Popover selection menus",
				// },
				// &listItem{
				// 	cssName:     "slider",
				// 	imgPath:     "/images/slider.svg",
				// 	description: "Range Controls",
				// },
				// &listItem{
				// 	cssName:     "snackbar",
				// 	imgPath:     "/images/ic_toast_24px.svg",
				// 	description: "Transient messages",
				// },
				// &listItem{
				// 	cssName:     "switch",
				// 	description: "On off switches",
				// },
				// &listItem{
				// 	cssName:     "tabs",
				// 	description: "Tabs with support for icon and text labels",
				// },
				// &listItem{
				// 	cssName:     "text-field",
				// 	imgPath:     "/images/ic_text_field_24px.svg",
				// 	description: "Single and multiline text fields",
				// },
				// &listItem{
				// 	cssName:     "theme",
				// 	description: "Using primary and secondary colors",
				// },
				// &listItem{
				// 	cssName:     "toolbar",
				// 	description: "Header and footers",
				// },
				// &listItem{
				// 	cssName:     "typography",
				// 	description: "Type hierarchy",
				// },
			),
		),
	)
}

// Render implements the vecty.Component interface.
func (c *listItem) Render() vecty.ComponentOrHTML {
	pathname := js.Global.Get("window").Get("location").Get("pathname").String()
	return elem.Anchor(
		vecty.Markup(
			vecty.Attribute("role", "listitem"),
			vecty.Class("mdc-list-item"),
			prop.Href(path.Clean(pathname+"/"+c.cssName)),
		),
		elem.Span(
			vecty.Markup(
				vecty.Class("demo-catalog-list-icon"),
				vecty.Class("mdc-list-item__graphic"),
			),
			elem.Image(
				vecty.Markup(
					prop.Src(c.imgPath),
				),
			),
		),
		elem.Span(
			vecty.Text(c.title()),
			elem.Span(
				vecty.Markup(
					vecty.Class("mdc-list-item__secondary-text"),
				),
				vecty.Text(c.description),
			),
		),
	)
}

func (c *listItem) title() string {
	words := strings.Split(c.cssName, "-")
	var t string
	for i, v := range words {
		if i == 0 {
			t = strings.Title(v)
			continue
		}
		t = t + " " + v
	}
	return t
}
