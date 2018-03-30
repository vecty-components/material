package icon

import (
	"agamigo.io/material/ripple"
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Size string

// I is a vecty-material icon component.
type I struct {
	vecty.Core
	ID            string
	Markup        []vecty.Applyer
	rootElement   *vecty.HTML
	HasRipple     bool
	rippleC       *ripple.R
	Name          string
	SizePX        int
	Inactive      bool
	Dark          bool
	ClassOverride []string
}

// Render implements the vecty.Component interface.
func (c *I) Render() vecty.ComponentOrHTML {
	sizeClass := js.InternalObject(c).Get("SizePX").String()
	switch sizeClass {
	case "undefined", "", "24":
		sizeClass = ""
	default:
		sizeClass = "md-" + sizeClass
	}
	isIconCode := false
	if c.Name != "" && string([]byte(c.Name)[0]) == "&" {
		isIconCode = true
	}
	c.rootElement = elem.Italic(
		vecty.Markup(
			vecty.Markup(c.Markup...),
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
		),
		vecty.If(!isIconCode, vecty.Text(c.Name)),
	)
	return c.rootElement
}

func (c *I) MDCRoot() *base.Base {
	return &base.Base{
		ID:        c.ID,
		Element:   c.rootElement,
		HasRipple: c.HasRipple,
		RippleC:   c.rippleC,
	}
}

func (c *I) Mount() {
	c.MDCRoot().Mount()
}

func (c *I) Unmount() {
	c.MDCRoot().Unmount()
}
