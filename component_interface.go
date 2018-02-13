package material // import "agamigo.io/material"

import "github.com/gopherjs/gopherjs/js"

// Componenter is the base interface for every material component
// implementation.
type Componenter interface {
	// SetComponent should replace a component implementation's Component with
	// the provided component.
	SetComponent(c *Component)

	// GetComponent should return a pointer to the component implementation's
	// underlying Component. Implementors that embed a *Component directly
	// get this for free.
	GetComponent() (c *Component)
}

// AfterStarter is implemented by components that need further setup ran
// after their underlying MDC foundation has been initialized.
type AfterStarter interface {
	AfterStart() error
}

type ComponentTyper interface {
	ComponentType() ComponentType
}

// MDCClasser is an interface that allows component users to specify the MDC
// class object that will be used to create/initialize the component.
type MDCClasser interface {
	MDCClass() *js.Object
}
