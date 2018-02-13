package material // import "agamigo.io/material"

// ComponentStatus holds a component's lifecycle status.
type ComponentStatus int

const (
	// An Uninitialized component has not been associated with the MDC library
	// yet. This package does not provide a way to access an Uninitialized
	// component.
	Uninitialized ComponentStatus = iota

	// A Stopped component has been associated with a JS Object constructed from
	// a MDC class. New() returns a Stopped component, and Stop() will stop a
	// Running component.
	Stopped

	// A Running component has had its underlying MDC init() method called,
	// which attaches the component to a specific HTMLElement in the DOM. It is
	// ready to be used.
	Running
)

// String returns the string version of a StatusType.
func (s ComponentStatus) String() string {
	switch s {
	case Stopped:
		return "stopped"
	case Running:
		return "running"
	}
	return "uninitialized"
}
