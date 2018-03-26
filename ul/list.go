package ul

import (
	"agamigo.io/vecty-material/base"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// L is a vecty-material list component.
type L struct {
	*base.Base
	*State
}

type State struct {
	Items          []vecty.ComponentOrHTML `vecty:"prop"`
	Dense          bool
	Avatar         bool
	NonInteractive bool
	ClickHandler   func(thisL *L, thisI *Item, e *vecty.Event)
	GroupSubheader string
}

// Item is a vecty-material list-item component.
type Item struct {
	*base.Base
	*ItemState
}

type ItemState struct {
	Primary      vecty.ComponentOrHTML
	Secondary    vecty.ComponentOrHTML
	Graphic      vecty.ComponentOrHTML
	Meta         vecty.ComponentOrHTML
	Selected     bool
	Activated    bool
	ClickHandler func(i *Item, e *vecty.Event)
	Href         string
}

// Group is a vecty-material list-group component.
type Group struct {
	*base.Base
	*GroupState
}

type GroupState struct {
	Lists []vecty.ComponentOrHTML
}

type divider struct {
	vecty.Core
}

func New(p *base.Props, s *State) *L {
	c := &L{}
	if s == nil {
		s = &State{}
	}
	c.State = s
	c.Base = base.New(p, nil)
	return c
}

func NewItem(p *base.Props, s *ItemState) *Item {
	c := &Item{}
	if s == nil {
		s = &ItemState{}
	}
	c.ItemState = s
	c.Base = base.New(p, nil)
	return c
}

func NewGroup(p *base.Props, s *GroupState) *Group {
	c := &Group{}
	if s == nil {
		s = &GroupState{}
	}
	c.GroupState = s
	c.Base = base.New(p, nil)
	return c
}

// Render implements the vecty.Component interface.
func (c *L) Render() vecty.ComponentOrHTML {
	items := make([]vecty.MarkupOrChild, len(c.Items))
	twoLine := false
	for i, li := range c.Items {
		switch t := li.(type) {
		case *Item:
			if t.Secondary != nil {
				twoLine = true
			}
			items[i] = t
		case *vecty.HTML:
			items[i] = base.RenderStoredChild(t)
		default:
			items[i] = li
		}
	}
	h := elem.UnorderedList(items...)
	vecty.Markup(
		vecty.Markup(c.Props.Markup...),
		vecty.Class("mdc-list"),
		vecty.MarkupIf(twoLine,
			vecty.Class("mdc-list--two-line")),
		vecty.MarkupIf(c.Dense,
			vecty.Class("mdc-list--dense")),
		vecty.MarkupIf(c.Avatar,
			vecty.Class("mdc-list--avatar-list")),
		vecty.MarkupIf(c.NonInteractive,
			vecty.Class("mdc-list--non-interactive")),
	).Apply(h)
	return h
}

// Render implements the vecty.Component interface.
func (c *Item) Render() vecty.ComponentOrHTML {
	tag := "li"
	if c.Href != "" {
		tag = "a"
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
			c.Primary,
			elem.Span(vecty.Markup(
				vecty.Class("mdc-list-item__secondary-text")),
				c.Secondary,
			))
	default:
		text = c.Primary
	}
	return c.Base.Render(vecty.Tag(tag,
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.Class("mdc-list-item"),
			vecty.MarkupIf(c.Selected,
				vecty.Class("mdc-list-item--selected")),
			vecty.MarkupIf(c.Activated,
				vecty.Class("mdc-list-item--activated")),
			vecty.MarkupIf(c.ClickHandler != nil,
				event.Click(c.wrapClickHandler()),
			),
			vecty.MarkupIf(c.Href != "", prop.Href(c.Href)),
		),
		graphic,
		base.RenderStoredChild(text),
		meta,
	))
}

// Render implements the vecty.Component interface.
func (c *Group) Render() vecty.ComponentOrHTML {
	return c.Base.Render(elem.Div(
		vecty.Markup(
			vecty.Markup(c.Props.Markup...),
			vecty.Class("mdc-list-group"),
		),
		c.listList(),
	))
}

func ListDivider() vecty.ComponentOrHTML {
	d := elem.HorizontalRule(
		vecty.Markup(
			vecty.Class("mdc-list-divider"),
		),
	)
	return base.RenderStoredChild(d)
}

func ListDividerInset() vecty.ComponentOrHTML {
	d := elem.HorizontalRule(
		vecty.Markup(
			vecty.Class("mdc-list-divider"),
			vecty.Class("mdc-list-divider--inset"),
		),
	)
	return base.RenderStoredChild(d)
}

func ItemDivider() vecty.ComponentOrHTML {
	d := elem.ListItem(
		vecty.Markup(
			vecty.Class("mdc-list-divider"),
			vecty.Attribute("role", "separator"),
		),
	)
	return base.RenderStoredChild(d)
}

func ItemDividerInset() vecty.ComponentOrHTML {
	d := elem.ListItem(
		vecty.Markup(
			vecty.Class("mdc-list-divider"),
			vecty.Class("mdc-list-divider--inset"),
			vecty.Attribute("role", "separator"),
		),
	)
	return base.RenderStoredChild(d)
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

func (c *L) wrapClickHandler() func(i *Item, e *vecty.Event) {
	return func(i *Item, e *vecty.Event) {
		c.ClickHandler(c, i, e)
	}
}

func (c *Item) wrapClickHandler() func(e *vecty.Event) {
	return func(e *vecty.Event) {
		c.ClickHandler(c, e)
	}
}

func setupGraphicOrMeta(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
	var graphic vecty.ComponentOrHTML
	if c != nil {
		graphic = c
		switch t := c.(type) {
		case vecty.Component:
			if h, ok := t.Render().(*vecty.HTML); ok {
				graphic = h
			}
		}
		if js.InternalObject(graphic).Get("tag").String() != "img" {
			graphic = elem.Span(graphic)
		}
	}
	return graphic
}
