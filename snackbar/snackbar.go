// snackbar implements a material snackbar component.
//
// See: https://material.io/components/web/catalog/snackbars/
package snackbar // import "agamigo.io/material/snackbar"

import (
	"errors"

	"agamigo.io/gojs"
	"agamigo.io/material"
	"github.com/gopherjs/gopherjs/js"
)

// S is a material snackbar component.
type S struct {
	*material.Component
	isNew bool

	// DismissOnAction causes the snackbar to be dimissed when the user presses
	// the action button. If you want the snackbar to remain visible until the
	// timeout is reached (regardless of whether the user pressed the action
	// button or not) you can set this to to false.
	DismissOnAction bool `js:"dismissOnAction"`

	// The text message to display.
	// Required.
	Message string `js:"message"`

	// The amount of time in milliseconds to show the snackbar.
	// Default is 2750.
	Timeout int `js:"timeout"`

	// The function to execute when the action is clicked.
	// Optional.
	ActionHandler func() `js:"actionHandler"`

	// The text to display for the action button.
	// Required if actionHandler is set.
	ActionText string `js:"actionText"`

	// Whether to show the snackbar with space for multiple lines of text.
	// Default is false.
	MultiLine bool `js:"multiline"`

	// Whether to show the action below the multiple lines of text.
	// Optional, applies when multiline is true. Default is false.
	ActionOnBottom bool `js:"actionOnBottom"`
}

// data holds configuration for the snackbar.
// type Data struct {
// 	*js.Object
// 	message        string `js:"message"`
// 	timeout        int    `js:"timeout"`
// 	actionHandler  func() `js:"actionHandler"`
// 	actionText     string `js:"actionText"`
// 	multiLine      bool   `js:"multiline"`
// 	actionOnBottom bool   `js:"actionOnBottom"`
// }

// ComponentType implements the ComponentTyper interface.
func (c *S) ComponentType() material.ComponentType {
	return material.ComponentType{
		MDCClassName:     "MDCSnackbar",
		MDCCamelCaseName: "snackbar",
	}
}

// SetComponent implements the Componenter interface and replaces the component's
// base Component with mdcC.
func (c *S) SetComponent(mdcC *material.Component) {
	c.Component = mdcC
}

// String returns the component's "ComponentType: status" information.
func (c *S) String() string {
	return c.ComponentType().String() + ": " + c.Component.String()
}

// Show displays the snackbar. If the configuration is invalid an error message
// will be returned and the snackbar will not be shown. For information on
// config requirements look at documentation for S.
func (c *S) Show() error {
	var err error
	gojs.CatchException(&err)
	if c.Message == "" {
		return errors.New("Snackbar Message is empty.")
	}
	if c.isNew && c.Timeout == 0 {
		c.Timeout = 2750
	}
	data := make(js.M)
	data["message"] = c.Message
	data["timeout"] = c.Timeout
	data["multiline"] = c.MultiLine
	data["actionOnBottom"] = c.ActionOnBottom
	if c.ActionHandler != nil {
		if c.ActionText == "" {
			return errors.New(
				"Snackbar has ActionHandler, but ActionText is empty.")
		}
		data["actionHandler"] = c.ActionHandler
		data["actionText"] = c.ActionText
	}
	c.GetObject().Call("show", data)
	c.isNew = false
	return err
}

// TODO: Handle custom events
// - MDCSnackbar:show
// - MDCSnackbar:hide
