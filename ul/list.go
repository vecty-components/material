package ul

import (
	"log"
	"reflect"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/vecty-components/material/base"
)

type nativeInputer interface {
	NativeInput() (*vecty.HTML, string)
}

// L is a vecty-material list component.
type L struct {
	*base.MDC
	vecty.Core
	Root           vecty.MarkupOrChild     `vecty:"prop"`
	Items          []vecty.ComponentOrHTML `vecty:"prop"`
	Dense          bool                    `vecty:"prop"`
	Avatar         bool                    `vecty:"prop"`
	NonInteractive bool                    `vecty:"prop"`
	GroupSubheader string                  `vecty:"prop"`
	twoLine        bool
}

// Item is a vecty-material list-item component.
type Item struct {
	*base.MDC
	base.KeyedComponent
	vecty.Core
	Root      vecty.MarkupOrChild   `vecty:"prop"`
	Primary   vecty.ComponentOrHTML `vecty:"prop"`
	Secondary vecty.ComponentOrHTML `vecty:"prop"`
	Graphic   vecty.ComponentOrHTML `vecty:"prop"`
	Meta      vecty.ComponentOrHTML `vecty:"prop"`
	Selected  bool                  `vecty:"prop"`
	Activated bool                  `vecty:"prop"`
	Alt       string                `vecty:"prop"`
	K         string                `vecty:"prop"`

	markup *base.LinkMarkup
}

// Group is a vecty-material list-group component.
type Group struct {
	*base.MDC
	vecty.Core
	Root  vecty.MarkupOrChild
	Lists []vecty.ComponentOrHTML
}

type divider struct {
	vecty.Core
}

// Render implements the vecty.Component interface.
func (c *L) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.UnorderedList(c.Root)
	}

	items := make([]vecty.MarkupOrChild, len(c.Items))
	c.twoLine = false
	for i, li := range c.Items {
		switch t := li.(type) {
		case *Item:
			if t.Secondary != nil {
				c.twoLine = true
			}
			items[i] = t
		case *vecty.HTML:
			items[i] = base.RenderStoredChild(t)
		default:
			items[i] = li
		}
	}

	root := elem.UnorderedList(items...)
	vecty.Markup(
		c,
		base.MarkupIfNotNil(rootMarkup),
	).Apply(root)
	return root
}

func (c *L) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
	}

	vecty.Markup(
		vecty.Class("mdc-list"),
		vecty.MarkupIf(c.twoLine,
			vecty.Class("mdc-list--two-line")),
		vecty.MarkupIf(c.Dense,
			vecty.Class("mdc-list--dense")),
		vecty.MarkupIf(c.Avatar,
			vecty.Class("mdc-list--avatar-list")),
		vecty.MarkupIf(c.NonInteractive,
			vecty.Class("mdc-list--non-interactive")),
	).Apply(h)
	c.MDC.RootElement = h
}

// Render implements the vecty.Component interface.
func (c *Item) Render() vecty.ComponentOrHTML {
	c.markup = base.ExtractMarkupFromLink(c.Primary)

	tag := "li"
	if c.markup.Href != "" {
		tag = "a"
	}

	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return vecty.Tag(tag, c.Root)
	}

	graphic := setupGraphicOrMeta(c.Graphic)
	if graphic != nil {
		if g, ok := graphic.(*vecty.HTML); ok {
			vecty.Class("mdc-list-item__graphic").Apply(g)
			vecty.Attribute("role", "presentation").Apply(g)
		}
	}

	meta := setupGraphicOrMeta(c.Meta)
	if meta != nil {
		if g, ok := meta.(*vecty.HTML); ok {
			vecty.Class("mdc-list-item__meta").Apply(g)
			vecty.Attribute("role", "presentation").Apply(g)
		}
	}

	var text vecty.ComponentOrHTML

	switch {
	case c.Secondary != nil:
		text = elem.Span(vecty.Markup(vecty.Class("mdc-list-item__text")),
			c.markup.Child,
			elem.Span(vecty.Markup(
				vecty.Class("mdc-list-item__secondary-text")),
				c.Secondary,
			))
	default:
		text = c.markup.Child
	}

	return vecty.Tag(tag,
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		graphic,
		base.RenderStoredChild(text),
		meta,
	)
}

