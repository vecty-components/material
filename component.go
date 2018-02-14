package material // import "agamigo.io/material"

import (
	"errors"

	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

// Component is a base type for all Material components.
type Component struct {
	*js.Object
	Type ComponentType
}

// Component implements the material.Componenter interface.
func (c *Component) Component() *js.Object {
	return c.Object
}

// SetComponent implements the material.Componenter interface and replaces the
// Component's base Component with mdc.
func (c *Component) SetComponent(mdc *js.Object) {
	c.Object = mdc
}

// ComponentType implements the ComponentTyper interface.
func (c *Component) ComponentType() ComponentType {
	return c.Type
}

// Start takes a component implementation (c) and initializes it with an
// HTMLElement (rootElem). Upon success err will be nil. If err is non-nil, it
// will contain any error thrown while calling the underlying MDC object's
// init() method. An error will also be returned if Component() is non-nil.  Use
// Stop to clean up the component before calling Start again.
//
// Important: If you are using a component from agamigo.io/material/*, you
// should use its Start method, not this function. Consult the component's
// documentation for info/examples.
//
// Implementing A Component
//
// If you are writing a component implementation the documentation for the
// Componenter interface provides useful information.
//
// Finding The MDC Library
//
// There are two ways Start knows of to find the MDC class needed to start a
// component. By default it uses values provided by the components in this
// project via the ComponentType method. This default works in the general case
// that the all-in-one MDC library is available under the global var "mdc".
//
// The second case, MDCClasser, is needed if the MDC code for your component is
// elsewhere, for example if you are using the individual MDC component
// "@material/checkbox" library instead of the all-in-one distribution.
// Implement the MDCClasser interface to provide Start with the exact object for
// the MDC component class.
//
// See: https://material.io/components/web/docs/framework-integration/
func Start(c Componenter, rootElem *js.Object) (err error) {
	defer gojs.CatchException(&err)

	switch {
	case rootElem == nil, rootElem == js.Undefined:
		return errors.New("rootElem is nil.")
	case c.Component() != nil:
		return errors.New("Refusing to Start non-nil component. " +
			"Use Stop() before starting it again.")
	}

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
	c.SetComponent(newMDCClassObj.New(rootElem))

	return err
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc. It then runs SetComponent(nil).
func Stop(c Componenter) (err error) {
	defer gojs.CatchException(&err)

	if c.Component() == nil {
		return errors.New("GetComponent() returned nil.")
	}

	c.Component().Call("destroy")
	c.SetComponent(nil)
	return err
}
