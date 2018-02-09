// The snackbar package implements a material snackbar component.
//
// See: https://material.io/components/web/catalog/snackbars/
package snackbar // import "agamigo.io/material/snackbar"

import (
	"agamigo.io/gojs"
	"agamigo.io/material/component"
	"github.com/gopherjs/gopherjs/js"
)

// S is a material snackbar component. It should only be created using the New
// function.
type S struct {
	*component.C
	*Data

	// By default the snackbar will be dimissed when the user presses the action
	// button. If you want the snackbar to remain visible until the timeout is
	// reached (regardless of whether the user pressed the action button or not)
	// you can set the dismissesOnAction property to false.
	DismissOnAction bool `js:"dismissOnAction"`
}

// Data holds configuration for the snackbar.
type Data struct {
	object *js.Object

	// The text message to display.
	// Required.
	Message string `js:"message"`

	// The amount of time in milliseconds to show the snackbar.
	// Optional (default 2750).
	Timeout int `js:"timeout"`

	// The function to execute when the action is clicked.
	// Optional.
	ActionHandler func() `js:"actionHandler"`

	// The text to display for the action button.
	// Required if actionHandler is set.
	ActionText string `js:"actionText"`

	// Whether to show the snackbar with space for multiple lines of text.
	// Optional
	MultiLine bool `js:"multiline"`

	// Whether to show the action below the multiple lines of text.
	// Optional, applies when multiline is true.
	ActionOnBottom bool `js:"actionOnBottom"`
}

// New creates a material snackbar component. It is a wrapper around
// component.New which instantiates the component from the MDC library.
func New() (*S, error) {
	newS, err := component.New(component.Snackbar)
	if err != nil {
		return nil, err
	}
	data := &Data{object: js.Global.Get("Object")}
	data.Timeout = 2750
	return &S{C: newS, Data: data}, err
}

// Show displays the snackbar.
func (s *S) Show() error {
	var err error
	gojs.CatchException(&err)
	s.GetObject().Call("show", s.Data)
	return err
}

// TODO: Handle custom events
// - MDCSnackbar:show
// - MDCSnackbar:hide
