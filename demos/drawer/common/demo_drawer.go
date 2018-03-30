package common

import (
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/drawer"
	"agamigo.io/vecty-material/icon"
	"agamigo.io/vecty-material/ul"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
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
		ID: "demo-drawer",
		Markup: []vecty.Applyer{
			vecty.Class("demo-drawer"),
		},
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
		Markup: []vecty.Applyer{
			vecty.Class("demo-drawer-list-item"),
		},
		Selected: selected,
		Graphic: vecty.List{
			&icon.I{
				Markup: []vecty.Applyer{
					vecty.Attribute("aria-hidden", "true"),
				},
				Name: ico,
			},
		},
		Primary: vecty.Text(text),
	}
}
