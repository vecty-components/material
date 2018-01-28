package mdc

import (
	"github.com/gopherjs/gopherjs/js"
)

type MDCName string

const (
	MDCNameCustom   MDCName = "custom"
	MDCNameBase             = "base"
	MDCNameCheckbox         = "checkbox"
)

type MDCStatus int

const (
	MDCStatusUninitialized MDCStatus = iota
	MDCStatusStopped
	MDCStatusRunning
)

var (
	nextID = 1
	// activeCs map[MDCName][]int
	mdc = js.Global.Get("mdc")
)

type Adaptable interface {
	MDCName() MDCName
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
	id     int
	status MDCStatus
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

func New(c Componenter) {
	o := makeMDComponent(c)
	if o == nil || o == js.Undefined {
		panic("Creating " + c.MDCName() +
			" failed, object nil or undefined")
	}
	c.SetObject(o)
	c.setStatus(MDCStatusStopped)
	c.setID(nextID)
	nextID = nextID + 1
}

func makeMDComponent(c Adaptable) *js.Object {
	switch c.MDCName() {
	case MDCNameCheckbox:
		return mdc.Get("checkbox").Get("MDCCheckbox")
	}
	return nil
}

func Start(c Componenter) {
	switch c.MDCName() {
	case MDCNameCheckbox:
		StartWith(c, "div.mdc-"+string(c.MDCName()))
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

	// activeCs[c.MDCName()] = append [c.MDCName()]c.ID()
}

func Stop(c Componenter) {
	if c.Status() == MDCStatusStopped {
		print("Attempted to stop already stopped component: ")
		println(c.MDCName()) //+ " #" + c.ID())
		return
	}

	if c.Status() != MDCStatusRunning {
		panic("Attempted to run Stop() an uninitialized component. Use mdc.New()")
	}

	c.GetObject().Call("destroy")
}
