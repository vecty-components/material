package main

import (
	"strconv"

	"syscall/js"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/vecty-material/material/checkbox"
	"github.com/vecty-material/material/demos/common"
	"github.com/vecty-material/material/formfield"
	"github.com/vecty-material/material/icon"
	"github.com/vecty-material/material/ripple"
	"github.com/vecty-material/material/ul"
)

const MDCImagesURL = "https://material-components-web.appspot.com/images/"

type demoSection struct {
	vecty.Core
	heading       string
	lists         []*demoList
	groups        []*ul.Group
	groupHeadings []string
}

type demoList struct {
	vecty.Core
	heading string
	markup  []vecty.Applyer
}

// listDemoView is our demo page component.
type listDemoView struct {
	vecty.Core
}

func main() {
	vecty.RenderBody(&listDemoView{})
}

// Render implements the vecty.Component interface.
func (c *listDemoView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(vecty.Class("mdc-typography")),
		&common.ToolbarHeader{
			Title:      "List",
			Navigation: common.NavBack,
		},
		elem.Main(
			elem.Div(vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust"))),
			&demoSection{
				heading: "hero",
				lists: []*demoList{
					&demoList{heading: "hero"}},
			},
			elem.Section(
				vecty.Markup(vecty.Class("preamble", "mdc-typography--body1")),
				elem.Aside(elem.Paragraph(
					elem.Emphasis(vecty.Text("NOTE:")),
					vecty.Text(" For the purposes of this demo, we've set a "+
						"max-width of 600px on all "),
					elem.Code(vecty.Text("mdc-list")),
					vecty.Text(" elements, and surrounded them by a 1px "+
						"border. This is not included in the base css, which "+
						"has the list take up as much width as possible "+
						"(since it's a block element)."),
				)),

				&formfield.FF{
					Label: "Toggle RTL",
					Input: &checkbox.CB{
						Input: vecty.Markup(
							prop.ID("toggle-rtl"),
						),
						OnChange: func(thisCB *checkbox.CB,
							e *vecty.Event) {
							w := js.Global().Get("window")
							d := w.Get("document")
							dw := d.Call("getElementById",
								"demo-wrapper")
							if dw.Call("getAttribute",
								"dir").String() == "rtl" {
								dw.Call("setAttribute", "dir", "ltr")
								return
							}
							dw.Call("setAttribute", "dir", "rtl")
						},
					},
				},
			),
			elem.Div(vecty.Markup(prop.ID("demo-wrapper")),
				&demoSection{
					heading: "Custom Colors",
					groups: []*ul.Group{
						&ul.Group{
							Root: vecty.Markup(
								vecty.Class("demo-list-group--custom"),
							),
						},
					},
					groupHeadings: []string{
						"Example - Two-Line Lists, Avatars, " +
							"Metadata, Inset Dividers"},
					lists: []*demoList{
						&demoList{heading: "Folders",
							markup: []vecty.Applyer{
								vecty.Class("demo-list--custom")}},
						&demoList{heading: "Files",
							markup: []vecty.Applyer{
								vecty.Class("demo-list--custom")}},
					},
				},
				&demoSection{
					heading: "Single-Line List",
					lists: []*demoList{
						&demoList{heading: "Text only, non-interactive " +
							"(no states)"},
						&demoList{heading: "Text only (dense)"},
						&demoList{heading: "Graphic"},
						&demoList{heading: "Graphic (dense)"},
						&demoList{heading: "Graphic Example - Icon with Text"},
						&demoList{heading: "Leading Checkbox"},
						&demoList{heading: "Avatar List"},
						&demoList{heading: "Avatar List (dense)"},
						&demoList{heading: "Example - Avatar with Text"},
						&demoList{heading: "Metadata"},
						&demoList{heading: "Metadata (Dense)"},
						&demoList{heading: "Trailing Checkbox"},
						&demoList{heading: "Avatar + Metadata"},
						&demoList{heading: "Avatar + Metadata (Dense)"},
						&demoList{heading: "Example - Avatar with Text " +
							"and Icon"},
					},
				},
				&demoSection{
					heading: "Two-Line List",
					lists: []*demoList{
						&demoList{heading: "Text-Only"},
						&demoList{heading: "Text-Only (Dense)"},
						&demoList{heading: "Graphic"},
						&demoList{heading: "Graphic (Dense)"},
						&demoList{heading: "Avatar List"},
						&demoList{heading: "Avatar List (dense)"},
						&demoList{heading: "Metadata"},
						&demoList{heading: "Metadata (Dense)"},
						&demoList{heading: "Example - Two-line Avatar + " +
							"Text + Icon"},
						&demoList{heading: "Lists w/ Ellipsis"},
					},
				},
				&demoSection{
					heading: "List Dividers",
					lists: []*demoList{
						&demoList{heading: "Full-Width Dividers"},
						&demoList{heading: "Inset Dividers"},
					},
				},
				&demoSection{
					heading: "List Groups",
					groups: []*ul.Group{
						&ul.Group{},
						&ul.Group{},
					},
					groupHeadings: []string{
						"Basic Usage",
						"Example - Two-Line Lists, Avatars, Metadata, " +
							"Inset Dividers",
					},
					lists: []*demoList{
						&demoList{heading: "List 1"},
						&demoList{heading: "List 2"},
						&demoList{heading: "Folders"},
						&demoList{heading: "Files"},
					},
				},
				&demoSection{
					heading: "Interactive Lists (with ink ripple)",
					lists: []*demoList{
						&demoList{heading: "Example - Interactive List"},
					},
				},
			),
		),
	)
}

