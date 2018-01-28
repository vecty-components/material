package checkbox

import (
	"agamigo.io/material-components-go/mdc"
	"github.com/gopherjs/gopherjs/js"
)

const (
	MDCNAME = mdc.MDCNameCheckbox
)

type C struct {
	*js.Object
	mdc.Component
}

func New() *C {
	c := &C{}
	mdc.New(c)
	return c
}

// MDCName implements mdc.Adaptable
func (c *C) MDCName() mdc.MDCName {
	return MDCNAME
}

// SetObject implements mdc.Adaptable
func (c *C) SetObject(o *js.Object) {
	c.Object = o
}

func (c *C) GetObject() *js.Object {
	return c.Object
}

func (c *C) Start() {
	mdc.Start(c)
}

func (c *C) StartWith(querySelector string) {
	mdc.StartWith(c, querySelector)
}

func (c *C) Stop() {
	mdc.Stop(c)
}
