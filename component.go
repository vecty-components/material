package material // import "agamigo.io/material"

import (
	"errors"

	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

// Start takes a component implementation (c) and initializes it with an
// HTMLElement (rootElem). Upon success err will be nil. If err is non-nil, it
// will contain any error thrown while calling the underlying MDC object's
// init() method. An error will also be returned if GetComponent() is non-nil.
// Use Stop to clean up the component before calling Start again.
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

// Stop stops a started component, removing its association with an HTMLElement
// and cleaning up event listeners, etc. It then runs SetComponent(nil).
func Stop(c Componenter) (err error) {
	defer gojs.CatchException(&err)

	if c.Component() == nil {
		return errors.New("GetComponent() returned nil.")
	}

	c.Component().Call("destroy")
	c.SetComponent(nil)
	return err
}
