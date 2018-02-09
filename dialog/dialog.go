// The dialog package implements a material dialog component.
//
// See: https://material.io/components/web/catalog/dialogs/
package dialog // import "agamigo.io/material/dialog"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
)

// D is a material dialog component. It should only be created using the New
// function.
type D struct {
	*component.C
	IsOpen bool `js:"open"`
	// AcceptChan chan *Event
	// CancelChan chan *Event
}

// New creates a material dialog component. It is a wrapper around component.New
// which instantiates the component from the MDC library.
func New() (*D, error) {
	newD, err := component.New(component.Dialog)
	if err != nil {
		return nil, err
	}
	d := &D{
		C: newD,
		// AcceptChan: make(chan *Event),
		// CancelChan: make(chan *Event),
	}
	return d, err
}

// Open shows the dialog. If the dialog is already open then Open is a no-op.
func (d *D) Open() error {
	var err error
	defer gojs.CatchException(&err)
	d.GetObject().Call("show")
	return err
}

// Close removes the dialog from view. If the dialog is already closed then
// Close is a no-op.
func (d *D) Close() error {
	var err error
	defer gojs.CatchException(&err)
	d.GetObject().Call("close")
	return err
}

// // Start wraps component.Start and adds event listeners that pass dialog events
// // over the channels provided by the dialog's OnAccessChan()/OnCancelChan()
// // methods.
// func (d *D) Start() error {
// 	var err error
// 	defer gojs.CatchException(&err)

// 	err = d.C.Start()
// 	if err != nil {
// 		return err
// 	}

// 	d.GetObject().Call("listen", "MDCDialog:accept",
// 		func(e *js.Object) {
// 			d.AcceptChan <- &Event{
// 				Type:      "MDCDialog:accept",
// 				Event:     e,
// 				Component: d,
// 			}
// 		},
// 	)

// 	d.GetObject().Call("listen", "MDCDialog:cancel",
// 		func(e *js.Object) {
// 			d.CancelChan <- &Event{
// 				Type:      "MDCDialog:cancel",
// 				Event:     e,
// 				Component: d,
// 			}
// 		},
// 	)

// 	return err
// }

// // Stop wraps component.Stop and signals to receivers of
// // OnAcceptChan()/OnCancelChan() to close communication.
// func (d *D) Stop() error {
// 	err := d.C.Stop()
// 	if err != nil {
// 		return err
// 	}
// 	d.AcceptChan <- nil
// 	return nil
// }

// OnAcceptChan returns a channel through which the dialog sends event details
// when a user chooses "accept" from the open dialog UI. Receivers should always
// check for a nil value sent on this channel which signals that the component
// has been stopped.
// func (d *D) OnAcceptChan() chan *component.Event {
// 	return d.acceptChan
// }

// OnCancelChan returns a channel through which the dialog sends event details
// when a user chooses "cancel" from the open dialog UI. Receivers should always
// check for a nil value sent on this channel which signals that the component
// has been stopped.
// func (d *D) OnCancelChan() chan *component.Event {
// 	return d.cancelChan
// }

// type Event struct {
// 	Type      string
// 	Event     *js.Object
// 	Component *D
// }
