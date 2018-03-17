package main

import (
	"strconv"

	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/formfield"
	"agamigo.io/vecty-material/icon"
	"agamigo.io/vecty-material/ul"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
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
				formfield.New(nil,
					&formfield.State{
						Label: "Toggle RTL",
						Input: checkbox.New(&base.Props{ID: "toggle-rtl"}, nil),
					},
				),
			),
			elem.Div(vecty.Markup(prop.ID("demo-wrapper")),
				&demoSection{
					heading: "Custom Colors",
					groups: []*ul.Group{
						ul.NewGroup(&base.Props{Markup: []vecty.Applyer{
							vecty.Class("demo-list-group--custom")}},
							nil)},
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
						ul.NewGroup(nil, nil),
						ul.NewGroup(nil, nil)},
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
	newL := ul.New(&base.Props{
		Markup: []vecty.Applyer{vecty.Class("demo-list")}},
		nil)
	newL.Markup = append(newL.Markup, c.markup...)
	newL.GroupSubheader = c.heading
	switch c.heading {
	case "Text only, non-interactive (no states)":
		makeSingleLineItems(newL)
		newL.NonInteractive = true
	case "Text-Only":
		makeTwoLineItems(newL)
	case "Text only (dense)":
		makeSingleLineItems(newL)
		newL.Dense = true
	case "Text-Only (Dense)":
		makeTwoLineItems(newL)
		newL.Dense = true
	case "Graphic":
		if section == "Two-Line List" {
			makeTwoLineItems(newL)
		} else {
			makeSingleLineItems(newL)
		}
		withIconPlaceholders(newL)
	case "Graphic (dense)":
		makeSingleLineItems(newL)
		withIconPlaceholders(newL)
		newL.Dense = true
	case "Graphic (Dense)":
		makeTwoLineItems(newL)
		withIconPlaceholders(newL)
		newL.Dense = true
	case "Graphic Example - Icon with Text":
		makeGraphicExampleItems(newL, false)
	case "Leading Checkbox":
		makeCheckboxItems(newL, true)
		newL.ID = "leading-checkbox-list"
	case "Trailing Checkbox":
		makeCheckboxItems(newL, false)
		newL.ID = "trailing-checkbox-list"
	case "Avatar List":
		if section == "Two-Line List" {
			makeTwoLineItems(newL)
		} else {
			makeSingleLineItems(newL)
		}
		withAvatars(newL)
		withIconPlaceholders(newL)
	case "Avatar List (dense)":
		if section == "Two-Line List" {
			makeTwoLineItems(newL)
		} else {
			makeSingleLineItems(newL)
		}
		withAvatars(newL)
		withIconPlaceholders(newL)
		newL.Dense = true
	case "Example - Avatar with Text":
		makeAvatarWTextItems(newL)
		withAvatars(newL)
	case "Metadata":
		if section == "Two-Line List" {
			makeTwoLineItems(newL)
		} else {
			makeSingleLineItems(newL)
		}
		withMetadata(newL)
	case "Metadata (Dense)":
		if section == "Two-Line List" {
			makeTwoLineItems(newL)
		} else {
			makeSingleLineItems(newL)
		}
		withMetadata(newL)
		newL.Dense = true
	case "Avatar + Metadata":
		makeSingleLineItems(newL)
		withAvatars(newL)
		withIconPlaceholders(newL)
		withMetadata(newL)
	case "Avatar + Metadata (Dense)":
		makeSingleLineItems(newL)
		withAvatars(newL)
		withIconPlaceholders(newL)
		withMetadata(newL)
		newL.Dense = true
	case "Example - Avatar with Text and Icon":
		makeAvatarWTextItems(newL)
		withAvatars(newL)
		withMetaIcons(newL)
		newL.Markup = append(newL.Markup,
			vecty.Class("demo-list--avatar-and-meta-icon"))
	case "Lists w/ Ellipsis":
		makeFoldersItems(newL)
		withAvatars(newL)
		withIconPlaceholders(newL)
		withEllipsis(newL)
	case "hero", "Folders", "Example - Two-line Avatar + Text + Icon":
		makeFoldersItems(newL)
		withAvatars(newL)
		withIconPlaceholders(newL)
	case "Files":
		makeFilesItems(newL)
		withAvatars(newL)
		withIconPlaceholders(newL)
	case "Full-Width Dividers":
		makeDividerItems(newL)
	case "Inset Dividers":
		makeDividerItems(newL)
		newL.Items[3] = ul.ItemDividerInset()
		withAvatars(newL)
		withIconPlaceholders(newL)
	case "List 1", "List 2":
		makeSingleLineItems(newL)
	case "Example - Interactive List":
		makeGraphicExampleItems(newL, true)
	}
	return newL
}

