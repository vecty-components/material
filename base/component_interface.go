package base // import "agamigo.io/material/base"

import "github.com/gopherjs/gopherjs/js"

// Componenter is a base interface for every material component implementation.
type Componenter interface {
	// Component should return the object that holds its MDC instance.
	Component() (c *Component)
}

type ComponentStartStopper interface {
	Componenter
	Start(rootElem *js.Object) error
	Stop() error
}

// ComponentSetter is a base interface for every material component
// implementation.
type ComponentSetter interface {
	Componenter

	// SetComponent should replace a component implementation's *js.Object
	// variable that holds its MDC instance.
	SetComponent(c *Component)
}

// ComponentTyper is one way to tell base.Start how to find the MDC library
// needed for a component. For more control, implement MDCClasser.
type ComponentTyper interface {
	ComponentType() ComponentType
}

// MDCClasser is an interface that allows component users to specify the MDC
// class object that will be used to create/initialize the component. It
// overrides ComponentTyper when calling base.Start.
type MDCClasser interface {
	MDCClass() *js.Object
}

type StateMapper interface {
	StateMap() StateMap
}
