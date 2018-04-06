package icon

import (
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Size string

// I is a vecty-material icon component.
type I struct {
	*base.MDCRoot
	vecty.Core
	Root          vecty.MarkupOrChild
	Name          string
	SizePX        int
	Inactive      bool
	Dark          bool
	ClassOverride []string
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
			vecty.MarkupIf(rootMarkup != nil, *rootMarkup),
		),
		vecty.If(!isIconCode, vecty.Text(c.Name)),
	)
}

func (c *I) Apply(h *vecty.HTML) {
	sizeClass, isIconCode := c.iconDetails()
	switch {
	case c.MDCRoot == nil:
		c.MDCRoot = &base.MDCRoot{}
	}
	c.MDCRoot.Element = h
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

func (c *I) iconDetails() (sizeClass string, isIconCode bool) {
	sizeClass = js.InternalObject(c).Get("SizePX").String()
	switch sizeClass {
	case "undefined", "", "24":
		sizeClass = ""
	default:
		sizeClass = "md-" + sizeClass
	}
	if c.Name != "" && string([]byte(c.Name)[0]) == "&" {
		isIconCode = true
	}
	return
}
