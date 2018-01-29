package mdc

import (
	"github.com/gopherjs/gopherjs/js"
)

type ComponentName int

const (
	Custom ComponentName = iota
	Checkbox
	FormField
)

type ComponentStatus int

const (
	Uninitialized ComponentStatus = iota
	Stopped
	Running
)

var (
	nextID = 1
	mdc    = js.Global.Get("mdc")
)

// type Componenter interface {
// 	Name() ComponentName
// 	SetObject(o *js.Object)
// 	GetObject() *js.Object
// 	ID() int
// 	setID(id int)
// 	Status() ComponentStatus
// 	setStatus(s ComponentStatus)
// }

type Component struct {
	*js.Object
	name   ComponentName
	id     int
	status ComponentStatus
}

func New(n ComponentName) *Component {
	c := &Component{}
	c.name = n
	o := makeMDComponent(c)
	if o == nil || o == js.Undefined {
		panic("Creating " + c.Name().classString() +
			" failed, object nil or undefined")
	}
	c.SetObject(o)
	c.setStatus(Stopped)
	c.setID(nextID)
	nextID = nextID + 1
	return c
}

func (c *Component) ID() int {
	return c.id
}

func (c *Component) setID(id int) {
	c.id = id
}

func (c *Component) Status() ComponentStatus {
	return c.status
}

func (c *Component) setStatus(s ComponentStatus) {
	c.status = s
}

func (c *Component) Name() ComponentName {
	return c.name
}

func (c *Component) GetObject() *js.Object {
	return c.Object
}

func (c *Component) SetObject(o *js.Object) {
	c.Object = o
}

func (n ComponentName) componentString() string {
	switch n {
	case Checkbox:
		return "MDCCheckbox"
	case FormField:
		return "MDCFormField"
	}

	panic("Failed to convert MDCName to component string.")
	return ""
}

func (n ComponentName) classString() string {
	switch n {
	case Checkbox:
		return "checkbox"
	case FormField:
		return "form-field"
	}

	panic("Failed to convert MDCName to class string.")
	return ""
}

func makeMDComponent(c *Component) *js.Object {
	switch c.Name() {
	case Checkbox:
		return mdc.Get("checkbox").Get(c.Name().componentString())
	}
	return nil
}

func (c *Component) Start() {
	switch c.Name() {
	case Checkbox:
		c.StartWith("div.mdc-" + string(c.Name().classString()))
	}
}

func (c *Component) StartWith(querySelector string) {
	if c.Status() == Running {
		return
	}
	if c.Status() != Stopped {
		panic("Attempted to run Start() an uninitialized component. Use mdc.New()")
	}

	e := js.Global.Get("document").Call("querySelector", querySelector)
	c.SetObject(c.GetObject().New(e))
	c.setStatus(Running)
}

func (c *Component) Stop() {
	if c.Status() == Stopped {
		println(c.Name().classString())
		print("Attempted to stop already stopped component: ")
		return
	}

	if c.Status() != Running {
		println(c.Name().classString())
		panic("Attempted to run Stop() an uninitialized component. Use mdc.New()")
	}

	c.GetObject().Call("destroy")
}
