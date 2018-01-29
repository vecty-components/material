package mdc

import (
	"github.com/gopherjs/gopherjs/js"
)

type ComponentName int

const (
	Custom ComponentName = iota
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

type ComponentStatus int

const (
	Uninitialized ComponentStatus = iota
	Stopped
	Running
)

var (
	nextID = 1
	mdc    = js.Global.Get("mdc")
)

type Component struct {
	*js.Object
	name   ComponentName
	id     int
	status ComponentStatus
}

func New(n ComponentName) *Component {
	c := &Component{}
	c.name = n
	o := makeMDComponent(c)
	if o == nil || o == js.Undefined {
		panic("Creating " + c.Name().classString() +
			" failed, object nil or undefined")
	}
	c.SetObject(o)
	c.setStatus(Stopped)
	c.setID(nextID)
	nextID = nextID + 1
	return c
}

func (c *Component) ID() int {
	return c.id
}

func (c *Component) setID(id int) {
	c.id = id
}

func (c *Component) Status() ComponentStatus {
	return c.status
}

func (c *Component) setStatus(s ComponentStatus) {
	c.status = s
}

func (c *Component) Name() ComponentName {
	return c.name
}

func (c *Component) GetObject() *js.Object {
	return c.Object
}

func (c *Component) SetObject(o *js.Object) {
	c.Object = o
}

func (n ComponentName) componentString() string {
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

func (n ComponentName) classString() string {
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

func makeMDComponent(c *Component) *js.Object {
	switch c.Name() {
	case Animation:
		return mdc.Get("animation").Get(c.Name().componentString())
	case Checkbox:
		return mdc.Get("checkbox").Get(c.Name().componentString())
	case Dialog:
		return mdc.Get("dialog").Get(c.Name().componentString())
	case PermanentDrawer:
		return mdc.Get("drawer").Get(c.Name().componentString())
	case PersistentDrawer:
		return mdc.Get("drawer").Get(c.Name().componentString())
	case SlidableDrawer:
		return mdc.Get("drawer").Get(c.Name().componentString())
	case TemporaryDrawer:
		return mdc.Get("drawer").Get(c.Name().componentString())
	case FormField:
		return mdc.Get("formField").Get(c.Name().componentString())
	case GridList:
		return mdc.Get("gridList").Get(c.Name().componentString())
	case IconToggle:
		return mdc.Get("iconToggle").Get(c.Name().componentString())
	case LinearProgress:
		return mdc.Get("linearProgress").Get(c.Name().componentString())
	case Menu:
		return mdc.Get("menu").Get(c.Name().componentString())
	case Radio:
		return mdc.Get("radio").Get(c.Name().componentString())
	case Ripple:
		return mdc.Get("ripple").Get(c.Name().componentString())
	case Select:
		return mdc.Get("select").Get(c.Name().componentString())
	// case SelectionControl:
	// 	return ""
	case Slider:
		return mdc.Get("slider").Get(c.Name().componentString())
	case Snackbar:
		return mdc.Get("snackbar").Get(c.Name().componentString())
	case Tab:
		return mdc.Get("tab").Get(c.Name().componentString())
	case TabBar:
		return mdc.Get("tab").Get(c.Name().componentString())
	case TabBarScroller:
		return mdc.Get("tab").Get(c.Name().componentString())
	// case Textfield:
	// 	return ""
	case Toolbar:
		return mdc.Get("toolbar").Get(c.Name().componentString())
	}
	return nil
}

func (c *Component) Start() {
	switch c.Name() {
	case Checkbox:
		c.StartWith("div.mdc-" + string(c.Name().classString()))
	}
}

func (c *Component) StartWith(querySelector string) {
	if c.Status() == Running {
		return
	}
	if c.Status() != Stopped {
		panic("Attempted to run Start() an uninitialized component. Use mdc.New()")
	}

	e := js.Global.Get("document").Call("querySelector", querySelector)
	c.SetObject(c.GetObject().New(e))
	c.setStatus(Running)
}

func (c *Component) Stop() {
	if c.Status() == Stopped {
		println(c.Name().classString())
		print("Attempted to stop already stopped component: ")
		return
	}

	if c.Status() != Running {
		println(c.Name().classString())
		panic("Attempted to run Stop() an uninitialized component. Use mdc.New()")
	}

	c.GetObject().Call("destroy")
}