func makeSingleLineItems(l *ul.L) {
	l.Items = []*ul.Item{
		ul.NewItem(nil, &ul.ItemState{Primary: vecty.Text("Single-line item")}),
		ul.NewItem(nil, &ul.ItemState{Primary: vecty.Text("Single-line item")}),
		ul.NewItem(nil, &ul.ItemState{Primary: vecty.Text("Single-line item")}),
	}
}
func makeTwoLineItems(l *ul.L) {
	l.Items = []*ul.Item{
		ul.NewItem(nil, &ul.ItemState{
			Primary:   vecty.Text("Two-line item"),
			Secondary: vecty.Text("Secondary text")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary:   vecty.Text("Two-line item"),
			Secondary: vecty.Text("Secondary text")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary:   vecty.Text("Two-line item"),
			Secondary: vecty.Text("Secondary text")}),
	}
}

func makeFoldersItems(l *ul.L) {
	l.Items = []*ul.Item{
		ul.NewItem(nil, &ul.ItemState{
			Graphic:   icon.New(nil, &icon.State{Name: "folder"}),
			Primary:   vecty.Text("Photos"),
			Secondary: vecty.Text("Jan 9, 2014"),
			Meta:      icon.New(nil, &icon.State{Name: "info"}),
		}),
		ul.NewItem(nil, &ul.ItemState{
			Graphic:   icon.New(nil, &icon.State{Name: "folder"}),
			Primary:   vecty.Text("Recipes"),
			Secondary: vecty.Text("Jan 17, 2014"),
			Meta:      icon.New(nil, &icon.State{Name: "info"}),
		}),
		ul.NewItem(nil, &ul.ItemState{
			Graphic:   icon.New(nil, &icon.State{Name: "folder"}),
			Primary:   vecty.Text("Work"),
			Secondary: vecty.Text("Jan 28, 2014"),
			Meta:      icon.New(nil, &icon.State{Name: "info"}),
		}),
	}
}

func makeFilesItems(l *ul.L) {
	l.Items = []*ul.Item{
		ul.NewItem(nil, &ul.ItemState{
			Graphic: icon.New(nil, &icon.State{
				Name: "insert_drive_file"}),
			Primary:   vecty.Text("Vacation Itinerary"),
			Secondary: vecty.Text("Jan 10, 2014"),
			Meta:      icon.New(nil, &icon.State{Name: "info"}),
		}),
		ul.NewItem(nil, &ul.ItemState{
			Graphic: icon.New(nil, &icon.State{
				Name: "insert_drive_file"}),
			Primary:   vecty.Text("Kitchen Remodel"),
			Secondary: vecty.Text("Jan 20, 2014"),
			Meta:      icon.New(nil, &icon.State{Name: "info"}),
		}),
	}
}

func makeDividerItems(l *ul.L) {
	l.Items = []*ul.Item{
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Single-line item - section 1")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Single-line item - section 1")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Single-line item - section 1")}),
		ul.ItemDivider(),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Single-line item - section 2")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Single-line item - section 2")}),
	}
}

func makeGraphicExampleItems(l *ul.L, interactive bool) {
	l.Items = []*ul.Item{
		ul.NewItem(nil, &ul.ItemState{
			Graphic: icon.New(nil, &icon.State{Name: "network_wifi"}),
			Primary: vecty.Text("Wi-Fi"),
		}),
		ul.NewItem(nil, &ul.ItemState{
			Graphic: icon.New(nil, &icon.State{Name: "bluetooth"}),
			Primary: vecty.Text("Bluetooth"),
		}),
		ul.NewItem(nil, &ul.ItemState{
			Graphic: icon.New(nil, &icon.State{Name: "data_usage"}),
			Primary: vecty.Text("Data Usage"),
		}),
	}
	if interactive {
		for _, item := range l.Items {
			item.Href = "#"
			item.Props.Markup = append(item.Props.Markup,
				event.Click(nil).PreventDefault())
			item.Ripple = true
		}
	}
}

