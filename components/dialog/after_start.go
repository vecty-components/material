package dialog

import (
	"syscall/js"

	"github.com/vecty-material/material/components/base"
)

func (c *D) afterStart() (err error) {
	proto := js.Global().Get("Object").Call("getPrototypeOf", c)
	ogGetter := js.Global().Get("Object").Call("getOwnPropertyDescriptor",
		proto, "open").Get("get")
	return base.DefineSetGet(c, "open",
		func(v interface{}) {
			b, ok := v.(bool)
			if !ok {
				print("WARNING: Non-bool set on dialog.Open")
				return
			}
			switch b {
			case true:
				err = c.setOpen()
				if err != nil {
					print(err)
				}
			case false:
				err = c.setClose()
				if err != nil {
					print(err)
				}
			}
		},
		ogGetter,
	)
}
