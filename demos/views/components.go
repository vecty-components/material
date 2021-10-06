package views

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/demos/components"
	router "marwan.io/vecty-router"
)

type ComponentImageList struct {
	vecty.Core
	images map[string]string
	list   []components.DemoLink
}

func NewComponentImageList(list []components.DemoLink) *ComponentImageList {
	return &ComponentImageList{
		images: make(map[string]string),
		list:   list,
	}
}

func (cl *ComponentImageList) Render() vecty.ComponentOrHTML {

	children := []vecty.MarkupOrChild{
		vecty.Markup(
			prop.ID("catalog-image-list"),
			vecty.Class(
				"mdc-image-list", "standard-image-list",
				"mdc-top-app-bar--fixed-adjust",
			),
		),
	}

	for _, item := range cl.list {
		children = append(
			children, cl.renderListItem(item.Name, item.Image, item.Url),
		)
	}

	return elem.UnorderedList(
		children...,
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
		base.RichLink(url,
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
