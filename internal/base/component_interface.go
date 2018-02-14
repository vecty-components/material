package base // import "agamigo.io/material/internal/base"

import "github.com/gopherjs/gopherjs/js"

// Componenter is the base interface for every material component
// implementation.
type Componenter interface {
	// GetComponent should return the object that holds its MDC instance.
	Component() (mdc *js.Object)

	// SetComponent should replace a component implementation's *js.Object
	// variable that holds its MDC instance.
	SetComponent(mdc *js.Object)
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
