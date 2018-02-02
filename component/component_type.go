package component

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