func (c *demoSection) Render() vecty.ComponentOrHTML {
	h := c.heading
	var lists vecty.List
	if c.groups != nil && len(c.groups) > 0 {
		if len(c.groups)*2 != len(c.lists) {
			panic("There should be two lists per list group.")
		}
		for i, group := range c.groups {
			lists = append(lists, elem.Heading3(vecty.Text(c.groupHeadings[i])))
			c.groups[i].Lists = append(group.Lists, c.lists[i*2].make(h))
			c.groups[i].Lists = append(group.Lists, c.lists[(i*2)+1].make(h))
			lists = append(lists, c.groups[i])
		}
	} else {
		for _, list := range c.lists {
			if h != "hero" {
				lists = append(lists, elem.Heading3(vecty.Text(list.heading)))
			}
			if h == "Single-Line List" && list.heading == "Graphic" {
				lists = append(lists, elem.Aside(elem.Paragraph(elem.Emphasis(
					vecty.Text("Note: The grey background is styled "+
						"using demo placeholder styles")))))
			}
			lists = append(lists, list.make(h))
		}
	}
	if c.heading == "hero" {
		return elem.Section(vecty.Markup(
			vecty.Class("hero")),
			lists)
	}
	return elem.Section(
		elem.Heading2(vecty.Text(c.heading)),
		elem.Section(lists),
	)
}

