// The dialog package implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/material/dialog"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
	"github.com/gopherjs/gopherjs/js"
)

// D is the interface for a material dialog component.
type D interface {
	component.C
	IsOpen() bool
	Open() error
	Close() error
	OnAcceptChan() chan *component.Event
	OnCancelChan() chan *component.Event
}

// dialog is the internal implementation of D made available publicly via
// New().
type dialog struct {
	component.C
	acceptChan chan *component.Event
	cancelChan chan *component.Event
}

// New creates a material dialog component that implement the D interface.
// It is a wrapper around component.New.
func New() (c D, err error) {
	newD, err := component.New(component.Dialog)
	if err != nil {
		return nil, err
	}
	d := &dialog{
		C:          newD,
		acceptChan: make(chan *component.Event),
		cancelChan: make(chan *component.Event),
	}
	return d, err
}

// IsOpen returns the state of the dialog, open or closed.
func (d *dialog) IsOpen() bool {
	// return d.GetObject().Get("foundation_").Get("isOpen_").Bool()
	return d.GetObject().Get("open").Bool()
}

// Open shows the dialog. If the dialog is already open then Open is a no-op.
func (d *dialog) Open() error {
	var err error
	defer gojs.CatchException(&err)
	d.GetObject().Call("show")
	return err
}

// Close removes the dialog from view. If the dialog is already closed then
// Close is a no-op.
func (d *dialog) Close() error {
	var err error
	defer gojs.CatchException(&err)
	d.GetObject().Call("close")
	return err
}

// Start wraps component.Start and adds event listeners that pass dialog events
// over the channels provided by the dialog's OnAccessChan()/OnCancelChan()
// methods.
func (d *dialog) Start() error {
	var err error
	defer gojs.CatchException(&err)

	err = d.C.Start()
	if err != nil {
		return err
	}

	d.GetObject().Call("listen", "MDCDialog:accept",
		func(e *js.Object) {
			d.acceptChan <- &component.Event{
				Type:      "MDCDialog:accept",
				Event:     e,
				Component: d,
			}
		},
	)

	d.GetObject().Call("listen", "MDCDialog:cancel",
		func(e *js.Object) {
			d.cancelChan <- &component.Event{
				Type:      "MDCDialog:cancel",
				Event:     e,
				Component: d,
			}
		},
	)

	return err
}

// Stop wraps component.Stop and signals to receivers of
// OnAcceptChan()/OnCancelChan() to close communication.
func (d *dialog) Stop() error {
	err := d.C.Stop()
	if err != nil {
		return err
	}
	d.acceptChan <- nil
	return nil
}

// OnAcceptChan returns a channel through which the dialog sends event details
// when a user chooses "accept" from the open dialog UI. Receivers should always
// check for a nil value sent on this channel which signals that the component
// has been stopped.
func (d *dialog) OnAcceptChan() chan *component.Event {
	return d.acceptChan
}

// OnCancelChan returns a channel through which the dialog sends event details
// when a user chooses "cancel" from the open dialog UI. Receivers should always
// check for a nil value sent on this channel which signals that the component
// has been stopped.
func (d *dialog) OnCancelChan() chan *component.Event {
	return d.cancelChan
}
