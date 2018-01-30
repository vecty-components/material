package component

import (
	"github.com/gopherjs/gopherjs/js"
)

type Type int

const (
	Custom Type = iota
	Animation
	Checkbox
	Dialog
	FormField
	PermanentDrawer
	PersistentDrawer
	SlidableDrawer
	TemporaryDrawer
	GridList
	IconToggle
	LinearProgress
	Menu
	Radio
	Ripple
	Select
	// SelectionControl
	Slider
	Snackbar
	Tab
	TabBar
	TabBarScroller
	// Textfield
	Toolbar
)

type StatusType int

const (
	Uninitialized StatusType = iota
	Stopped
	Running
)

var (
	mdcObject = js.Global.Get("mdc")
)

type C interface {
	GetObject() *js.Object
	Start()
	StartWith(querySelector string)
	Stop()
}

type component struct {
	*js.Object
	name   Type
	status StatusType
}

func New(n Type) C {
	c := &component{}
	c.name = n
	o := makeMDComponent(c)
	if o == nil || o == js.Undefined {
		panic("Creating " + c.name.classString() +
			" failed, object nil or undefined")
	}
	c.setObject(o)
	c.setStatus(Stopped)
	return c
}

func (c *component) setStatus(s StatusType) {
	c.status = s
}

func (c *component) GetObject() *js.Object {
	return c.Object
}

func (c *component) setObject(o *js.Object) {
	c.Object = o
}

func (n Type) componentString() string {
	switch n {
	case Animation:
		return "MDCAnimation"
	case Checkbox:
		return "MDCCheckbox"
	case Dialog:
		return "MDCDialog"
	case PermanentDrawer:
		return "MDCDrawer"
	case PersistentDrawer:
		return "MDCDrawer"
	case SlidableDrawer:
		return "MDCDrawer"
	case TemporaryDrawer:
		return "MDCDrawer"
	case FormField:
		return "MDCFormField"
	case GridList:
		return "MDCGridList"
	case IconToggle:
		return "MDCIconToggle"
	case LinearProgress:
		return "MDCLinearProgress"
	case Menu:
		return "MDCMenu"
	case Radio:
		return "MDCRadio"
	case Ripple:
		return "MDCRipple"
	case Select:
		return "MDCSelect"
	// case SelectionControl:
	// 	return ""
	case Slider:
		return "MDCSlider"
	case Snackbar:
		return "MDCSnackbar"
	case Tab:
		return "MDCTab"
	case TabBar:
		return "MDCTabBar"
	case TabBarScroller:
		return "MDCTabBarScroller"
	// case Textfield:
	// 	return ""
	case Toolbar:
		return "MDCToolbar"
	}

	panic("Failed to convert MDCName to component string.")
	return ""
}

func (n Type) classString() string {
	switch n {
	case Animation:
		return "animation"
	case Checkbox:
		return "checkbox"
	case Dialog:
		return "dialog"
	case PermanentDrawer:
		return "drawer--permanent"
	case PersistentDrawer:
		return "drawer--persistent"
	case SlidableDrawer:
		return "drawer--slidable"
	case TemporaryDrawer:
		return "drawer--temporary"
	case FormField:
		return "form-field"
	case GridList:
		return "grid-list"
	case IconToggle:
		return "icon-toggle"
	case LinearProgress:
		return "linear-progress"
	case Menu:
		return "menu"
	case Radio:
		return "radio"
	case Ripple:
		return "ripple"
	case Select:
		return "select"
	// case SelectionControl:
	// 	return ""
	case Slider:
		return "slider"
	case Snackbar:
		return "snackbar"
	case Tab:
		return "tab"
	case TabBar:
		return "tab-bar"
	case TabBarScroller:
		return "tab-bar-scroller"
	// case Textfield:
	// 	return ""
	case Toolbar:
		return "toolbar"
	}

	panic("Failed to convert MDCName to class string.")
	return ""
}

func makeMDComponent(c *component) *js.Object {
	switch c.name {
	case Animation:
		return mdcObject.Get("animation").Get(c.name.componentString())
	case Checkbox:
		return mdcObject.Get("checkbox").Get(c.name.componentString())
	case Dialog:
		return mdcObject.Get("dialog").Get(c.name.componentString())
	case PermanentDrawer:
		return mdcObject.Get("drawer").Get(c.name.componentString())
	case PersistentDrawer:
		return mdcObject.Get("drawer").Get(c.name.componentString())
	case SlidableDrawer:
		return mdcObject.Get("drawer").Get(c.name.componentString())
	case TemporaryDrawer:
		return mdcObject.Get("drawer").Get(c.name.componentString())
	case FormField:
		return mdcObject.Get("formField").Get(c.name.componentString())
	case GridList:
		return mdcObject.Get("gridList").Get(c.name.componentString())
	case IconToggle:
		return mdcObject.Get("iconToggle").Get(c.name.componentString())
	case LinearProgress:
		return mdcObject.Get("linearProgress").Get(c.name.componentString())
	case Menu:
		return mdcObject.Get("menu").Get(c.name.componentString())
	case Radio:
		return mdcObject.Get("radio").Get(c.name.componentString())
	case Ripple:
		return mdcObject.Get("ripple").Get(c.name.componentString())
	case Select:
		return mdcObject.Get("select").Get(c.name.componentString())
	// case SelectionControl:
	// 	return ""
	case Slider:
		return mdcObject.Get("slider").Get(c.name.componentString())
	case Snackbar:
		return mdcObject.Get("snackbar").Get(c.name.componentString())
	case Tab:
		return mdcObject.Get("tab").Get(c.name.componentString())
	case TabBar:
		return mdcObject.Get("tab").Get(c.name.componentString())
	case TabBarScroller:
		return mdcObject.Get("tab").Get(c.name.componentString())
	// case Textfield:
	// 	return ""
	case Toolbar:
		return mdcObject.Get("toolbar").Get(c.name.componentString())
	}
	return nil
}

func (c *component) Start() {
	switch c.name {
	case Checkbox:
		c.StartWith("div.mdc-" + string(c.name.classString()))
	}
}

func (c *component) StartWith(querySelector string) {
	if c.status == Running {
		return
	}
	if c.status != Stopped {
		panic("Attempted to run Start() an uninitialized component. Use mdc.New()")
	}

	e := js.Global.Get("document").Call("querySelector", querySelector)
	c.setObject(c.GetObject().New(e))
	c.setStatus(Running)
}

func (c *component) Stop() {
	if c.status == Stopped {
		println(c.name.classString())
		print("Attempted to stop already stopped component: ")
		return
	}

	if c.status != Running {
		println(c.name.classString())
		panic("Attempted to run Stop() an uninitialized component. Use mdc.New()")
	}

	c.GetObject().Call("destroy")
}
