package checkbox

import (
	"errors"

	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

type BasicCB struct {
	vecty.Core
	id            string
	element       *js.Object
	customClasses vecty.ClassMap
	checked       bool
	disabled      bool
	indeterminate bool
	value         string
}

func NewBasic(id string) CBInterface {
	return &BasicCB{
		id:            id,
		customClasses: make(vecty.ClassMap, 0),
	}
}

func (c *BasicCB) Render() vecty.ComponentOrHTML {
	return render(c)
}

func (c *BasicCB) Checked() bool {
	v, err := c.getInputBoolProp("checked")
	if err != nil {
		return c.checked
	}
	return v
}

func (c *BasicCB) SetChecked(v bool) {
	c.checked = v
	_ = c.setInputProp("checked", v)
	// if err != nil {
	// 	panic(err)
	// }
}

func (c *BasicCB) Disabled() bool {
	v, err := c.getInputBoolProp("disabled")
	if err != nil {
		return c.disabled
	}
	return v
}

func (c *BasicCB) SetDisabled(v bool) {
	c.disabled = v
	_ = c.setInputProp("disabled", v)
	// if err != nil {
	// 	panic(err)
	// }
}

func (c *BasicCB) Indeterminate() bool {
	v, err := c.getInputBoolProp("indeterminate")
	if err != nil {
		return c.indeterminate
	}
	return v
}

func (c *BasicCB) SetIndeterminate(v bool) {
	c.indeterminate = v
	_ = c.setInputProp("indeterminate", v)
	// if err != nil {
	// 	panic(err)
	// }
}

func (c *BasicCB) Value() string {
	v, err := c.getInputStringProp("value")
	if err != nil {
		return c.value
	}
	return v
}

func (c *BasicCB) SetValue(v string) {
	c.value = v
	_ = c.setInputProp("checked", v)
	// if err != nil {
	// 	panic(err)
	// }
}

func (c *BasicCB) ID() string {
	return c.id
}

func (c *BasicCB) Element() *js.Object {
	return c.element
}

func (c *BasicCB) getInputProp(prop string) (value *js.Object, err error) {
	gojs.CatchException(&err)
	e := js.Global.Get("document").Call("getElementById", c.ID())
	if e == nil || e == js.Undefined {
		return nil, errors.New("Unable to find Input Element \"" + c.ID() + "\".")
	}
	return e.Get(prop), err
}

func (c *BasicCB) getInputStringProp(prop string) (value string, err error) {
	v, err := c.getInputProp(prop)
	if err != nil || v == nil || v == js.Undefined {
		return "", err
	}
	return v.String(), err
}

func (c *BasicCB) getInputBoolProp(prop string) (value bool, err error) {
	v, err := c.getInputProp(prop)
	if err != nil || v == nil || v == js.Undefined {
		return false, err
	}
	return v.Bool(), err
}

func (c *BasicCB) setInputProp(prop string, value interface{}) (err error) {
	gojs.CatchException(&err)
	e := js.Global.Get("document").Call("getElementById", c.ID())
	if e == nil || e == js.Undefined {
		return errors.New("Unable to find Input Element \"" + c.ID() + "\".")
	}
	e.Set(prop, value)
	return err
}

func (c *BasicCB) Mount() {
	ie := js.Global.Get("document").Call("getElementById", c.ID())
	c.element = ie.Get("parentNode")
}

func (c *BasicCB) Unmount() {
}

func (c *BasicCB) AddClass(class string) {
	c.customClasses[class] = true
}

func (c *BasicCB) DelClass(class string) {
	c.customClasses[class] = false
}

func (c *BasicCB) getClasses() vecty.ClassMap {
	return c.customClasses
}
