package main

import (
	mmenu "agamigo.io/material/menu"
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/formfield"
	"agamigo.io/vecty-material/menu"
	"agamigo.io/vecty-material/radio"
	"agamigo.io/vecty-material/ul"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// MenuDemoView is our demo page component.
type MenuDemoView struct {
	vecty.Core
	RememberSelected                               bool
	SelectedIndex                                  int
	menuItems, menuItemsLarge, menuItemsExtraLarge []vecty.ComponentOrHTML
}

func main() {
	mdv := &MenuDemoView{}
	mdv.menuItems = []vecty.ComponentOrHTML{
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Back")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Forward")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Reload")}),
		ul.ItemDivider(),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Save as...")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Help")}),
	}
	mdv.menuItemsLarge = append(mdv.menuItems,
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Settings")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Feedback")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Options...")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 1")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 2")}),
	)
	mdv.menuItemsExtraLarge = append(mdv.menuItemsLarge,
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 3")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 4")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 5")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 6")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 7")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 8")}),
		ul.NewItem(nil, &ul.ItemState{
			Primary: vecty.Text("Item 9")}),
	)
	vecty.RenderBody(mdv)
}

// Render implements the vecty.Component interface.
func (c *MenuDemoView) Render() vecty.ComponentOrHTML {
	heroM := menu.New(nil, &menu.State{
		Open: true,
		List: ul.New(nil, &ul.State{Items: []vecty.ComponentOrHTML{
			ul.NewItem(nil, &ul.ItemState{
				Primary: vecty.Text("Back")}),
			ul.NewItem(nil, &ul.ItemState{
				Primary: vecty.Text("Forward")}),
			ul.NewItem(nil, &ul.ItemState{
				Primary: vecty.Text("Reload")}),
			ul.ItemDivider(),
			ul.NewItem(nil, &ul.ItemState{
				Primary: vecty.Text("Help & Feedback")}),
			ul.NewItem(nil, &ul.ItemState{
				Primary: vecty.Text("Settings")}),
		}}),
	})

	demoM := menu.New(
		&base.Props{
			ID: "demo-menu",
			Markup: []vecty.Applyer{
				vecty.Style("top", "0"),
				vecty.Style("left", "0"),
			},
		},
		&menu.State{
			List: ul.New(nil, nil),
		},
	)
	demoM.List.(*ul.L).Items = c.menuItems
	demoM.SelectHandler = func(index int, item vecty.ComponentOrHTML,
		e *vecty.Event) {
		c.SelectedIndex = index
		for _, lItem := range demoM.List.(*ul.L).Items {
			if ulItem, ok := lItem.(*ul.Item); ok {
				ulItem.Selected = false
			}
		}
		if ulItem, ok := item.(*ul.Item); ok && c.RememberSelected {
			ulItem.Selected = true
		}
		vecty.Rerender(c)
	}
	menuBtnClasses := vecty.ClassMap{"demo-button--normal": true}
	menuBtn := button.New(
		&base.Props{
			ID: "menu-button",
			Markup: []vecty.Applyer{
				vecty.Class("demo-button"),
				menuBtnClasses,
			}},
		&button.State{
			Raised: true,
			Label: vecty.List{
				vecty.Text("Show"),
				elem.Span(vecty.Markup(
					vecty.Class("demo-button__normal-text")),
					vecty.Text(" Menu"),
				),
				elem.Span(vecty.Markup(
					vecty.Class("demo-button__long-text")),
					vecty.Text(" From Here Now!"),
				)},
			ClickHandler: func(thisB *button.B, e *vecty.Event) {
				demoM.Open = !demoM.Open
			},
		})
	demoM.AnchorElement = menuBtn
	btnPositionFunc := func(thisR *radio.R, e *vecty.Event) {
		if thisR.Checked {
			s := demoM.Element.Node().Get("parentNode").Get("style")
			s.Call("removeProperty", "top")
			s.Call("removeProperty", "right")
			s.Call("removeProperty", "bottom")
			s.Call("removeProperty", "left")
			switch thisR.Value {
			case "top left":
				s.Call("setProperty", "top", 0)
				s.Call("setProperty", "left", 0)
			case "top right":
				s.Call("setProperty", "top", 0)
				s.Call("setProperty", "right", 0)
			case "middle left":
				s.Call("setProperty", "top", "35%")
				s.Call("setProperty", "left", 0)
			case "middle right":
				s.Call("setProperty", "top", "35%")
				s.Call("setProperty", "right", 0)
			case "bottom left":
				s.Call("setProperty", "bottom", "0")
				s.Call("setProperty", "left", "0")
			case "bottom right":
				s.Call("setProperty", "bottom", "0")
				s.Call("setProperty", "right", "0")
			}
		}
	}

	menuPositionFunc := func(thisR *radio.R, e *vecty.Event) {
		if thisR.Checked {
			switch thisR.Value {
			case "top start":
				demoM.SetAnchorCorner(mmenu.TOP_START)
			case "top end":
				demoM.SetAnchorCorner(mmenu.TOP_END)
			case "bottom start":
				demoM.SetAnchorCorner(mmenu.BOTTOM_START)
			case "bottom end":
				demoM.SetAnchorCorner(mmenu.BOTTOM_END)
			}
		}
	}

	menuMarginFunc := func(e *vecty.Event) {
		id := e.Target.Get("id").String()
		val := e.Target.Get("value").Int()
		m := demoM.AnchorMargins()
		switch id {
		case "top-margin":
			m.Top = val
		case "bottom-margin":
			m.Bottom = val
		case "left-margin":
			m.Left = val
		case "right-margin":
			m.Right = val
		}
		demoM.SetAnchorMargins(m)
	}

	rtlFunc := func(thisCB *checkbox.CB, e *vecty.Event) {
		d := js.Global.Get("window").Get("document")
		dw := d.Call("getElementById", "demo-wrapper")
		if thisCB.Checked {
			dw.Call("setAttribute", "dir", "rtl")
			return
		}
		dw.Call("removeAttribute", "dir")
	}

	rememberCBFunc := func(thisCB *checkbox.CB, e *vecty.Event) {
		c.RememberSelected = thisCB.Checked
	}

	disableOpenAnimFunc := func(thisCB *checkbox.CB, e *vecty.Event) {
		demoM.QuickOpen = thisCB.Checked
	}

	menuSizeFunc := func(thisR *radio.R, e *vecty.Event) {
		if thisR.Checked {
			switch thisR.Value {
			case "small":
				demoM.List.(*ul.L).Items = c.menuItems
			case "large":
				demoM.List.(*ul.L).Items = c.menuItemsLarge
			case "tall":
				demoM.List.(*ul.L).Items = c.menuItemsExtraLarge
			}
			vecty.Rerender(demoM)
		}
	}

	btnWidthFunc := func(thisR *radio.R, e *vecty.Event) {
		if thisR.Checked {
			switch thisR.Value {
			case "tiny":
				menuBtnClasses["demo-button--normal"] = false
				menuBtnClasses["demo-button--long"] = false
			case "regular":
				menuBtnClasses["demo-button--normal"] = true
				menuBtnClasses["demo-button--long"] = false
			case "wide":
				menuBtnClasses["demo-button--normal"] = true
				menuBtnClasses["demo-button--long"] = true
			}
			vecty.Rerender(menuBtn)
		}
	}

	sItem := c.menuItemsExtraLarge[c.SelectedIndex].(*ul.Item)
	sItemText := js.InternalObject(sItem.Primary).Get("text").String()

	return elem.Body(
		vecty.Markup(
			vecty.Class("mdc-typography"),
		),
		&common.ToolbarHeader{
			Title:      "Menu",
			Navigation: common.NavBack,
		},
		elem.Main(
			elem.Div(vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust"))),
			elem.Section(
				vecty.Markup(vecty.Class("hero")),
				heroM.Render(),
			),
			elem.Div(vecty.Markup(vecty.Class("demo-content")),
				elem.Div(vecty.Markup(prop.ID("demo-wrapper")),
					demoM,
				),
				elem.Div(vecty.Markup(vecty.Class("demo-controls-container")),
					elem.Div(vecty.Markup(vecty.Class("demo-controls")),
						elem.Div(
							vecty.Markup(
								vecty.Class("left-column-controls"),
							),
							vecty.Text("Button Position"),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Top left",
									Input: radio.New(nil,
										&radio.State{
											Name:          "position",
											Value:         "top left",
											Checked:       true,
											ChangeHandler: btnPositionFunc,
										}),
								}),
							),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Top right",
									Input: radio.New(nil,
										&radio.State{
											Name:          "position",
											Value:         "top right",
											ChangeHandler: btnPositionFunc,
										}),
								}),
							),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Middle left",
									Input: radio.New(nil,
										&radio.State{
											Name:          "position",
											Value:         "middle left",
											ChangeHandler: btnPositionFunc,
										}),
								}),
							),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Middle right",
									Input: radio.New(nil,
										&radio.State{
											Name:          "position",
											Value:         "middle right",
											ChangeHandler: btnPositionFunc,
										}),
								}),
							),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Bottom left",
									Input: radio.New(nil,
										&radio.State{
											Name:          "position",
											Value:         "bottom left",
											ChangeHandler: btnPositionFunc,
										}),
								}),
							),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Bottom right",
									Input: radio.New(nil,
										&radio.State{
											Name:          "position",
											Value:         "bottom right",
											ChangeHandler: btnPositionFunc,
										}),
								}),
							),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("right-column-controls"),
							),
							vecty.Text("Default Menu Position"),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Top start",
									Input: radio.New(nil, &radio.State{
										Name:          "menu-position",
										Value:         "top start",
										Checked:       true,
										ChangeHandler: menuPositionFunc,
									}),
								}),
							),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Top end",
									Input: radio.New(nil, &radio.State{
										Name:          "menu-position",
										Value:         "top end",
										ChangeHandler: menuPositionFunc,
									}),
								}),
							),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Bottom start",
									Input: radio.New(nil, &radio.State{
										Name:          "menu-position",
										Value:         "bottom start",
										ChangeHandler: menuPositionFunc,
									}),
								}),
							),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label: "Bottom end",
									Input: radio.New(nil, &radio.State{
										Name:          "menu-position",
										Value:         "bottom end",
										ChangeHandler: menuPositionFunc,
									}),
								}),
							),
						),
						elem.Paragraph(elem.Div(vecty.Markup(
							prop.ID("margin-inputs"),
							vecty.Class("margin-inputs")),
							vecty.Text("Anchor Margins:"),
							elem.Div(
								formfield.New(nil, &formfield.State{
									Label:    "T: ",
									AlignEnd: true,
									Input: elem.Input(vecty.Markup(
										prop.Type(prop.TypeText),
										prop.ID("top-margin"),
										prop.Value("0"),
										vecty.Property("size", 3),
										vecty.Property("max-length", 3),
										event.Change(menuMarginFunc),
									)),
								}),
								formfield.New(nil, &formfield.State{
									Label:    "B: ",
									AlignEnd: true,
									Input: elem.Input(vecty.Markup(
										prop.Type(prop.TypeText),
										prop.ID("bottom-margin"),
										prop.Value("0"),
										vecty.Property("size", 3),
										event.Change(menuMarginFunc),
									)),
								}),
								formfield.New(nil, &formfield.State{
									Label:    "L: ",
									AlignEnd: true,
									Input: elem.Input(vecty.Markup(
										prop.Type(prop.TypeText),
										prop.ID("left-margin"),
										prop.Value("0"),
										vecty.Property("size", 3),
										event.Change(menuMarginFunc),
									)),
								}),
								formfield.New(nil, &formfield.State{
									Label:    "R: ",
									AlignEnd: true,
									Input: elem.Input(vecty.Markup(
										prop.Type(prop.TypeText),
										prop.ID("right-margin"),
										prop.Value("0"),
										vecty.Property("size", 3),
										event.Change(menuMarginFunc),
									)),
								}),
							),
						)),
						elem.Div(formfield.New(nil, &formfield.State{
							Label: "RTL",
							Input: checkbox.New(nil,
								&checkbox.State{
									ChangeHandler: rtlFunc,
								},
							),
						})),
						elem.Div(formfield.New(nil, &formfield.State{
							Label: "Remember Selected Item",
							Input: checkbox.New(nil,
								&checkbox.State{
									ChangeHandler: rememberCBFunc,
								},
							),
						})),
						elem.Div(formfield.New(nil, &formfield.State{
							Label: "Disable Open Animation",
							Input: checkbox.New(nil,
								&checkbox.State{
									ChangeHandler: disableOpenAnimFunc,
								},
							),
						})),
						elem.Paragraph(
							elem.Div(
								vecty.Markup(
									vecty.Class("left-column-controls"),
								),
								vecty.Text("Menu Sizes:"),
								elem.Div(
									formfield.New(nil, &formfield.State{
										Label: "Regular menu",
										Input: radio.New(nil, &radio.State{
											Name:          "menu-length",
											Value:         "small",
											Checked:       true,
											ChangeHandler: menuSizeFunc,
										}),
									}),
								),
								elem.Div(
									formfield.New(nil, &formfield.State{
										Label: "Large menu",
										Input: radio.New(nil, &radio.State{
											Name:          "menu-length",
											Value:         "large",
											ChangeHandler: menuSizeFunc,
										}),
									}),
								),
								elem.Div(
									formfield.New(nil, &formfield.State{
										Label: "Extra tall menu",
										Input: radio.New(nil, &radio.State{
											Name:          "menu-length",
											Value:         "tall",
											ChangeHandler: menuSizeFunc,
										}),
									}),
								),
							),
							elem.Div(vecty.Markup(
								vecty.Class("right-column-controls")),
								vecty.Text("Anchor Widths"),
								elem.Div(
									formfield.New(nil, &formfield.State{
										Label: "Small button",
										Input: radio.New(nil, &radio.State{
											Name:          "anchor-width",
											Value:         "tiny",
											ChangeHandler: btnWidthFunc,
										}),
									}),
								),
								elem.Div(
									formfield.New(nil, &formfield.State{
										Label: "Comparable to menu",
										Input: radio.New(nil, &radio.State{
											Name:          "anchor-width",
											Value:         "regular",
											Checked:       true,
											ChangeHandler: btnWidthFunc,
										}),
									}),
								),
								elem.Div(
									formfield.New(nil, &formfield.State{
										Label: "Wider than menu",
										Input: radio.New(nil, &radio.State{
											Name:          "anchor-width",
											Value:         "wide",
											ChangeHandler: btnWidthFunc,
										}),
									}),
								),
							),
							elem.HorizontalRule(),
							elem.Div(elem.Span(
								vecty.Text("Last Selected item: "),
								elem.Emphasis(vecty.Markup(
									prop.ID("last-selected")),
									vecty.Text(sItemText),
								),
							)),
						),
					),
				),
			),
		),
	)
}
