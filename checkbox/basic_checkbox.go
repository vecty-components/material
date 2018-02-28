package checkbox

import (
	"errors"

	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

type basicCB struct {
	vecty.Core
	id            string
	element       *js.Object
	customClasses vecty.ClassMap
	checked       bool
	disabled      bool
	indeterminate bool
	value         string
}

func NewBasic(id string) CB {
	return &basicCB{
		id:            id,
		customClasses: make(vecty.ClassMap, 0),
	}
}

func (c *basicCB) Render() vecty.ComponentOrHTML {
	return render(c)
}

func (c *basicCB) Checked() bool {
	v, err := c.getInputBoolProp("checked")
	if err != nil {
		return c.checked
	}
	return v
}

func (c *basicCB) SetChecked(v bool) {
	c.checked = v
	_ = c.setInputProp("checked", v)
	// if err != nil {
	// 	panic(err)
	// }
}

func (c *basicCB) Disabled() bool {
	v, err := c.getInputBoolProp("disabled")
	if err != nil {
		return c.disabled
	}
	return v
}

func (c *basicCB) SetDisabled(v bool) {
	c.disabled = v
	_ = c.setInputProp("disabled", v)
	// if err != nil {
	// 	panic(err)
	// }
}

func (c *basicCB) Indeterminate() bool {
	v, err := c.getInputBoolProp("indeterminate")
	if err != nil {
		return c.indeterminate
	}
	return v
}

func (c *basicCB) SetIndeterminate(v bool) {
	c.indeterminate = v
	_ = c.setInputProp("indeterminate", v)
	// if err != nil {
	// 	panic(err)
	// }
}

func (c *basicCB) Value() string {
	v, err := c.getInputStringProp("value")
	if err != nil {
		return c.value
	}
	return v
}

func (c *basicCB) SetValue(v string) {
	c.value = v
	_ = c.setInputProp("checked", v)
	// if err != nil {
	// 	panic(err)
	// }
}

func (c *basicCB) ID() string {
	return c.id
}

func (c *basicCB) Element() *js.Object {
	return c.element
}

func (c *basicCB) getInputProp(prop string) (value *js.Object, err error) {
	gojs.CatchException(&err)
	e := js.Global.Get("document").Call("getElementById", c.ID())
	if e == nil || e == js.Undefined {
		return nil, errors.New("Unable to find Input Element \"" + c.ID() + "\".")
	}
	return e.Get(prop), err
}

func (c *basicCB) getInputStringProp(prop string) (value string, err error) {
	v, err := c.getInputProp(prop)
	if err != nil || v == nil || v == js.Undefined {
		return "", err
	}
	return v.String(), err
}

func (c *basicCB) getInputBoolProp(prop string) (value bool, err error) {
	v, err := c.getInputProp(prop)
	if err != nil || v == nil || v == js.Undefined {
		return false, err
	}
	return v.Bool(), err
}

func (c *basicCB) setInputProp(prop string, value interface{}) (err error) {
	gojs.CatchException(&err)
	e := js.Global.Get("document").Call("getElementById", c.ID())
	if e == nil || e == js.Undefined {
		return errors.New("Unable to find Input Element \"" + c.ID() + "\".")
	}
	e.Set(prop, value)
	return err
}

func (c *basicCB) Mount() {
	ie := js.Global.Get("document").Call("getElementById", c.ID())
	c.element = ie.Get("parentNode")
}

func (c *basicCB) Unmount() {
}

func (c *basicCB) AddClass(class string) {
	c.customClasses[class] = true
}

func (c *basicCB) DelClass(class string) {
	c.customClasses[class] = false
}

func (c *basicCB) getClasses() vecty.ClassMap {
	return c.customClasses
}