func (c *demoList) make(section string) *ul.L {
	c.markup = append(c.markup, vecty.Class("demo-list"))
	newL := &ul.L{}
	newL.GroupSubheader = c.heading
	switch c.heading {
	case "Text only, non-interactive (no states)":
		c.makeSingleLineItems(newL)
		newL.NonInteractive = true
	case "Text-Only":
		c.makeTwoLineItems(newL)
	case "Text only (dense)":
		c.makeSingleLineItems(newL)
		newL.Dense = true
	case "Text-Only (Dense)":
		c.makeTwoLineItems(newL)
		newL.Dense = true
	case "Graphic":
		if section == "Two-Line List" {
			c.makeTwoLineItems(newL)
		} else {
			c.makeSingleLineItems(newL)
		}
		c.withIconPlaceholders(newL)
	case "Graphic (dense)":
		c.makeSingleLineItems(newL)
		c.withIconPlaceholders(newL)
		newL.Dense = true
	case "Graphic (Dense)":
		c.makeTwoLineItems(newL)
		c.withIconPlaceholders(newL)
		newL.Dense = true
	case "Graphic Example - Icon with Text":
		c.makeGraphicExampleItems(newL, false)
	case "Leading Checkbox":
		c.makeCheckboxItems(newL, true)
		c.markup = append(c.markup, prop.ID("leading-checkbox-list"))
	case "Trailing Checkbox":
		c.makeCheckboxItems(newL, false)
		c.markup = append(c.markup, prop.ID("trailing-checkbox-list"))
	case "Avatar List":
		if section == "Two-Line List" {
			c.makeTwoLineItems(newL)
		} else {
			c.makeSingleLineItems(newL)
		}
		c.withAvatars(newL)
		c.withIconPlaceholders(newL)
	case "Avatar List (dense)":
		if section == "Two-Line List" {
			c.makeTwoLineItems(newL)
		} else {
			c.makeSingleLineItems(newL)
		}
		c.withAvatars(newL)
		c.withIconPlaceholders(newL)
		newL.Dense = true
	case "Example - Avatar with Text":
		c.makeAvatarWTextItems(newL)
		c.withAvatars(newL)
	case "Metadata":
		if section == "Two-Line List" {
			c.makeTwoLineItems(newL)
		} else {
			c.makeSingleLineItems(newL)
		}
		c.withMetadata(newL)
	case "Metadata (Dense)":
		if section == "Two-Line List" {
			c.makeTwoLineItems(newL)
		} else {
			c.makeSingleLineItems(newL)
		}
		c.withMetadata(newL)
		newL.Dense = true
	case "Avatar + Metadata":
		c.makeSingleLineItems(newL)
		c.withAvatars(newL)
		c.withIconPlaceholders(newL)
		c.withMetadata(newL)
	case "Avatar + Metadata (Dense)":
		c.makeSingleLineItems(newL)
		c.withAvatars(newL)
		c.withIconPlaceholders(newL)
		c.withMetadata(newL)
		newL.Dense = true
	case "Example - Avatar with Text and Icon":
		c.makeAvatarWTextItems(newL)
		c.withAvatars(newL)
		c.withMetaIcons(newL)
		c.markup = append(c.markup, vecty.Class("demo-list--avatar-and-meta-icon"))
	case "Lists w/ Ellipsis":
		c.makeFoldersItems(newL)
		c.withAvatars(newL)
		c.withIconPlaceholders(newL)
		c.withEllipsis(newL)
	case "hero", "Folders", "Example - Two-line Avatar + Text + Icon":
		c.makeFoldersItems(newL)
		c.withAvatars(newL)
		c.withIconPlaceholders(newL)
	case "Files":
		c.makeFilesItems(newL)
		c.withAvatars(newL)
		c.withIconPlaceholders(newL)
	case "Full-Width Dividers":
		c.makeDividerItems(newL)
	case "Inset Dividers":
		c.makeDividerItems(newL)
		newL.Items[3] = ul.ItemDividerInset()
		c.withAvatars(newL)
		c.withIconPlaceholders(newL)
	case "List 1", "List 2":
		c.makeSingleLineItems(newL)
	case "Example - Interactive List":
		c.makeGraphicExampleItems(newL, true)
	}
	newL.Root = vecty.Markup(c.markup...)
	return newL
}

func (dl *demoList) makeSingleLineItems(l *ul.L) {
	l.Items = []vecty.ComponentOrHTML{
		&ul.Item{Primary: vecty.Text("Single-line item")},
		&ul.Item{Primary: vecty.Text("Single-line item")},
		&ul.Item{Primary: vecty.Text("Single-line item")},
	}
}
func (dl *demoList) makeTwoLineItems(l *ul.L) {
	l.Items = []vecty.ComponentOrHTML{
		&ul.Item{
			Primary:   vecty.Text("Two-line item"),
			Secondary: vecty.Text("Secondary text")},
		&ul.Item{
			Primary:   vecty.Text("Two-line item"),
			Secondary: vecty.Text("Secondary text")},
		&ul.Item{
			Primary:   vecty.Text("Two-line item"),
			Secondary: vecty.Text("Secondary text")},
	}
}

