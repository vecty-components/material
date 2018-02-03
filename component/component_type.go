package component

type Type int

const (
	Checkbox Type = iota
	Dialog
	FormField
	PersistentDrawer
	TemporaryDrawer
	GridList
	IconToggle
	LinearProgress
	Menu
	Radio
	Ripple
	Select
	Slider
	Snackbar
	Tab
	TabBar
	TabBarScroller
	TextField
	Toolbar
	TypeCount
)

func ComponentTypes() []Type {
	types := make([]Type, TypeCount, TypeCount)
	for i := Type(0); i < TypeCount; i = i + 1 {
		types[i] = i
	}
	return types
}

func (n Type) String() string {
	switch n {
	case Checkbox:
		return "MDCCheckbox"
	case Dialog:
		return "MDCDialog"
	case PersistentDrawer:
		return "MDCPersistentDrawer"
	case TemporaryDrawer:
		return "MDCTemporaryDrawer"
	case FormField:
		return "MDCFormField"
	case GridList:
		return "MDCGridList"
	case IconToggle:
		return "MDCIconToggle"
	case LinearProgress:
		return "MDCLinearProgress"
	case Menu:
		return "MDCSimpleMenu"
	case Radio:
		return "MDCRadio"
	case Ripple:
		return "MDCRipple"
	case Select:
		return "MDCSelect"
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
	case TextField:
		return "MDCTextField"
	case Toolbar:
		return "MDCToolbar"
	}

	panic("Failed to convert MDCName to component string.")
}

func (n Type) classString() string {
	switch n {
	case Checkbox:
		return "checkbox"
	case Dialog:
		return "dialog"
	// case PermanentDrawer:
	// 	return "drawer--permanent"
	case PersistentDrawer:
		return "drawer--persistent"
	// case SlidableDrawer:
	// 	return "drawer--slidable"
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
	case TextField:
		return "text-field"
	case Toolbar:
		return "toolbar"
	}

	panic("Failed to convert MDCName to class string.")
}
