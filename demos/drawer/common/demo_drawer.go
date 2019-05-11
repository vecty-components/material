package common

import (
	"github.com/vecty-material/vecty-material/base"
	"github.com/vecty-material/vecty-material/drawer"
	"github.com/vecty-material/vecty-material/icon"
	"github.com/vecty-material/vecty-material/ul"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func NewDemoDrawer(dType drawer.Type) *drawer.D {
	var toolbarSpacer *vecty.HTML
	var header vecty.ComponentOrHTML
	switch dType {
	case drawer.Temporary:
		header = elem.Div(
			vecty.Markup(
				vecty.Class("mdc-theme--text-primary-on-primary"),
				vecty.Class("mdc-theme--primary-bg"),
			),
			base.RenderStoredChild(vecty.Text("Header here")),
		)
	case drawer.Persistent:
		toolbarSpacer = elem.Div()
	}
	return &drawer.D{
		Root: vecty.Markup(
			vecty.Class("demo-drawer"),
			prop.ID("demo-drawer"),
		),
		Type:          dType,
		Header:        header,
		ToolbarSpacer: base.RenderStoredChild(toolbarSpacer),
		Content: &ul.Group{
			Lists: []vecty.ComponentOrHTML{
				&ul.L{
					Items: []vecty.ComponentOrHTML{
						iconListItem("inbox", "Inbox"),
						iconListItem("star", "Star"),
						iconListItem("send", "Sent Mail"),
						iconListItem("drafts", "Drafts")},
					OnClick: func(l *ul.L, i *ul.Item,
						e *vecty.Event) {
						for _, v := range l.Items {
							if ulItem, ok := v.(*ul.Item); ok {
								switch {
								case i == ulItem:
									i.Activated = true
									vecty.Rerender(i)
								default:
									ulItem.Activated = false
									vecty.Rerender(ulItem)
								}
							}
						}
					},
				},
				ul.ListDivider(),
				&ul.L{
					Items: []vecty.ComponentOrHTML{
						iconListItem("email", "All Mail"),
						iconListItem("delete", "Trash"),
						iconListItem("report", "Spam"),
					}},
			}},
	}
}

func iconListItem(ico, text string) *ul.Item {
	var selected bool
	if ico == "inbox" {
		selected = true
	}
	return &ul.Item{
		Root: vecty.Markup(
			vecty.Class("demo-drawer-list-item"),
		),
		Selected: selected,
		Graphic: vecty.List{
			&icon.I{
				Root: vecty.Markup(
					vecty.Attribute("aria-hidden", "true"),
				),
				Name: ico,
			},
		},
		Primary: vecty.Text(text),
	}
}
