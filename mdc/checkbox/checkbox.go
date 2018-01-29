package checkbox

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

func (c *C) State() StateType {
	s := UNKNOWN
	checked := c.Get("checked").Bool()
	switch {
	case c.Get("indeterminate").Bool():
		s = INDETERMINATE
	case !checked:
		s = UNCHECKED
	case checked:
		s = CHECKED
	}

	if c.Get("disabled").Bool() {
		s = s + DISABLED
	}

	if s == UNKNOWN {
		println("Warning: State of input is UNKNOWN.")
	}

	return s
}

func (c *C) SetState(s StateType) {
	print("SetState called with:")
	print(s)
	switch s {
	case UNKNOWN:
		panic("SetState failed, invalid state given.")
	case INDETERMINATE, INDETERMINATE_DISABLED:
		c.Set("indeterminate", true)
	case UNCHECKED, UNCHECKED_DISABLED:
		c.Set("checked", false)
		c.Set("indeterminate", false)
	case CHECKED, CHECKED_DISABLED:
		c.Set("checked", true)
		c.Set("indeterminate", false)
	}

	if s%2 != 0 {
		c.Set("disabled", true)
		return
	}

	c.Set("disabled", false)
}

func (c *C) Value() string {
	return c.Get("value").String()
}

func (c *C) SetValue(v string) {
	c.Set("value", v)
}
