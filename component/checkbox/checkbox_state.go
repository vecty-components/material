package checkbox

type StateType int

const (
	// Unset state is zero
	UNKNOWN StateType = iota
	DISABLED
	// Enabled states are even, disabled are odd
	UNCHECKED
	UNCHECKED_DISABLED
	CHECKED
	CHECKED_DISABLED
	INDETERMINATE
	INDETERMINATE_DISABLED
)

type StateCfg struct {
	State    string
	Disabled bool
}

// func (s *StateType) String() string {
// 	var str string
// 	switch s {
// 	case INDETERMINATE, INDETERMINATE_DISABLED:
// 		str = "indeterminate"
// 	case UNCHECKED, UNCHECKED_DISABLED:
// 		c.GetObject().Set("checked", false)
// 		c.GetObject().Set("indeterminate", false)
// 	case CHECKED, CHECKED_DISABLED:
// 		c.GetObject().Set("checked", true)
// 		c.GetObject().Set("indeterminate", false)
// 	}

// 	if s%2 != 0 {
// 		c.GetObject().Set("disabled", true)
// 		return
// 	}

// 	c.GetObject().Set("disabled", false)
// }
