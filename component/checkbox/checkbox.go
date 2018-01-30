package checkbox

import "agamigo.io/material/component"

type StateType int

const (
	// Unset state is zero
	UNKNOWN StateType = iota
	DISABLED
	// Enabled states are even, disabled are odd
	UNCHECKED
	UNCHECKED_DISABLED
	CHECKED
	CHECKED_DISABLED
	INDETERMINATE
	INDETERMINATE_DISABLED
)

type C interface {
	component.C
	State() StateType
	SetState(s StateType)
	Value() string
	SetValue(v string)
}

type c struct {
	component.C
}

func New() C {
	return &c{
		component.New(component.Checkbox),
	}
}

func (c *c) State() StateType {
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

func (c *c) SetState(s StateType) {
	print("SetState called with:")
	print(s)
	switch s {
	case UNKNOWN:
		panic("SetState failed, invalid state given.")
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

func (c *c) Value() string {
	return c.GetObject().Get("value").String()
}

func (c *c) SetValue(v string) {
	c.GetObject().Set("value", v)
}
