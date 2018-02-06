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

// C is the base interface that all material components implement.
type C interface {
	// GetObject provides access to the underlying MDC JavaScript object.
	GetObject() *js.Object

	// Start{With,WithElement} starts the component, associating it with an
	// HTMLElement, making it ready for use.
	Start() error
	StartWith(querySelector string) error
	StartWithElement(element *js.Object) error

	// Stop stops the component, disassociating it with its HTMLElement and
	// cleaning up event-listeners.
	Stop() error

	String() string
	CType() Type
}

// component is the internal implementation of C.
type component struct {
	*js.Object
	name   Type
	status StatusType
}

// New creates a material component that implements the C interface. It assumes
// the MDC library and resulting component will live in the js.Global scope.
func New(n Type) (mdcComponent C, err error) {
	defer gojs.CatchException(&err)

	c, err := NewWith(n, js.Global)
	return c, err
}

// NewWith is like New(), with added option of specifying a *js.Object to store
// the component. This is primarily intended for use in tests where we may want
// to emulate a DOM somewhere other than Node's global scope.
func NewWith(n Type, dom *js.Object) (mdcComponent C, err error) {
	defer gojs.CatchException(&err)

	c := &component{}
	c.name = n

	o, err := makeMDComponent(c, dom)
	if err != nil {
		return nil, err
	}

	c.setObject(o)
	c.setStatus(Stopped)
	return c, err
}

// String returns a JSON string for a component which includes the MDC
// component's type, and status.
func (c *component) String() string {
	return "{\"component\":\"" + c.name.String() + "\"," +
		"\"status\":\"" + c.status.String() + "\"}"
}

func (c *component) setStatus(s StatusType) {
	c.status = s
}

// GetObject implements the C interface.
func (c *component) GetObject() *js.Object {
	return c.Object
}

func (c *component) setObject(o *js.Object) error {
	var err error
	defer gojs.CatchException(&err)
	c.Object = o
	return err
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

func makeMDComponent(c *component, dom *js.Object) (*js.Object, error) {
	var err error
	defer gojs.CatchException(&err)

	mdcObject := dom.Get("mdc")

	switch c.name {
	case Checkbox:
		return mdcObject.Get("checkbox").Get(c.name.String()), err
	case Dialog:
		return mdcObject.Get("dialog").Get(c.name.String()), err
	case PersistentDrawer:
		return mdcObject.Get("drawer").Get(c.name.String()), err
	case TemporaryDrawer:
		return mdcObject.Get("drawer").Get(c.name.String()), err
	case FormField:
		return mdcObject.Get("formField").Get(c.name.String()), err
	case GridList:
		return mdcObject.Get("gridList").Get(c.name.String()), err
	case IconToggle:
		return mdcObject.Get("iconToggle").Get(c.name.String()), err
	case LinearProgress:
		return mdcObject.Get("linearProgress").Get(c.name.String()), err
	case Menu:
		return mdcObject.Get("menu").Get(c.name.String()), err
	case Radio:
		return mdcObject.Get("radio").Get(c.name.String()), err
	case Ripple:
		return mdcObject.Get("ripple").Get(c.name.String()), err
	case Select:
		return mdcObject.Get("select").Get(c.name.String()), err
	case Slider:
		return mdcObject.Get("slider").Get(c.name.String()), err
	case Snackbar:
		return mdcObject.Get("snackbar").Get(c.name.String()), err
	case Tab:
		return mdcObject.Get("tabs").Get(c.name.String()), err
	case TabBar:
		return mdcObject.Get("tabs").Get(c.name.String()), err
	case TabBarScroller:
		return mdcObject.Get("tabs").Get(c.name.String()), err
	case TextField:
		return mdcObject.Get("textField").Get(c.name.String()), err
	case Toolbar:
		return mdcObject.Get("toolbar").Get(c.name.String()), err
	}
	return nil, err
}

// Start implements the C interface. It associates the component to an
// HTMLElement using a default querySelector that matches the first
// "div.mdc-[component-class]" element it finds. For more fine-grained control
// over the HTMLElement a component starts with, use the StartWith and
// StartWithElement methods.
//
// Upon success the component's status will be Running, and err will be nil.
//
// If err is non-nil, it will contain any error thrown while calling the
// underlying MDC object's init() method, and the component's status will remain
// Stopped.
func (c *component) Start() (err error) {
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
func (c *component) StartWith(querySelector string) (err error) {
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
func (c *component) StartWithElement(e *js.Object) (err error) {
	defer gojs.CatchException(&err)

	if c.status == Running {
		return nil
	}

	if c.status != Stopped {
		return errors.New("Attempted to Start() an uninitialized component: " +
			c.String() + ". Use mdc.New()")
	}

	o := c.GetObject().New(e)
	err = c.setObject(o)
	c.setStatus(Running)

	return err
}

// Stop stops a Running component, removing its association with an HTMLElement
// and cleaning up event listeners, etc. It changes the component's status to
// Stopped.
func (c *component) Stop() (err error) {
	defer gojs.CatchException(&err)

	if c.status == Stopped {
		return errors.New("Cannot Stop() already stopped component: " +
			c.String())
	}

	if c.status != Running {
		return errors.New("Cannot Stop() an uninitialized component: " +
			c.String() + ". Use mdc.New()")
	}

	c.GetObject().Call("destroy")

	return err
}

// CType returns the component's Type
func (c *component) CType() Type {
	return c.name
}
