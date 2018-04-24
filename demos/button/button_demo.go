package main

import (
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/formfield"
	"agamigo.io/vecty-material/icon"
	"agamigo.io/vecty-material/ripple"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// buttonDemo is our main page component.
type buttonDemoView struct {
	vecty.Core
	buttons []*button.B
}

func main() {
	vecty.RenderBody(&buttonDemoView{})
}

// Render implements the vecty.Component interface.
func (c *buttonDemoView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(vecty.Class("mdc-typography")),
		&common.ToolbarHeader{
			Title:      "Buttons",
			Navigation: common.NavBack,
		},
		elem.Main(
			elem.Div(vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust"))),
			elem.Section(
				vecty.Markup(vecty.Class("hero")),
				c.newBtn(&button.B{Label: vecty.Text("Flat")}),
				c.newBtn(&button.B{Label: vecty.Text("Raised"), Raised: true}),
				elem.Small(
					vecty.Markup(vecty.Class("note")),
					vecty.Text(`Note: "secondary" was previously called `+
						`"accent" in the Material spec.`),
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("demo-wrapper")),
				&formfield.FF{
					Label: "Disable buttons (excluding links)",
					Input: &checkbox.CB{
						Input: vecty.Markup(
							prop.ID("toggle-disabled"),
						),
						OnChange: func(thisCB *checkbox.CB,
							e *vecty.Event) {
							checked := e.Target.Get("checked").Bool()
							for _, b := range c.buttons {
								if b.Href != "" {
									continue
								}
								b.Disabled = checked
								vecty.Rerender(b)
							}
						},
					},
				},
			),
		),
		c.renderBtnFieldSets("Ripple Enabled", true),
		c.renderBtnFieldSets("CSS Only", false),
	)
}

func r(applyRipple bool) vecty.Applyer {
	return vecty.MarkupIf(applyRipple,
		&ripple.R{},
	)
}

func (c *buttonDemoView) renderBtnFieldSet(title string, Ripple bool,
	b *button.B) vecty.ComponentOrHTML {
	if b == nil {
		b = &button.B{}
	}
	return elem.FieldSet(
		elem.Legend(
			vecty.Markup(vecty.Class("mdc-typography--title")),
			vecty.Text(title),
		),
		elem.Div(
			c.newBtn(
				&button.B{
					Label:      vecty.Text("Baseline"),
					Raised:     b.Raised,
					Unelevated: b.Unelevated,
					Outlined:   b.Outlined,
					Root: vecty.Markup(
						r(Ripple),
					),
				},
			),
			c.newBtn(
				&button.B{
					Label:      vecty.Text("Dense"),
					Dense:      true,
					Raised:     b.Raised,
					Unelevated: b.Unelevated,
					Outlined:   b.Outlined,
					Root: vecty.Markup(
						r(Ripple),
					),
				},
			),
			c.newBtn(
				&button.B{
					Root: vecty.Markup(
						vecty.Class("secondary-text-button"),
						r(Ripple),
					),
					Label:      vecty.Text("Secondary"),
					Raised:     b.Raised,
					Unelevated: b.Unelevated,
					Outlined:   b.Outlined,
				},
			),
			c.newBtn(
				&button.B{
					Label: vecty.Text("Icon"),
					Icon: &icon.I{
						Name: "favorite",
					},
					Raised:     b.Raised,
					Unelevated: b.Unelevated,
					Outlined:   b.Outlined,
					Root: vecty.Markup(
						r(Ripple),
					),
				},
			),
			c.newBtn(
				&button.B{
					Label:      vecty.Text("Link"),
					Href:       "javascript:void(0)",
					Raised:     b.Raised,
					Unelevated: b.Unelevated,
					Outlined:   b.Outlined,
					Root: vecty.Markup(
						r(Ripple),
					),
				},
			),
		),
	)
}

func (c *buttonDemoView) renderBtnFieldSets(heading string,
	Ripple bool) vecty.ComponentOrHTML {
	return vecty.List{
		elem.Heading1(
			vecty.Markup(vecty.Class("mdc-typography--display2")),
			vecty.Text(heading),
		),
		c.renderBtnFieldSet("Text Button", Ripple, nil),
		c.renderBtnFieldSet("Raised Button", Ripple,
			&button.B{Raised: true},
		),
		c.renderBtnFieldSet("Unelevated Button (Experimental)", Ripple,
			&button.B{Unelevated: true},
		),
		c.renderBtnFieldSet("Outlined Button", Ripple,
			&button.B{Outlined: true},
		),
		elem.FieldSet(
			elem.Legend(
				vecty.Markup(
					vecty.Class("mdc-typography--title"),
				),
				vecty.Text("Custom button (Experimental)"),
			),
			elem.Div(
				c.newBtn(
					&button.B{
						Root: vecty.Markup(
							vecty.Class("big-round-corner-button"),
							r(Ripple),
						),
						Unelevated: true,
						Label:      vecty.Text("Corner Radius"),
					},
				),
				c.newBtn(
					&button.B{
						Root: vecty.Markup(
							vecty.Class("thick-stroke-button"),
							r(Ripple),
						),
						Outlined: true,
						Label:    vecty.Text("Thick Stroke Width"),
					},
				),
			),
		),
	}
}

// Wraps button.New() and keeps track of created buttons
func (c *buttonDemoView) newBtn(b *button.B) *button.B {
	c.buttons = append(c.buttons, b)
	return b
}
