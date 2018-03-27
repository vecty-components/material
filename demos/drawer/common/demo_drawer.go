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
	var toolbarSpacer, header *vecty.HTML
	switch dType {
	case drawer.Temporary:
		header = elem.Div(
			vecty.Markup(
				vecty.Class("mdc-theme--text-primary-on-primary"),
				vecty.Class("mdc-theme--primary-bg"),
			),
			vecty.Text("Header here"),
		)
	case drawer.Persistent:
		toolbarSpacer = elem.Div()
	}
	return drawer.New(
		&base.Props{
			ID: "demo-drawer",
			Markup: []vecty.Applyer{
				vecty.Class("demo-drawer"),
			},
		},
		&drawer.State{
			Type:          dType,
			Header:        header,
			ToolbarSpacer: toolbarSpacer,
			Content: ul.NewGroup(nil,
				&ul.GroupState{Lists: []vecty.ComponentOrHTML{
					ul.New(nil,
						&ul.State{Items: []vecty.ComponentOrHTML{
							iconListItem("inbox", "Inbox"),
							iconListItem("star", "Star"),
							iconListItem("send", "Sent Mail"),
							iconListItem("drafts", "Drafts")},
							ClickHandler: func(l *ul.L, i *ul.Item,
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
					),
					ul.ListDivider(),
					ul.New(nil,
						&ul.State{Items: []vecty.ComponentOrHTML{
							iconListItem("email", "All Mail"),
							iconListItem("delete", "Trash"),
							iconListItem("report", "Spam"),
						}},
					),
				}},
			),
		},
	)
}

func iconListItem(ico, text string) *ul.Item {
	var selected bool
	if ico == "inbox" {
		selected = true
	}
	return ul.NewItem(
		&base.Props{
			Markup: []vecty.Applyer{
				vecty.Class("demo-drawer-list-item"),
			},
		},
		&ul.ItemState{
			Selected: selected,
			Graphic: vecty.List{
				icon.New(
					&base.Props{Markup: []vecty.Applyer{
						vecty.Attribute("aria-hidden", "true"),
					}},
					&icon.State{Name: ico},
				),
			},
			Primary: vecty.Text(text),
		},
	)
}
