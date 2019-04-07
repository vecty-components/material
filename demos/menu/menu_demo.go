package main

import (
	"strconv"

	mmenu "agamigo.io/material/menu"
	"agamigo.io/vecty-material/base/applyer"
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/formfield"
	"agamigo.io/vecty-material/menu"
	"agamigo.io/vecty-material/radio"
	"agamigo.io/vecty-material/ul"
	"github.com/gopherjs/gopherwasm/js"
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
	demoM                                          *menu.M
}

type lastSelectedItem struct {
	vecty.Core
	name  string
	index int
}

func main() {
	mdv := &MenuDemoView{}
	mdv.menuItems = []vecty.ComponentOrHTML{
		&ul.Item{Primary: vecty.Text("Back")},
		&ul.Item{Primary: vecty.Text("Forward")},
		&ul.Item{Primary: vecty.Text("Reload")},
		ul.ItemDivider(),
		&ul.Item{Primary: vecty.Text("Save as...")},
		&ul.Item{Primary: vecty.Text("Help")},
	}
	mdv.menuItemsLarge = append(mdv.menuItems,
		&ul.Item{Primary: vecty.Text("Settings")},
		&ul.Item{Primary: vecty.Text("Feedback")},
		&ul.Item{Primary: vecty.Text("Options...")},
		&ul.Item{Primary: vecty.Text("Item 1")},
		&ul.Item{Primary: vecty.Text("Item 2")},
	)
	mdv.menuItemsExtraLarge = append(mdv.menuItemsLarge,
		&ul.Item{Primary: vecty.Text("Item 3")},
		&ul.Item{Primary: vecty.Text("Item 4")},
		&ul.Item{Primary: vecty.Text("Item 5")},
		&ul.Item{Primary: vecty.Text("Item 6")},
		&ul.Item{Primary: vecty.Text("Item 7")},
		&ul.Item{Primary: vecty.Text("Item 8")},
		&ul.Item{Primary: vecty.Text("Item 9")},
	)
	vecty.RenderBody(mdv)
}

