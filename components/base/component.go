/*
The base package contains code shared by implementations of material components
for GopherJS.
*/
package base // import "github.com/vecty-material/material/components/base"

import (
	"errors"

	"syscall/js"

	"github.com/vecty-material/material/gojs"
)

// Component is a base type for all Material components.
type Component struct {
	js.Value
	*MDCState
	Type ComponentType
}

type MDCState struct {
	Basic       bool
	Started     bool
	RootElement js.Value
}

type StateMap map[string]interface{}

// Component implements the base.Componenter interface.
func (c *Component) Component() *Component {
	if c.Value.IsUndefined() {
		c.Value = js.Global().Get("Object").New()
	}
	if c.MDCState == nil {
		c.MDCState = &MDCState{}
	}
	return c
}

// SetComponent implements the base.ComponentSetter interface and replaces the
// Component's properties with those of c's.
func (c *Component) SetComponent(newC *Component) {
	c = newC
}

// ComponentType implements the ComponentTyper interface.
func (c *Component) ComponentType() ComponentType {
	return c.Type
}

func (c *Component) Start(rootElem js.Value) error {
	return nil
}

func (c *Component) Stop() error {
	return nil
}

func (c *Component) SetState(sm StateMap) *Component {
	if sm != nil {
		for k, v := range sm {
			if v != nil {
				c.Component().Set(k, v)
			}
		}
	}
	return c
}

// Start takes a component implementation (c) and initializes it with an
// HTMLElement (rootElem). Upon success err will be nil. If err is non-nil, it
// will contain any error thrown while calling the underlying MDC object's
// init() method. An error will also be returned if Component() is non-nil.  Use
// Stop to clean up the component before calling Start again.
//
// Important: If you are using a component from github.com/vecty-material/material/*, you
// should use its Start method, not this function. Consult the component's
// documentation for info/examples.
//
// Implementing A Component
//
// If you are writing a component implementation the documentation for the
// Componenter{Setter} interfaces provides useful information.
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
func Start(c Componenter, rootElem js.Value) (err error) {
	defer gojs.CatchException(&err)

	backup := StateMap{}
	if sm, ok := c.(StateMapper); ok {
		backup = sm.StateMap()
		defer c.Component().SetState(backup)
	}

	if c.Component().MDCState.Basic {
		return nil
	}
	if c.Component().MDCState.Started {
		err = Stop(c)
		if err != nil {
			return err
		}
	}
	if rootElem.IsUndefined() {
		return errors.New("rootElem is nil.")
	}

	var newMDCClassObj js.Value
	switch t := c.(type) {
	case MDCClasser:
		newMDCClassObj = t.MDCClass()
	default:
		CCaseName := t.Component().ComponentType().MDCCamelCaseName
		ClassName := t.Component().ComponentType().MDCClassName
		if CCaseName == "" || ClassName == "" {
			return errors.New("Empty string in ComponentType")
		}
		mdcObject := js.Global().Get("mdc")
		newMDCClassObj = mdcObject.Get(CCaseName).Get(ClassName)
	}

	// Create a new MDC component instance tied to rootElem
	c.Component().Value = newMDCClassObj.New(rootElem)
	c.Component().MDCState.RootElement = rootElem
	c.Component().MDCState.Started = true

	return err
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc. It then runs SetComponent(nil).
func Stop(c Componenter) (err error) {
	defer gojs.CatchException(&err)

	if !c.Component().MDCState.Started {
		return errors.New("Refusing to stop non-started component.")
	}
	if c.Component() == nil {
		return errors.New("GetComponent() returned nil.")
	}
	c.Component().Call("destroy")
	c.Component().SetComponent(nil)
	return err
}
