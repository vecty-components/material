package main

import (
	"agamigo.io/vecty-material/base"
	"agamigo.io/vecty-material/button"
	"agamigo.io/vecty-material/checkbox"
	"agamigo.io/vecty-material/demos/common"
	"agamigo.io/vecty-material/formfield"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
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
		&common.ToolbarHeader{Title: "Buttons"},
		elem.Main(
			vecty.Markup(vecty.Class("mdc-toolbar-fixed-adjust")),
			elem.Section(
				vecty.Markup(vecty.Class("hero")),
				c.newBtn(nil,
					&button.State{Label: vecty.Text("Flat")},
				),
				c.newBtn(nil,
					&button.State{Label: vecty.Text("Raised"), Raised: true},
				),
				elem.Small(
					vecty.Markup(vecty.Class("note")),
					vecty.Text(`Note: "secondary" was previously called `+
						`"accent" in the Material spec.`),
				),
			),
			elem.Section(
				vecty.Markup(vecty.Class("demo-wrapper")),
				formfield.New(nil,
					&formfield.State{
						Label: "Disable buttons (excluding links)",
						Input: checkbox.New(
							&base.Props{
								ID: "toggle-disabled",
							},
							&checkbox.State{
								ChangeHandler: func(e *vecty.Event) {
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
						),
					},
				),
				c.renderBtnFieldSets("Ripple Enabled", false),
				c.renderBtnFieldSets("CSS Only", true),
			),
		),
	)
}

func (c *buttonDemoView) renderBtnFieldSet(title string, noRipple bool,
	s *button.State) vecty.ComponentOrHTML {
	if s == nil {
		s = &button.State{}
	}
	return elem.FieldSet(
		elem.Legend(
			vecty.Markup(vecty.Class("mdc-typography--title")),
			vecty.Text(title),
		),
		elem.Div(
			c.newBtn(&base.Props{NoRipple: noRipple},
				&button.State{
					Label:      vecty.Text("Baseline"),
					Raised:     s.Raised,
					Unelevated: s.Unelevated,
					Stroked:    s.Stroked,
				},
			),
			c.newBtn(&base.Props{NoRipple: noRipple},
				&button.State{
					Label:      vecty.Text("Compact"),
					Compact:    true,
					Raised:     s.Raised,
					Unelevated: s.Unelevated,
					Stroked:    s.Stroked,
				},
			),
			c.newBtn(&base.Props{NoRipple: noRipple},
				&button.State{
					Label:      vecty.Text("Dense"),
					Dense:      true,
					Raised:     s.Raised,
					Unelevated: s.Unelevated,
					Stroked:    s.Stroked,
				},
			),
			c.newBtn(
				&base.Props{
					NoRipple: noRipple,
					Markup: vecty.Markup(
						vecty.Class("secondary-text-button"),
					),
				},
				&button.State{
					Label:      vecty.Text("Secondary"),
					Raised:     s.Raised,
					Unelevated: s.Unelevated,
					Stroked:    s.Stroked,
				},
			),
			c.newBtn(&base.Props{NoRipple: noRipple},
				&button.State{
					Label: vecty.List{
						elem.Italic(
							vecty.Markup(
								vecty.Class("material-icons"),
								vecty.Class("mdc-button__icon"),
							),
							vecty.Text("favorite"),
						),
						vecty.Text("Icon"),
					},
					Raised:     s.Raised,
					Unelevated: s.Unelevated,
					Stroked:    s.Stroked,
				},
			),
			c.newBtn(&base.Props{NoRipple: noRipple},
				&button.State{
					Label:      vecty.Text("Link"),
					Href:       "javascript:void(0)",
					Raised:     s.Raised,
					Unelevated: s.Unelevated,
					Stroked:    s.Stroked,
				},
			),
		),
	)
}

func (c *buttonDemoView) renderBtnFieldSets(heading string,
	noRipple bool) vecty.ComponentOrHTML {
	return vecty.List{
		elem.Heading1(
			vecty.Markup(vecty.Class("mdc-typography--display2")),
			vecty.Text(heading),
		),
		c.renderBtnFieldSet("Text Button", noRipple, nil),
		c.renderBtnFieldSet("Raised Button", noRipple,
			&button.State{Raised: true},
		),
		c.renderBtnFieldSet("Unelevated Button (Experimental)", noRipple,
			&button.State{Unelevated: true},
		),
		c.renderBtnFieldSet("Stroked Button", noRipple,
			&button.State{Stroked: true},
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
					&base.Props{
						Markup: vecty.Markup(
							vecty.Class("big-round-corner-button")),
						NoRipple: noRipple,
					},
					&button.State{
						Unelevated: true,
						Label:      vecty.Text("Corner Radius"),
					},
				),
				c.newBtn(
					&base.Props{
						Markup: vecty.Markup(
							vecty.Class("thick-stroke-button")),
						NoRipple: noRipple,
					},
					&button.State{
						Stroked: true,
						Label:   vecty.Text("Thick Stroke Width"),
					},
				),
			),
		),
	}
}

// Wraps button.New() and keeps track of created buttons
func (c *buttonDemoView) newBtn(p *base.Props, s *button.State) *button.B {
	c.buttons = append(c.buttons, button.New(p, s))
	return c.buttons[len(c.buttons)-1]
}