// Render implements the vecty.Component interface.
func (c *MenuDemoView) Render() vecty.ComponentOrHTML {
	heroM := &menu.M{
		Open: true,
		Root: vecty.Markup(
			applyer.CSSOnly(),
		),
		List: &ul.L{Items: []vecty.ComponentOrHTML{
			&ul.Item{Primary: vecty.Text("Back")},
			&ul.Item{Primary: vecty.Text("Forward")},
			&ul.Item{Primary: vecty.Text("Reload")},
			ul.ItemDivider(),
			&ul.Item{Primary: vecty.Text("Help & Feedback")},
			&ul.Item{Primary: vecty.Text("Settings")},
		}},
	}

	lsiStatus := &lastSelectedItem{}

	demoM := &menu.M{
		Root: vecty.Markup(
			prop.ID("demo-menu"),
			vecty.Style("top", "0"),
			vecty.Style("left", "0"),
		),
		List: &ul.L{},
	}
	demoM.List.(*ul.L).Items = c.menuItems
	demoM.OnSelect = func(index int, item vecty.ComponentOrHTML,
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
		sItem := c.menuItemsExtraLarge[index].(*ul.Item)
		lsiStatus.name = js.InternalObject(sItem.Primary).Get("text").String()
		lsiStatus.index = index
		vecty.Rerender(lsiStatus)
		vecty.Rerender(demoM)
	}
	menuBtnClasses := vecty.ClassMap{"demo-button--normal": true}
	menuBtn := &button.B{
		Root: vecty.Markup(
			prop.ID("menu-button"),
			vecty.Class("demo-button"),
			menuBtnClasses,
		),
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
		OnClick: func(thisB *button.B, e *vecty.Event) {
			demoM.Open = !demoM.Open
			vecty.Rerender(demoM)
		},
	}
	demoM.AnchorElement = menuBtn
	btnPositionFunc := func(thisR *radio.R, e *vecty.Event) {
		if thisR.Checked {
			s := demoM.MDC.Component.Component().RootElement.Get(
				"parentNode").Get("style")
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
		vecty.Rerender(c)
	}

	rtlFunc := func(thisCB *checkbox.CB, e *vecty.Event) {
		d := js.Global().Get("window").Get("document")
		dw := d.Call("getElementById", "demo-wrapper")
		if thisCB.Checked {
			dw.Call("setAttribute", "dir", "rtl")
			return
		}
		dw.Call("removeAttribute", "dir")
		vecty.Rerender(demoM)
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
				heroM,
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
								&formfield.FF{
									Label: "Top left",
									Input: &radio.R{
										Name:     "position",
										Value:    "top left",
										Checked:  true,
										OnChange: btnPositionFunc,
									},
								},
							),
							elem.Div(
								&formfield.FF{
									Label: "Top right",
									Input: &radio.R{
										Name:     "position",
										Value:    "top right",
										OnChange: btnPositionFunc,
									},
								},
							),
							elem.Div(
								&formfield.FF{
									Label: "Middle left",
									Input: &radio.R{
										Name:     "position",
										Value:    "middle left",
										OnChange: btnPositionFunc,
									},
								},
							),
							elem.Div(
								&formfield.FF{
									Label: "Middle right",
									Input: &radio.R{
										Name:     "position",
										Value:    "middle right",
										OnChange: btnPositionFunc,
									},
								},
							),
							elem.Div(
								&formfield.FF{
									Label: "Bottom left",
									Input: &radio.R{
										Name:     "position",
										Value:    "bottom left",
										OnChange: btnPositionFunc,
									},
								},
							),
							elem.Div(
								&formfield.FF{
									Label: "Bottom right",
									Input: &radio.R{
										Name:     "position",
										Value:    "bottom right",
										OnChange: btnPositionFunc,
									},
								},
							),
						),
						elem.Div(
							vecty.Markup(
								vecty.Class("right-column-controls"),
							),
							vecty.Text("Default Menu Position"),
							elem.Div(
								&formfield.FF{
									Label: "Top start",
									Input: &radio.R{
										Name:     "menu-position",
										Value:    "top start",
										Checked:  true,
										OnChange: menuPositionFunc,
									},
								},
							),
							elem.Div(
								&formfield.FF{
									Label: "Top end",
									Input: &radio.R{
										Name:     "menu-position",
										Value:    "top end",
										OnChange: menuPositionFunc,
									},
								},
							),
							elem.Div(
								&formfield.FF{
									Label: "Bottom start",
									Input: &radio.R{
										Name:     "menu-position",
										Value:    "bottom start",
										OnChange: menuPositionFunc,
									},
								},
							),
							elem.Div(
								&formfield.FF{
									Label: "Bottom end",
									Input: &radio.R{
										Name:     "menu-position",
										Value:    "bottom end",
										OnChange: menuPositionFunc,
									},
								},
							),
						),
						elem.Paragraph(elem.Div(vecty.Markup(
							prop.ID("margin-inputs"),
							vecty.Class("margin-inputs")),
							vecty.Text("Anchor Margins:"),
							elem.Div(
								&formfield.FF{
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
								},
								&formfield.FF{
									Label:    "B: ",
									AlignEnd: true,
									Input: elem.Input(vecty.Markup(
										prop.Type(prop.TypeText),
										prop.ID("bottom-margin"),
										prop.Value("0"),
										vecty.Property("size", 3),
										event.Change(menuMarginFunc),
									)),
								},
								&formfield.FF{
									Label:    "L: ",
									AlignEnd: true,
									Input: elem.Input(vecty.Markup(
										prop.Type(prop.TypeText),
										prop.ID("left-margin"),
										prop.Value("0"),
										vecty.Property("size", 3),
										event.Change(menuMarginFunc),
									)),
								},
								&formfield.FF{
									Label:    "R: ",
									AlignEnd: true,
									Input: elem.Input(vecty.Markup(
										prop.Type(prop.TypeText),
										prop.ID("right-margin"),
										prop.Value("0"),
										vecty.Property("size", 3),
										event.Change(menuMarginFunc),
									)),
								},
							),
						)),
						elem.Div(&formfield.FF{
							Label: "RTL",
							Input: &checkbox.CB{
								OnChange: rtlFunc,
							},
						}),
						elem.Div(&formfield.FF{
							Label: "Remember Selected Item",
							Input: &checkbox.CB{
								OnChange: rememberCBFunc,
							},
						}),
						elem.Div(&formfield.FF{
							Label: "Disable Open Animation",
							Input: &checkbox.CB{
								OnChange: disableOpenAnimFunc,
							},
						}),
						elem.Paragraph(
							elem.Div(
								vecty.Markup(
									vecty.Class("left-column-controls"),
								),
								vecty.Text("Menu Sizes:"),
								elem.Div(
									&formfield.FF{
										Label: "Regular menu",
										Input: &radio.R{
											Name:     "menu-length",
											Value:    "small",
											Checked:  true,
											OnChange: menuSizeFunc,
										},
									},
								),
								elem.Div(
									&formfield.FF{
										Label: "Large menu",
										Input: &radio.R{
											Name:     "menu-length",
											Value:    "large",
											OnChange: menuSizeFunc,
										},
									},
								),
								elem.Div(
									&formfield.FF{
										Label: "Extra tall menu",
										Input: &radio.R{
											Name:     "menu-length",
											Value:    "tall",
											OnChange: menuSizeFunc,
										},
									},
								),
							),
							elem.Div(vecty.Markup(
								vecty.Class("right-column-controls")),
								vecty.Text("Anchor Widths"),
								elem.Div(
									&formfield.FF{
										Label: "Small button",
										Input: &radio.R{
											Name:     "anchor-width",
											Value:    "tiny",
											OnChange: btnWidthFunc,
										},
									},
								),
								elem.Div(
									&formfield.FF{
										Label: "Comparable to menu",
										Input: &radio.R{
											Name:     "anchor-width",
											Value:    "regular",
											Checked:  true,
											OnChange: btnWidthFunc,
										},
									},
								),
								elem.Div(
									&formfield.FF{
										Label: "Wider than menu",
										Input: &radio.R{
											Name:     "anchor-width",
											Value:    "wide",
											OnChange: btnWidthFunc,
										},
									},
								),
							),
							elem.HorizontalRule(),
							lsiStatus,
						),
					),
				),
			),
		),
	)
}

func (c *lastSelectedItem) Render() vecty.ComponentOrHTML {
	return elem.Div(elem.Span(
		vecty.Text("Last Selected item: "),
		elem.Emphasis(vecty.Markup(
			prop.ID("last-selected")),
			vecty.If(c.name == "", vecty.Text("<none selected>")),
			vecty.If(c.name != "",
				vecty.Text("\""+c.name+"\" at index "+strconv.Itoa(c.index)),
			),
		),
	))
}
