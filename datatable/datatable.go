package datatable

import (
	"fmt"
	"math/rand"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/checkbox"
)

type fakeComponent struct {
	vecty.Core
	Id    string `vecty:"prop"`
	mfunc func()
}

func newFakeComponent(mfunc func()) *fakeComponent {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	n := 10
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return &fakeComponent{
		Id:    string(s),
		mfunc: mfunc,
	}
}

func (fc *fakeComponent) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			prop.ID(fc.Id),
		),
	)
}

func (fc *fakeComponent) Mount() {
	fc.mfunc()
}

type R struct {
	Cells []*C
}

func (c *R) renderHead() *vecty.HTML {
	cells := make([]vecty.MarkupOrChild, len(c.Cells))
	for i, c := range c.Cells {
		cells[i] = c.renderHead()
	}

	return elem.TableRow(
		append([]vecty.MarkupOrChild{
			vecty.Markup(
				vecty.Class("mdc-data-table__header-row"),
			),
		}, cells...)...,
	)
}

func (c *R) renderRow() *vecty.HTML {
	cells := make([]vecty.MarkupOrChild, len(c.Cells))
	for i, c := range c.Cells {
		cells[i] = c.renderRow()
	}

	return elem.TableRow(
		append([]vecty.MarkupOrChild{
			vecty.Markup(
				vecty.Class("mdc-data-table__row"),
			),
		}, cells...)...,
	)
}

type C struct {
	Label vecty.ComponentOrHTML
}

func (c *C) renderHead() *vecty.HTML {
	label := c.Label
	cb, ok := c.Label.(*checkbox.CB)
	if ok {
		label = cb.Render()
		vecty.Markup(
			vecty.Class("mdc-data-table__header-row-checkbox"),
		).Apply(label.(*vecty.HTML))
	}

	return elem.TableHeader(
		vecty.Markup(
			vecty.Class("mdc-data-table__header-cell"),
			vecty.MarkupIf(
				ok, vecty.Class("mdc-data-table__header-cell--checkbox"),
			),
			vecty.Attribute("role", "columnheader"),
			vecty.Attribute("scope", "col"),
		),
		label,
	)
}

func (c *C) renderRow() *vecty.HTML {
	label := c.Label
	cb, ok := c.Label.(*checkbox.CB)
	if ok {
		label = cb.Render()
		vecty.Markup(
			vecty.Class("mdc-data-table__row-checkbox"),
		).Apply(label.(*vecty.HTML))
	}

	return elem.TableData(
		vecty.Markup(
			vecty.MarkupIf(
				ok, vecty.Class("mdc-data-table__cell--checkbox"),
			),
			vecty.Class("mdc-data-table__cell"),
			vecty.Attribute("scope", "row"),
		),
		label,
	)
}

// TB is a vecty-material datatable component.
type DT struct {
	*base.MDC
	vecty.Core
	Root vecty.MarkupOrChild
	Head *R
	Rows []*R
}

// Render implements the vecty.Component interface.
func (c *DT) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	rows := make([]vecty.MarkupOrChild, len(c.Rows))
	for i, row := range c.Rows {
		rows[i] = row.renderRow()
	}

	// Built in root element.
	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		elem.Table(
			vecty.Markup(
				vecty.Class("mdc-data-table__table"),
			),
			elem.TableHead(
				c.Head.renderHead(),
			),
			elem.TableBody(
				append([]vecty.MarkupOrChild{
					vecty.Markup(
						vecty.Class("mdc-data-table__content"),
					),
				}, rows...)...,
			),
		),
		newFakeComponent(func() {
			fmt.Println("call layout")
			c.Component.Component().Call("layout")
		}),
	)
}

func (c *DT) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCDataTable",
				MDCCamelCaseName: "dataTable",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{})
	}
	c.MDC.RootElement = h
	vecty.Markup(
		vecty.Class("mdc-data-table"),
	).Apply(h)
}
