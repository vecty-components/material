package icon

import (
	"strconv"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/vecty-components/material/base"
)

type Size string

// I is a vecty-material icon component.
type I struct {
	*base.MDC
	vecty.Core
	Root          vecty.MarkupOrChild `vecty:"prop"`
	Name          string              `vecty:"prop"`
	SizePX        int                 `vecty:"prop"`
	Inactive      bool                `vecty:"prop"`
	Dark          bool                `vecty:"prop"`
	ClassOverride []string            `vecty:"prop"`
}

// Render implements the vecty.Component interface.
func (c *I) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	_, isIconCode := c.iconDetails()

	return elem.Italic(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		vecty.If(!isIconCode, vecty.Text(c.Name)),
	)
}

func (c *I) Apply(h *vecty.HTML) {
	sizeClass, isIconCode := c.iconDetails()
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
	}
	c.MDC.RootElement = h
	vecty.Markup(
		vecty.MarkupIf(c.ClassOverride == nil,
			vecty.Class("material-icons"),
		),
		vecty.MarkupIf(c.ClassOverride != nil,
			vecty.Class(c.ClassOverride...),
		),
		vecty.MarkupIf(c.Inactive,
			vecty.Class("md-inactive"),
		),
		vecty.MarkupIf(c.Dark,
			vecty.Class("md-dark"),
		),
		vecty.MarkupIf(sizeClass != "",
			vecty.Class(sizeClass),
		),
		vecty.MarkupIf(isIconCode,
			vecty.UnsafeHTML(c.Name),
		),
	).Apply(h)
}

func (c *I) iconDetails() (string, bool) {
	isIconCode := false
	sizeClass := strconv.Itoa(c.SizePX)
	switch sizeClass {
	case "undefined", "", "24":
		sizeClass = ""
	default:
		sizeClass = "md-" + sizeClass
	}
	if c.Name != "" && string([]byte(c.Name)[0]) == "&" {
		isIconCode = true
	}
	return sizeClass, isIconCode
}
