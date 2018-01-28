package textfield

import "github.com/gopherjs/gopherjs/js"

var mdcTextfield = js.Global.Get("MDCTextField")

type props struct {
	Value             string `js:"value"`
	HelperTextContent string `js:"helperTextContent"`
	Disabled          bool   `js:"disabled"`
	Valid             bool   `js:"valid"`
	Required          bool   `js:"required"`
	//TODO Ripple            interface{} `js:"ripple"`
}

type Textfield struct {
	*js.Object
	props
}

// type tf struct {
// 	*js.Object
// 	*Props
// }

func New() *Textfield {
	return &Textfield{Object: mdcTextfield}
}
