package views

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-material/material/base"
	router "marwan.io/vecty-router"
)

type ComponentImageList struct {
	vecty.Core
	images map[string]string
}

func NewComponentImageList() *ComponentImageList {
	return &ComponentImageList{
		images: make(map[string]string),
	}
}

func (cl *ComponentImageList) Render() vecty.ComponentOrHTML {
	vecty.AddStylesheet("/assets/styles/ImageListCatalog.css")

	buttonImg := "/assets/images/buttons_180px.svg"
	// cardsImg := "/assets/images/cards_180px.svg"
	// checkboxImg := "/assets/images/checkbox_180px.svg"
	// chipsImg := "/assets/images/chips_180px.svg"
	// dataTableImg := "/assets/images/data_table_180px.svg"
	// dialogImg := "/assets/images/dialog_180px.svg"
	// drawerImg := "/assets/images/drawer_180px.svg"
	// elevationImg := "/assets/images/elevation_180px.svg"
	// fabImg := "/assets/images/floating_action_button_180px.svg"
	// iconButtonImg := "/assets/images/icon_button_180px.svg"
	// imageListImg := "/assets/images/image_list_180px.svg"
	// inputImg := "/assets/images/form_field_180px.svg"
	// layoutGridImg := "/assets/images/layout_grid_180px.svg"
	// listImg := "/assets/images/list_180px.svg"
	// linearProgressImg := "/assets/images/linear_progress_180px.svg"
	menuImg := "/assets/images/menu_180px.svg"
	// radioImg := "/assets/images/radio_180px.svg"
	// rippleImg := "/assets/images/ripple_180px.svg"
	// sliderImg := "/assets/images/slider_180px.svg"
	// snackbarImg := "/assets/images/snackbar_180px.svg"
	// switchImg := "/assets/images/switch_180px.svg"
	// tabsImg := "/assets/images/tabs_180px.svg"
	// themeImg := "/assets/images/ic_theme_24px.svg"
	// topAppBarImg := "/assets/images/top_app_bar_180px.svg"
	// typographyImg := "/assets/images/fonts_180px.svg"

	return elem.Div(
		elem.UnorderedList(
			vecty.Markup(
				prop.ID("catalog-image-list"),
				vecty.Class(
					"mdc-image-list", "standard-image-list",
					"mdc-top-app-bar--fixed-adjust",
				),
			),
			cl.renderListItem("Button", buttonImg, "button"),
			cl.renderListItem("Menu", menuImg, "menu"),
		),
	)
}

func (cl *ComponentImageList) renderListItem(
	title, imageSource, url string) vecty.ComponentOrHTML {

	if _, ok := cl.images[imageSource]; !ok {
		cl.images[imageSource] = ""
		go func() {

			r, _ := http.Get(imageSource)
			svg, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()
			source := string(svg)

			/* TODO: fix this to manipulate html */
			source = strings.ReplaceAll(
				source, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", "",
			)
			source = strings.ReplaceAll(
				source, "<svg width=\"180px\" height=\"180px\"", "<svg",
			)

			cl.images[imageSource] = source
			vecty.Rerender(cl)
		}()
	}

	return elem.ListItem(
		vecty.Markup(
			vecty.Class("catalog-image-list-item", "mdc-image-list__item"),
		),
		base.RichLink("/"+url,
			[]vecty.ComponentOrHTML{
				elem.Div(
					vecty.Markup(
						vecty.Class(
							"catalog-image-list-item-container",
							"mdc-image-list__image-aspect-container",
							"mdc-ripple-surface",
						),
					),
					elem.Div(
						vecty.Markup(
							vecty.Class("mdc-image-list__image"),
							vecty.UnsafeHTML(cl.images[imageSource]),
						),
					),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("mdc-image-list__supporting"),
					),
					elem.Span(
						vecty.Markup(
							vecty.Class("catalog-image-list-label", "mdc-image-list__label"),
						),
						vecty.Text(title),
					),
				),
			}, router.LinkOptions{Class: "catalog-image-link"},
		),
	)
}
