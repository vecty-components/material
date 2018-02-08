package component // import "agamigo.io/material/component"

import (
	"errors"

	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

// StatusType holds a component's lifecycle status.
type StatusType int

const (
	// An Uninitialized component has not been associated with the MDC library
	// yet. This package does not provide a way to access an Uninitialized
	// component.
	Uninitialized StatusType = iota

	// A Stopped component has been associated with a JS Object constructed from
	// a MDC class. New() returns a Stopped component, and Stop() will stop a
	// Running component.
	Stopped

	// A Running component has had its underlying MDC init() method called,
	// which attaches the component to a specific HTMLElement in the DOM. It is
	// ready to be used.
	Running
)

// C is the base material component type
type C struct {
	mdc    *js.Object
	name   Type
	status StatusType
}

// New creates a material component of Type t. It assumes the MDC library and
// resulting component will live in the js.Global scope.
func New(t Type) (mdcComponent *C, err error) {
	defer gojs.CatchException(&err)
	c, err := NewWith(t, js.Global)
	return c, err
}

// NewWith is like New(), with added option of specifying a *js.Object to store
// the component. This is primarily intended for use in tests where we may want
// to emulate a DOM somewhere other than Node's global scope.
func NewWith(t Type, dom *js.Object) (mdcComponent *C, err error) {
	defer gojs.CatchException(&err)
	c := &C{}
	c.mdc, err = makeMDComponent(t, dom)
	if err != nil {
		return nil, err
	}

	c.name = t
	c.status = Stopped
	return c, err
}

// String returns a JSON string for a component which includes the MDC
// component's type, and status.
func (c *C) String() string {
	return "{\"component\":\"" + c.name.String() + "\"," +
		"\"status\":\"" + c.status.String() + "\"}"
}

// String returns the string representation of a StatusType. One of
// "uninitialized", "stopped", or "running".
func (s StatusType) String() string {
	switch s {
	case Stopped:
		return "stopped"
	case Running:
		return "running"
	}

	return "uninitialized"
}

func makeMDComponent(t Type, dom *js.Object) (*js.Object, error) {
	var err error
	defer gojs.CatchException(&err)

	mdcObject := dom.Get("mdc")

	switch t {
	case Checkbox:
		return mdcObject.Get("checkbox").Get(t.String()), err
	case Dialog:
		return mdcObject.Get("dialog").Get(t.String()), err
	case PersistentDrawer:
		return mdcObject.Get("drawer").Get(t.String()), err
	case TemporaryDrawer:
		return mdcObject.Get("drawer").Get(t.String()), err
	case FormField:
		return mdcObject.Get("formField").Get(t.String()), err
	case GridList:
		return mdcObject.Get("gridList").Get(t.String()), err
	case IconToggle:
		return mdcObject.Get("iconToggle").Get(t.String()), err
	case LinearProgress:
		return mdcObject.Get("linearProgress").Get(t.String()), err
	case Menu:
		return mdcObject.Get("menu").Get(t.String()), err
	case Radio:
		return mdcObject.Get("radio").Get(t.String()), err
	case Ripple:
		return mdcObject.Get("ripple").Get(t.String()), err
	case Select:
		return mdcObject.Get("select").Get(t.String()), err
	case Slider:
		return mdcObject.Get("slider").Get(t.String()), err
	case Snackbar:
		return mdcObject.Get("snackbar").Get(t.String()), err
	case Tab:
		return mdcObject.Get("tabs").Get(t.String()), err
	case TabBar:
		return mdcObject.Get("tabs").Get(t.String()), err
	case TabBarScroller:
		return mdcObject.Get("tabs").Get(t.String()), err
	case TextField:
		return mdcObject.Get("textField").Get(t.String()), err
	case Toolbar:
		return mdcObject.Get("toolbar").Get(t.String()), err
	}
	return nil, err
}

// Start associates the component to an HTMLElement using a default
// querySelector that matches the first "div.mdc-[component-class]" element it
// finds. For more fine-grained control over the HTMLElement a component starts
// with, use the StartWith and StartWithElement methods.
//
// Upon success the component's status will be Running, and err will be nil.
//
// If err is non-nil, it will contain any error thrown while calling the
// underlying MDC object's init() method, and the component's status will remain
// Stopped.
func (c *C) Start() (err error) {
	return c.StartWith(".mdc-" + string(c.name.classString()))
}

// StartWith is like Start(), but allows you to specify the querySelector string
// used to associate a component with an HTMLElement.
//
// Upon success the component's status will be Running, and err will be nil.
//
// If err is non-nil, it will contain any error thrown while calling the
// underlying MDC object's init() method, and the component's status will remain
// Stopped.
func (c *C) StartWith(querySelector string) (err error) {
	defer gojs.CatchException(&err)

	e := js.Global.Get("window").Get("document").Call("querySelector",
		querySelector)

	return c.StartWithElement(e)
}

// StartWithElement is like StartWith, but accepts a *js.Object that must
// contain a valid HTMLElement for the component to associate itself with.
//
// Upon success the component's status will be Running, and err will be nil.
//
// If err is non-nil, it will contain any error thrown while calling the
// underlying MDC object's init() method, and the component's status will remain
// Stopped.
func (c *C) StartWithElement(e *js.Object) (err error) {
	defer gojs.CatchException(&err)

	if c.status == Running {
		return nil
	}

	if c.status != Stopped {
		return errors.New("Attempted to Start() an uninitialized component: " +
			c.String() + ". Use mdc.New()")
	}

	o := c.mdc.New(e)
	c.mdc = o
	c.status = Running

	return err
}

// Stop stops a Running component, removing its association with an HTMLElement
// and cleaning up event listeners, etc. It changes the component's status to
// Stopped.
func (c *C) Stop() (err error) {
	defer gojs.CatchException(&err)

	if c.status == Stopped {
		return errors.New("Cannot Stop() already stopped component: " +
			c.String())
	}

	if c.status != Running {
		return errors.New("Cannot Stop() an uninitialized component: " +
			c.String() + ". Use mdc.New()")
	}

	c.mdc.Call("destroy")

	return err
}

// CType returns the component's Type
func (c *C) CType() Type {
	return c.name
}

// GetObject returns the MDC component's JavaScript object
func (c *C) GetObject() *js.Object {
	return c.mdc
}
