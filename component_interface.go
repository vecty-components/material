package material // import "agamigo.io/material"

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

// AfterStarter is implemented by components that need further setup ran
// after their underlying MDC foundation has been initialized.
type AfterStarter interface {
	AfterStart() error
}

// ComponentTyper is one way to tell material.Start how to find the MDC library
// needed for a component. For more control, implement MDCClasser.
type ComponentTyper interface {
	ComponentType() ComponentType
}

// MDCClasser is an interface that allows component users to specify the MDC
// class object that will be used to create/initialize the component.
type MDCClasser interface {
	MDCClass() *js.Object
}