func (dl *demoList) makeFoldersItems(l *ul.L) {
	l.Items = []vecty.ComponentOrHTML{
		&ul.Item{
			Graphic:   &icon.I{Name: "folder"},
			Primary:   vecty.Text("Photos"),
			Secondary: vecty.Text("Jan 9, 2014"),
			Meta:      &icon.I{Name: "info"},
		},
		&ul.Item{
			Graphic:   &icon.I{Name: "folder"},
			Primary:   vecty.Text("Recipes"),
			Secondary: vecty.Text("Jan 17, 2014"),
			Meta:      &icon.I{Name: "info"},
		},
		&ul.Item{
			Graphic:   &icon.I{Name: "folder"},
			Primary:   vecty.Text("Work"),
			Secondary: vecty.Text("Jan 28, 2014"),
			Meta:      &icon.I{Name: "info"},
		},
	}
}

func (dl *demoList) makeFilesItems(l *ul.L) {
	l.Items = []vecty.ComponentOrHTML{
		&ul.Item{
			Graphic: &icon.I{
				Name: "insert_drive_file"},
			Primary:   vecty.Text("Vacation Itinerary"),
			Secondary: vecty.Text("Jan 10, 2014"),
			Meta:      &icon.I{Name: "info"},
		},
		&ul.Item{
			Graphic: &icon.I{
				Name: "insert_drive_file"},
			Primary:   vecty.Text("Kitchen Remodel"),
			Secondary: vecty.Text("Jan 20, 2014"),
			Meta:      &icon.I{Name: "info"},
		},
	}
}

func (dl *demoList) makeDividerItems(l *ul.L) {
	l.Items = []vecty.ComponentOrHTML{
		&ul.Item{
			Primary: vecty.Text("Single-line item - section 1")},
		&ul.Item{
			Primary: vecty.Text("Single-line item - section 1")},
		&ul.Item{
			Primary: vecty.Text("Single-line item - section 1")},
		ul.ItemDivider(),
		&ul.Item{
			Primary: vecty.Text("Single-line item - section 2")},
		&ul.Item{
			Primary: vecty.Text("Single-line item - section 2")},
	}
}

func (dl *demoList) makeGraphicExampleItems(l *ul.L, interactive bool) {
	l.Items = []vecty.ComponentOrHTML{
		&ul.Item{
			Graphic: &icon.I{Name: "network_wifi"},
			Primary: vecty.Text("Wi-Fi"),
		},
		&ul.Item{
			Graphic: &icon.I{Name: "bluetooth"},
			Primary: vecty.Text("Bluetooth"),
		},
		&ul.Item{
			Graphic: &icon.I{Name: "data_usage"},
			Primary: vecty.Text("Data Usage"),
		},
	}
	if interactive {
		for _, cItem := range l.Items {
			if item, ok := cItem.(*ul.Item); ok {
				item.Href = "#"
				item.Root = vecty.Markup(
					event.Click(nil).PreventDefault(),
					&ripple.R{},
				)
			}
		}
	}
}

func (dl *demoList) makeCheckboxItems(l *ul.L, isLeading bool) {
	cbs := []*checkbox.CB{
		&checkbox.CB{
			Input: vecty.Markup(
				prop.ID("leading-checkbox-blueberries"),
			),
		},
		&checkbox.CB{
			Input: vecty.Markup(
				prop.ID("leading-checkbox-boysenberries"),
			),
		},
		&checkbox.CB{
			Input: vecty.Markup(
				prop.ID("leading-checkbox-strawberries"),
			),
		},
	}
	l.Items = []vecty.ComponentOrHTML{
		&ul.Item{
			Primary: elem.Label(vecty.Markup(
				prop.For("leading-checkbox-blueberries")),
				vecty.Text("Blueberries")),
		},
		&ul.Item{
			Primary: elem.Label(vecty.Markup(
				prop.For("leading-checkbox-boysenberries")),
				vecty.Text("Boysenberries")),
		},
		&ul.Item{
			Primary: elem.Label(vecty.Markup(
				prop.For("leading-checkbox-strawberries")),
				vecty.Text("Strawberries")),
		},
	}
	for i, cItem := range l.Items {
		if item, ok := cItem.(*ul.Item); ok {
			item.Root = vecty.Markup(
				vecty.Class("checkbox-list-ripple-surface"),
				&ripple.R{},
			)
			if isLeading {
				item.Graphic = cbs[i]
			} else {
				item.Meta = cbs[i]
			}
			item.OnClick = func(it *ul.Item, e *vecty.Event) {
				var cb *checkbox.CB
				if isLeading {
					cb = item.Graphic.(*checkbox.CB)
					cb.Checked = !cb.Checked
					vecty.Rerender(cb)
				} else {
					cb = item.Meta.(*checkbox.CB)
					cb.Checked = !cb.Checked
					vecty.Rerender(cb)
				}
			}
		}
	}
}