func (c *Item) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
		fallthrough
	case c.MDC.Component == nil:
		c.MDC.Component = &base.Component{
			Type: base.ComponentType{
				MDCClassName:     "MDCRipple",
				MDCCamelCaseName: "ripple",
			},
		}

		c.MDC.Component.Component().SetState(base.StateMap{})
	}

	vecty.Markup(
		vecty.Class("mdc-list-item"),
		vecty.MarkupIf(c.Selected,
			vecty.Class("mdc-list-item--selected")),
		vecty.MarkupIf(c.Activated,
			vecty.Class("mdc-list-item--activated")),
		vecty.MarkupIf(c.markup.OnClick != nil,
			event.Click(c.markup.OnClick).PreventDefault(),
		),
		vecty.MarkupIf(c.markup.Href != "", prop.Href(c.markup.Href)),
	).Apply(h)
	c.MDC.RootElement = h
}

// Render implements the vecty.Component interface.
func (c *Group) Render() vecty.ComponentOrHTML {
	rootMarkup := base.MarkupOnly(c.Root)
	if c.Root != nil && rootMarkup == nil {
		// User supplied root element.
		return elem.Div(c.Root)
	}

	return elem.Div(
		vecty.Markup(
			c,
			base.MarkupIfNotNil(rootMarkup),
		),
		c.listList(),
	)
}

func (c *Group) Apply(h *vecty.HTML) {
	switch {
	case c.MDC == nil:
		c.MDC = &base.MDC{}
	}

	vecty.Markup(
		vecty.Class("mdc-list-group"),
	).Apply(h)
	c.MDC.RootElement = h
}

type DividerComponent struct {
	vecty.Core
	base.KeyedComponent             //`vecty:"prop"`
	Root                *vecty.HTML `vecty:"prop"` //????????????
}

func NewDividerComponent(h *vecty.HTML) vecty.ComponentOrHTML {
	return &DividerComponent{Root: h}
}

func (dc *DividerComponent) Render() vecty.ComponentOrHTML { //bug caused from here?????
	return base.RenderStoredChild(dc.Root) //package base --> base/tools.go
}

func ListDivider() vecty.ComponentOrHTML {
	d := elem.HorizontalRule(
		vecty.Markup(
			vecty.Class("mdc-list-divider"),
		),
	)
	return NewDividerComponent(d)
}

func ListDividerInset() vecty.ComponentOrHTML {
	d := elem.HorizontalRule(
		vecty.Markup(
			vecty.Class("mdc-list-divider"),
			vecty.Class("mdc-list-divider--inset"),
		),
	)
	return NewDividerComponent(d)
}

type D struct {
	vecty.Core
	key  string
	Root *vecty.HTML
}

func (d *D) Key() interface{} {
	return d.key
}

func NewD(h *vecty.HTML) vecty.ComponentOrHTML {
	D := &D{Root: h}
	D.key = base.Key()
	return D
}

func (d *D) Render() vecty.ComponentOrHTML {
	return d.Root
}

func ItemDivider() vecty.ComponentOrHTML { //bug caused from here?????
	d := elem.ListItem(
		vecty.Markup(
			vecty.Class("mdc-list-divider"),
			vecty.Attribute("role", "separator"),
			vecty.Key(base.Key()), //panic: vecty: all siblings must have keys when using keyed elements
		),
	)
	ndc := NewD(d)
	//ndc := NewDividerComponent(d)
	log.Printf("%v %v %p %+v %v %v %p %+v", "ul/list.ItemDivider()1", "d =", d, d, d.Key().(string), "NewDividerComponent =", ndc, ndc)
	//return NewDividerComponent(d)
	return ndc
	//return d //panic: vecty: all siblings must have keys when using keyed elements
}

func ItemDividerInset() vecty.ComponentOrHTML {
	d := elem.ListItem(
		vecty.Markup(
			vecty.Class("mdc-list-divider"),
			vecty.Class("mdc-list-divider--inset"),
			vecty.Attribute("role", "separator"),
		),
	)
	return NewDividerComponent(d)
}

func (c *Group) listList() vecty.List {
	lists := make(vecty.List, len(c.Lists)*2)
	for _, cList := range c.Lists {
		if list, ok := cList.(*L); ok {
			if list.GroupSubheader != "" {
				lists = append(lists,
					elem.Heading3(vecty.Markup(
						vecty.Class("mdc-list-group__subheader")),
						vecty.Text(list.GroupSubheader)),
				)
			}
			lists = append(lists, list)
		}
	}
	return lists
}

func setupGraphicOrMeta(graphic vecty.ComponentOrHTML) vecty.ComponentOrHTML {

	if graphic != nil && reflect.ValueOf(graphic).Kind().String() != "slice" {
		tag := reflect.ValueOf(graphic).Elem().FieldByName("tag").String()
		if tag != "span" {
			graphic = elem.Span(
				graphic,
			)
		}
	}

	if graphic != nil {
		graphic = base.RenderStoredChild(graphic)
	}

	return graphic
}
