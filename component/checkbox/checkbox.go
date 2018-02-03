package checkbox

import (
	"agamigo.io/material/component"
)

const (
	html = `<div class="mdc-checkbox">
				<input class="mdc-checkbox__native-control" type="checkbox">
			</div>`
)

type CB interface {
	component.C
	component.HTMLElementer
	State() StateType
	SetState(s StateType)
	Value() string
	SetValue(v string)
}

type checkbox struct {
	component.C
	html string
}

func New() (c CB, err error) {
	newC, err := component.New(component.Checkbox)
	if err != nil {
		return nil, err
	}
	return &checkbox{newC, html}, err
}

func (c *checkbox) State() StateType {
	s := UNKNOWN
	checked := c.GetObject().Get("checked").Bool()

	switch {
	case c.GetObject().Get("indeterminate").Bool():
		s = INDETERMINATE
	case !checked:
		s = UNCHECKED
	case checked:
		s = CHECKED
	}

	if c.GetObject().Get("disabled").Bool() {
		s = s + DISABLED
	}

	if s == UNKNOWN {
		println("Warning: State of input is UNKNOWN.")
	}

	return s
}

func (c *checkbox) SetState(s StateType) {
	switch s {
	case UNKNOWN:
		println("SetState failed, invalid state given.")
	case INDETERMINATE, INDETERMINATE_DISABLED:
		c.GetObject().Set("indeterminate", true)
	case UNCHECKED, UNCHECKED_DISABLED:
		c.GetObject().Set("checked", false)
		c.GetObject().Set("indeterminate", false)
	case CHECKED, CHECKED_DISABLED:
		c.GetObject().Set("checked", true)
		c.GetObject().Set("indeterminate", false)
	}

	if s%2 != 0 {
		c.GetObject().Set("disabled", true)
		return
	}

	c.GetObject().Set("disabled", false)
}

func (c *checkbox) Value() string {
	return c.GetObject().Get("value").String()
}

func (c *checkbox) SetValue(v string) {
	c.GetObject().Set("value", v)
}

func (c *checkbox) HTML() string {
	return c.html
}

func (c *checkbox) SetHTML(html string) {
	c.html = html
}
