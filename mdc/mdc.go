package mdc

import (
	"github.com/gopherjs/gopherjs/js"
)

type MDCName int

const (
	MDCNameCustom MDCName = iota
	MDCNameBase
	MDCNameCheckbox
	MDCNameFormField
)

type MDCStatus int

const (
	MDCStatusUninitialized MDCStatus = iota
	MDCStatusStopped
	MDCStatusRunning
)

var (
	nextID = 1
	mdc    = js.Global.Get("mdc")
)

type Adaptable interface {
	Name() MDCName
	SetObject(o *js.Object)
	GetObject() *js.Object
}

type Componenter interface {
	Adaptable
	ID() int
	setID(id int)
	Status() MDCStatus
	setStatus(s MDCStatus)
}

type Component struct {
	*js.Object
	name   MDCName
	id     int
	status MDCStatus
}

func New(n MDCName) *Component {
	c := &Component{}
	c.name = n
	o := makeMDComponent(c)
	if o == nil || o == js.Undefined {
		panic("Creating " + c.Name().classString() +
			" failed, object nil or undefined")
	}
	c.SetObject(o)
	c.setStatus(MDCStatusStopped)
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

func (c *Component) Status() MDCStatus {
	return c.status
}

func (c *Component) setStatus(s MDCStatus) {
	c.status = s
}

func (c *Component) Name() MDCName {
	return c.name
}

func (c *Component) GetObject() *js.Object {
	return c.Object
}

func (c *Component) SetObject(o *js.Object) {
	c.Object = o
}

func (n MDCName) componentString() string {
	switch n {
	case MDCNameCheckbox:
		return "MDCCheckbox"
	case MDCNameFormField:
		return "MDCFormField"
	}

	panic("Failed to convert MDCName to component string.")
	return ""
}

func (n MDCName) classString() string {
	switch n {
	case MDCNameCheckbox:
		return "checkbox"
	case MDCNameFormField:
		return "form-field"
	}

	panic("Failed to convert MDCName to class string.")
	return ""
}

func makeMDComponent(c Adaptable) *js.Object {
	switch c.Name() {
	case MDCNameCheckbox:
		return mdc.Get("checkbox").Get(c.Name().componentString())
	}
	return nil
}

func Start(c Componenter) {
	switch c.Name() {
	case MDCNameCheckbox:
		StartWith(c, "div.mdc-"+string(c.Name().classString()))
	}
}

func StartWith(c Componenter, querySelector string) {
	if c.Status() == MDCStatusRunning {
		return
	}
	if c.Status() != MDCStatusStopped {
		panic("Attempted to run Start() an uninitialized component. Use mdc.New()")
	}

	e := js.Global.Get("document").Call("querySelector", querySelector)
	c.SetObject(c.GetObject().New(e))
	c.setStatus(MDCStatusRunning)
}

func Stop(c Componenter) {
	if c.Status() == MDCStatusStopped {
		println(c.Name().classString())
		print("Attempted to stop already stopped component: ")
		return
	}

	if c.Status() != MDCStatusRunning {
		println(c.Name().classString())
		panic("Attempted to run Stop() an uninitialized component. Use mdc.New()")
	}

	c.GetObject().Call("destroy")
}
