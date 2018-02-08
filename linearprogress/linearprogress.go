// The linearprogress package implements a material linearprogress component.
//
// See: https://material.io/components/web/catalog/linear-progress/
package linearprogress // import "agamigo.io/material/linearprogress"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
	"github.com/gopherjs/gopherjs/js"
)

// LP is a material libearprogress component. It should only be created using
// the New function.
type LP struct {
	*component.C
	*State
}

// State holds the linearprogress's r/w settings. It is in its own struct due to
// the structure of the MDC component storing properties in the "foundation_".
// The properties here can still be accessed directly from an instance of LP.
//
// TODO: Find a way to put these properties into the LP struct directly.
type State struct {
	*state
	Determinate bool `js:"determinate"`
	Reverse     bool `js:"reverse"`
	Progress    int  `js:"progress"`
	Buffer      int  `js:"buffer"`
}

// state exists as the unexposed portion of State.
type state struct {
	*js.Object
}

// New creates a material linearprogress component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*LP, error) {
	newLP, err := component.New(component.LinearProgress)
	if err != nil {
		return nil, err
	}
	lp := &LP{newLP, &State{state: &state{}}}
	return lp, err
}

// Start wraps component.Start.
func (lp *LP) Start() (err error) {
	err = lp.C.Start()
	if err != nil {
		return err
	}

	lp.State.Object = lp.GetObject().Get("foundation_")
	return err
}

// StartWith wraps component.StartWith.
func (lp *LP) StartWith(querySelector string) (err error) {
	err = lp.C.StartWith(querySelector)
	if err != nil {
		return err
	}

	lp.State = &State{state: &state{Object: lp.Get("foundation_")}}
	return err
}

// StartWithElement wraps component.StartWithElement.
func (lp *LP) StartWithElement(e *js.Object) (err error) {
	err = lp.C.StartWithElement(e)
	if err != nil {
		return err
	}

	lp.State = &State{state: &state{Object: lp.Get("foundation_")}}
	return err
}

// Open opens the linearProgress component.
func (lp *LP) Open() (err error) {
	gojs.CatchException(&err)
	lp.GetObject().Call("open")
	return err
}

// Close closes the linearProgress component.
func (lp *LP) Close() (err error) {
	gojs.CatchException(&err)
	lp.GetObject().Call("close")
	return err
}
