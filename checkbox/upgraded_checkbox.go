package checkbox

import (
	mdccb "agamigo.io/material/checkbox"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

type CB struct {
	*mdccb.CB
	*BasicCB
	started bool
}

func NewUpgraded(id string) CBInterface {
	cb := &CB{}
	cb.CB = &mdccb.CB{}
	cb.BasicCB = NewBasic(id).(*BasicCB)
	return cb
}

func (c *CB) Render() vecty.ComponentOrHTML {
	return render(c)
}

func (c *CB) Checked() bool {
	if c.started {
		return c.CB.Checked
	}
	return c.BasicCB.Checked()
}

func (c *CB) SetChecked(v bool) {
	if c.started {
		c.CB.Checked = v
	}
	c.BasicCB.SetChecked(v)
}

func (c *CB) Disabled() bool {
	if c.started {
		return c.CB.Disabled
	}
	return c.BasicCB.Disabled()
}

func (c *CB) SetDisabled(v bool) {
	if c.started {
		c.CB.Disabled = v
	}
	c.BasicCB.SetDisabled(v)
}

func (c *CB) Indeterminate() bool {
	if c.started {
		return c.CB.Indeterminate
	}
	return c.BasicCB.Indeterminate()
}

func (c *CB) SetIndeterminate(v bool) {
	if c.started {
		c.CB.Indeterminate = v
	}
	c.BasicCB.SetIndeterminate(v)
}

func (c *CB) Value() string {
	if c.started {
		return c.CB.Value
	}
	return c.BasicCB.Value()
}

func (c *CB) SetValue(v string) {
	if c.started {
		c.CB.Value = v
	}
	c.BasicCB.SetValue(v)
}

func (c *CB) Mount() {
	c.BasicCB.Mount()
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

func (c *CB) Unmount() {
	err := c.Stop()
	if err != nil {
		panic(err)
	}
	c.started = false
}