func (dl *demoList) makeAvatarWTextItems(l *ul.L) {
	l.Items = []vecty.ComponentOrHTML{
		&ul.Item{
			Graphic: elem.Image(vecty.Markup(
				prop.Src(MDCImagesURL+"animal1.svg"),
				vecty.Property("width", 56),
				vecty.Property("height", 56),
				vecty.Property("alt", "Panda"))),
			Primary: vecty.Text("Panda"),
		},
		&ul.Item{
			Graphic: elem.Image(vecty.Markup(
				prop.Src(MDCImagesURL+"animal2.svg"),
				vecty.Property("width", 56),
				vecty.Property("height", 56),
				vecty.Property("alt", "Sloth"))),
			Primary: vecty.Text("Sloth"),
		},
		&ul.Item{
			Graphic: elem.Image(vecty.Markup(
				prop.Src(MDCImagesURL+"animal3.svg"),
				vecty.Property("width", 56),
				vecty.Property("height", 56),
				vecty.Property("alt", "Brown Bear"))),
			Primary: vecty.Text("Brown Bear"),
		},
	}
}

func (dl *demoList) withAvatars(l *ul.L) {
	dl.markup = append(dl.markup, vecty.Class("demo-list--with-avatars"))
	l.Avatar = true
}

func (dl *demoList) withIconPlaceholders(l *ul.L) {
	dl.markup = append(dl.markup, vecty.Class("demo-list--icon-placeholders"))
	for _, cItem := range l.Items {
		if item, ok := cItem.(*ul.Item); ok {
			if item.Graphic == nil {
				item.Graphic = elem.Span()
			}
		}
	}
}

func (dl *demoList) withMetadata(l *ul.L) {
	for i, cItem := range l.Items {
		if item, ok := cItem.(*ul.Item); ok {
			item.Meta = vecty.Text("$" + strconv.Itoa(i+1) + "0.00")
		}
	}
}

func (dl *demoList) withMetaIcons(l *ul.L) {
	newItems := make([]vecty.ComponentOrHTML, len(l.Items))
	for i, cItem := range l.Items {
		if item, ok := cItem.(*ul.Item); ok {
			// Add icon to metadata
			iName := "favorite_border"
			if i == len(l.Items)-1 {
				iName = "favorite"
			}
			item.Meta = &icon.I{Name: iName}
			// Reverse the Items order
			newItems[len(l.Items)-i-1] = item
		}
	}
	l.Items = newItems
}

func (dl *demoList) withEllipsis(l *ul.L) {
	if len(l.Items) != 3 {
		print(len(l.Items))
		panic("Expected 3 items for Ellipsis example.")
	}
	for _, cItem := range l.Items {
		if item, ok := cItem.(*ul.Item); ok {
			item.Meta = &icon.I{Name: "folder"}
		}
	}
	l.Items[0].(*ul.Item).Secondary = vecty.Text("This is some secondary text")
	l.Items[1].(*ul.Item).Primary = vecty.Text(
		"Photos of my best photography using my finely tuned skills and eye")
	l.Items[1].(*ul.Item).Secondary = vecty.Text("This is some secondary text")
	l.Items[2].(*ul.Item).Primary = vecty.Text("Work Photos")
	l.Items[2].(*ul.Item).Secondary = vecty.Text(
		"This is a description of work photos from the years 2018 to " +
			"present time while I was a barista")
}
