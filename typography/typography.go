package typography

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func Subtitle1(markup ...vecty.MarkupOrChild) *vecty.HTML {
	return elem.Heading3(
		append(
			[]vecty.MarkupOrChild{
				vecty.Markup(
					vecty.Class("mdc-typography--subtitle1"),
				),
			},
			markup...,
		)...,
	)
}
