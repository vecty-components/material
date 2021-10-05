package datatable

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-material/material/base"
	"github.com/vecty-material/material/components/datatable"
)

type R struct {
	vecty.Core
	Cells []*C

	head bool
}

func (c *R) Render() vecty.ComponentOrHTML {
	return elem.TableRow(
		vecty.Markup(
			vecty.MarkupIf(
				c.head, vecty.Class("mdc-data-table__header-row"),
			),
			vecty.MarkupIf(
				!c.head, vecty.Class("mdc-data-table__row"),
			),
		),
		// c.Cells...,
	)
}

type C struct {
	vecty.Core
}

func (c *C) Render() vecty.ComponentOrHTML {
	return nil
}

// TB is a vecty-material datatable component.
type TB struct {
	*base.MDC
	vecty.Core
	Root vecty.MarkupOrChild
	Head *R
	Rows []*R
}

// Render implements the vecty.Component interface.
func (c *TB) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}
	c.Head.head = true

	rows := make([]vecty.MarkupOrChild, len(c.Rows))
	for i, row := range c.Rows {
		row.head = false
		rows[i] = row
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
				c.Head,
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

func (c *TB) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = datatable.New()
	}
	c.MDC.RootElement = h
	vecty.Markup(
		vecty.Class("mdc-data-table"),
	).Apply(h)
}
