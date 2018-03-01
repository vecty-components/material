package base

type IDer interface {
	ID() string
}

type IDSetter interface {
	SetID(id string)
}
