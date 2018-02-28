package checkbox

import (
	mdccb "agamigo.io/material/checkbox"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

type upgradedCB struct {
	*mdccb.CB
	*basicCB
	started bool
}

func NewUpgraded(id string) CB {
	cb := &upgradedCB{}
	cb.CB = &mdccb.CB{}
	cb.basicCB = NewBasic(id).(*basicCB)
	return cb
}

func (c *upgradedCB) Render() vecty.ComponentOrHTML {
	return render(c)
}

func (c *upgradedCB) Checked() bool {
	if c.started {
		return c.CB.Checked
	}
	return c.basicCB.Checked()
}

func (c *upgradedCB) SetChecked(v bool) {
	if c.started {
		c.CB.Checked = v
	}
	c.basicCB.SetChecked(v)
}

func (c *upgradedCB) Disabled() bool {
	if c.started {
		return c.CB.Disabled
	}
	return c.basicCB.Disabled()
}

func (c *upgradedCB) SetDisabled(v bool) {
	if c.started {
		c.CB.Disabled = v
	}
	c.basicCB.SetDisabled(v)
}

func (c *upgradedCB) Indeterminate() bool {
	if c.started {
		return c.CB.Indeterminate
	}
	return c.basicCB.Indeterminate()
}

func (c *upgradedCB) SetIndeterminate(v bool) {
	if c.started {
		c.CB.Indeterminate = v
	}
	c.basicCB.SetIndeterminate(v)
}

func (c *upgradedCB) Value() string {
	if c.started {
		return c.CB.Value
	}
	return c.basicCB.Value()
}

func (c *upgradedCB) SetValue(v string) {
	if c.started {
		c.CB.Value = v
	}
	c.basicCB.SetValue(v)
}

func (c *upgradedCB) Mount() {
	c.basicCB.Mount()
	e := c.Element()
	if e == nil || e == js.Undefined {
		panic("Element() is null while mounting upgradedCB.")
	}
	err := c.Start(e)
	if err != nil {
		panic(err)
	}
	c.started = true
}

func (c *upgradedCB) Unmount() {
	err := c.Stop()
	if err != nil {
		panic(err)
	}
	c.started = false
}
