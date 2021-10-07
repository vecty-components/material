package datatable

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/checkbox"
)

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
	Root vecty.MarkupOrChild `vecty:"prop"`
	Head *R                  `vecty:"prop"`
	Rows []*R                `vecty:"prop"`
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

	if c.MDC != nil && c.MDC.Component != nil {
		component := c.MDC.Component.(*base.Component)
		go func() {
			time.Sleep(50 * time.Millisecond)
			component.Call("layout")
		}()
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
