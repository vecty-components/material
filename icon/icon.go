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
	*base.Base
	*State
}

type State struct {
	Name          string
	SizePX        int
	Inactive      bool
	Dark          bool
	ClassOverride string
}

func New(p *base.Props, s *State) *I {
	c := &I{}
	if s == nil {
		s = &State{}
	}
	c.State = s
	c.Base = base.New(p, nil)
	return c
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
	return c.Base.Render(elem.Italic(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.MarkupIf(c.ClassOverride == "",
				vecty.Class("material-icons"),
			),
			vecty.MarkupIf(c.ClassOverride != "",
				vecty.Class(c.ClassOverride),
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
	))
}
