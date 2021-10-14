package datatable

import (
	"syscall/js"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
	"github.com/vecty-components/material/checkbox"
)

type R struct {
	vecty.Core
	Cells []*C   `vecty:"prop"`
	K     string `vecty:"prop"`
}

func (c *R) Render() vecty.ComponentOrHTML {
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

func (c *R) Key() interface{} {
	if c.K == "" {
		c.K = base.Key()
	}

	return c.K
}

type C struct {
	Label vecty.ComponentOrHTML
	Root  vecty.MarkupList
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
		c.Root,
		label,
	)
}

// TB is a vecty-material datatable component.
type DT struct {
	*base.MDC
	vecty.Core
	Root vecty.MarkupOrChild `vecty:"prop"`
	Head []*C                `vecty:"prop"`
	Rows []*R                `vecty:"prop"`
}

func (c *DT) renderHead() *vecty.HTML {
	cells := make([]vecty.MarkupOrChild, len(c.Head))
	for i, c := range c.Head {
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

// Render implements the vecty.Component interface.
func (c *DT) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	rows := make([]vecty.MarkupOrChild, len(c.Rows))
	for i, row := range c.Rows {
		rows[i] = row
	}

	id := base.Key()
	if c.MDC != nil && c.MDC.Component != nil {
		component := c.MDC.Component.(*base.Component)
		go func() {
			for {
				time.Sleep(15 * time.Millisecond)
				if !js.Global().Get("document").
					Call("getElementById", id).IsUndefined() {
					break
				}
			}
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
				c.renderHead(),
			),
			elem.TableBody(
				append([]vecty.MarkupOrChild{
					vecty.Markup(
						vecty.Class("mdc-data-table__content"),
					),
				}, rows...)...,
			),
		),
		elem.Div(
			vecty.Markup(
				prop.ID(id),
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
