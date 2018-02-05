package checkbox // import "agamigo.io/material/component/checkbox"

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