func makeCheckboxItems(l *ul.L, isLeading bool) {
	cbs := []*checkbox.CB{
		checkbox.New(&base.Props{ID: "leading-checkbox-blueberries"}, nil),
		checkbox.New(&base.Props{ID: "leading-checkbox-boysenberries"}, nil),
		checkbox.New(&base.Props{ID: "leading-checkbox-strawberries"}, nil),
	}
	l.Items = []*ul.Item{
		ul.NewItem(nil, &ul.ItemState{
			Primary: elem.Label(vecty.Markup(
				prop.For("leading-checkbox-blueberries")),
				vecty.Text("Blueberries")),
		}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: elem.Label(vecty.Markup(
				prop.For("leading-checkbox-boysenberries")),
				vecty.Text("Boysenberries")),
		}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: elem.Label(vecty.Markup(
				prop.For("leading-checkbox-strawberries")),
				vecty.Text("Strawberries")),
		}),
	}
	for i, item := range l.Items {
		item.Markup = append(item.Markup,
			vecty.Class("checkbox-list-ripple-surface"))
		item.Ripple = true
		if isLeading {
			item.Graphic = cbs[i]
		} else {
			item.Meta = cbs[i]
		}
		cb := cbs[i]
		item.ClickHandler = func(it *ul.Item, e *vecty.Event) {
			// TODO: Figure out why we have to update the native input
			// directly
			if e.Target.Get("tagName").String() == "LI" {
				cb.Checked = !cb.Checked
				cbEl := e.Target.Call("querySelector",
					".mdc-checkbox__native-control")
				cbEl.Set("checked", cb.Checked)
			}
		}
	}
}

func makeAvatarWTextItems(l *ul.L) {
	l.Items = []*ul.Item{
		ul.NewItem(nil,
			&ul.ItemState{
				Graphic: elem.Image(vecty.Markup(
					prop.Src(MDCImagesURL+"animal1.svg"),
					vecty.Property("width", 56),
					vecty.Property("height", 56),
					vecty.Property("alt", "Panda"))),
				Primary: vecty.Text("Panda"),
			},
		),
		ul.NewItem(nil,
			&ul.ItemState{
				Graphic: elem.Image(vecty.Markup(
					prop.Src(MDCImagesURL+"animal2.svg"),
					vecty.Property("width", 56),
					vecty.Property("height", 56),
					vecty.Property("alt", "Sloth"))),
				Primary: vecty.Text("Sloth"),
			},
		),
		ul.NewItem(nil,
			&ul.ItemState{
				Graphic: elem.Image(vecty.Markup(
					prop.Src(MDCImagesURL+"animal3.svg"),
					vecty.Property("width", 56),
					vecty.Property("height", 56),
					vecty.Property("alt", "Brown Bear"))),
				Primary: vecty.Text("Brown Bear"),
			},
		),
	}
}

func withAvatars(l *ul.L) {
	l.Markup = append(l.Markup, vecty.Class("demo-list--with-avatars"))
	l.Avatar = true
}

func withIconPlaceholders(l *ul.L) {
	l.Markup = append(l.Markup, vecty.Class("demo-list--icon-placeholders"))
	for _, item := range l.Items {
		if item.Graphic == nil {
			item.Graphic = elem.Span()
		}
	}
}

func withMetadata(l *ul.L) {
	for i, item := range l.Items {
		item.Meta = vecty.Text("$" + strconv.Itoa(i+1) + "0.00")
	}
}

func withMetaIcons(l *ul.L) {
	newItems := make([]*ul.Item, len(l.Items))
	for i, item := range l.Items {
		// Add icon to metadata
		iName := "favorite_border"
		if i == len(l.Items)-1 {
			iName = "favorite"
		}
		item.Meta = icon.New(nil, &icon.State{Name: iName})
		// Reverse the Items order
		newItems[len(l.Items)-i-1] = item
	}
	l.Items = newItems
}

func withEllipsis(l *ul.L) {
	if len(l.Items) != 3 {
		print(len(l.Items))
		panic("Expected 3 items for Ellipsis example.")
	}
	for _, item := range l.Items {
		item.Meta = icon.New(nil, &icon.State{Name: "folder"})
	}
	l.Items[0].Secondary = vecty.Text("This is some secondary text")
	l.Items[1].Primary = vecty.Text("Photos of my best photography using my " +
		"finely tuned skills and eye")
	l.Items[1].Secondary = vecty.Text("This is some secondary text")
	l.Items[2].Primary = vecty.Text("Work Photos")
	l.Items[2].Secondary = vecty.Text("This is a description of work photos from " +
		"the years 2018 to present time while I was a barista")
}
