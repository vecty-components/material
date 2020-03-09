package base // import "github.com/vecty-material/material/base"

import "syscall/js"

// Componenter is a base interface for every material component implementation.
type Componenter interface {
	// Component should return the object that holds its MDC instance.
	Component() (c *Component)
}

// ComponentStartStopper is an interface that all material components implement.
// It is not used within the material project, but is intended for use by its
// consumers that embed a material component directly. Then frameworks/functions
// etc. can accept any component.
type ComponentStartStopper interface {
	Componenter
	Start(rootElem js.Value) error
	Stop() error
}

// MDCClasser is an interface that allows component users to specify the MDC
// class object that will be used to create/initialize the component. It
// overrides ComponentTyper when calling base.Start.
type MDCClasser interface {
	MDCClass() js.Value
}

// StateMapper is an interface that components implement in order to provide a
// map of state values which can be used for backup/restore.
type StateMapper interface {
	StateMap() StateMap
}
