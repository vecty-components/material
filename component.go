/*
The material package provides interfaces and functions needed to implement/run
material components in GopherJS.

Quickstart Guide

1. In your project include the all-in-one distribution of the MDC javascript
library and set it to the global variable "mdc". This can be done a number of
ways (HTML script element, webpack, filename "mdc.inc.js" for gopherjs to pick
up, etc).

2. Import a Material component from this project in your Go progrem.

	import "agamigo.io/material"
	import "agamigo.io/material/checkbox"

3. Make the HTML suitable for that MDC component available to your GopherJS
program. See: https://material.io/components/web/catalog/

	<html>
		<body>
			<div class="mdc-checkbox">
				<input class="mdc-checkbox__native-control" type="checkbox">
			</div>
		</body>
	</html>

4. Put that HTMLElement into a GopherJS object.

	cbElem := js.Global.Get("document").Get("body").Get("firstElementChild")

5. Create a new instance of the component and start it.

	cb := checkbox.CB{}
	material.Start(cb, cbElem)
*/
package material // import "agamigo.io/material"

import (
	"errors"

	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

// Component is the base material component type. Types that embed Component and
// implement Componenter can use the material.Start and material.Stop functions.
type Component struct {
	mdc    *js.Object
	status ComponentStatus
}

// Start takes a component implementation (c) and initializes it with an
// HTMLElement (rootElem). Upon success the component's status will be Running,
// and err will be nil.  If err is non-nil, it will contain any error thrown
// while calling the underlying MDC object's init() method, and the component's
// status will remain Stopped.
//
// Finding The MDC Library
//
// There are two ways Start knows of to find the MDC class needed to start a
// component. By default it uses values provided by the components in this
// project via the ComponentType method. This default works in the general case
// that the all-in-one MDC library is available under the global var "mdc".
//
// The second case, MDCClasser, is needed if the MDC code you need is elsewhere,
// for example if you are using individual MDC component "@material/checkbox"
// libraries instead of the all-in-one distribution. Implement the MDCClasser to
// provide Start with the exact object for the MDC component class.
//
// Implementing A Component
//
// If you are writing a component implementation the documentation for the
// Componenter interface provides useful information.
//
// If you need to perform additional work on the Component after initialization,
// read the AfterStarter interface documentation. If AfterStart returns a
// non-nill error then Stop will be called on the component.
//
// See: https://material.io/components/web/docs/framework-integration/
func Start(c Componenter, rootElem *js.Object) (err error) {
	defer gojs.CatchException(&err)

	switch {
	case rootElem == nil, rootElem == js.Undefined:
		return errors.New("rootElem is nil.")
	case c.GetComponent() == nil:
		c.SetComponent(&Component{})
	case c.GetComponent().status == Running:
		return errors.New("Component already started.")
	}

	// We create a new instance of the MDC component if c is Stopped or
	// Uninitialized.
	var newMDCClassObj *js.Object
	switch co := c.(type) {
	case MDCClasser:
		newMDCClassObj = co.MDCClass()
	case ComponentTyper:
		CCaseName := co.ComponentType().MDCCamelCaseName
		ClassName := co.ComponentType().MDCClassName
		if CCaseName == "" || ClassName == "" {
			return errors.New("Empty string in ComponentType")
		}
		mdcObject := js.Global.Get("mdc")
		newMDCClassObj = mdcObject.Get(CCaseName).Get(ClassName)
	default:
		return errors.New("The provided component does not implement " +
			"material.ComponentTyper or material.MDCClasser.")
	}

	// Create a new MDC component instance tied to rootElem
	newMDCObj := newMDCClassObj.New(rootElem)
	c.GetComponent().mdc = newMDCObj
	c.GetComponent().status = Running

	switch co := c.(type) {
	case AfterStarter:
		err = co.AfterStart()
		if err != nil {
			Stop(c)
			return err
		}
	}

	return err
}

// Stop stops a Running component, removing its association with an HTMLElement
// and cleaning up event listeners, etc. It changes the component's status to
// Stopped.
func Stop(c Componenter) (err error) {
	defer gojs.CatchException(&err)

	if c.GetComponent() == nil {
		return errors.New("GetComponent() returned nil.")
	}

	switch c.GetComponent().status {
	case Stopped:
		return errors.New("Component already stopped")
	case Uninitialized:
		return errors.New("Component is uninitialized")
	}
	c.GetComponent().mdc.Call("destroy")
	c.SetComponent(&Component{status: Stopped})
	return err
}

// GetComponent implements the Componenter interface. Component implementations
// can use this method as-is when embedding an exposed material.Component.
func (c *Component) GetComponent() *Component {
	return c
}

// ComponentType implements the Componenter interface. This should be shadowed
// by a component implementation.
func (c *Component) ComponentType() ComponentType {
	return ComponentType{}
}

// GetObject returns the component's MDC JavaScript object.
func (c *Component) GetObject() *js.Object {
	return c.mdc
}

// String returns the Component's StatusType as text.
func (c *Component) String() string {
	if c == nil || c.status == Uninitialized {
		return Uninitialized.String()
	}
	return c.Status().String()
}

// Status returns the component's StatusType. For the string version use
// Status().String().
func (c *Component) Status() ComponentStatus {
	return c.status
}
