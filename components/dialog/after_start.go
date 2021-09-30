package dialog

import (
	"syscall/js"

	"github.com/vecty-material/material/components/base"
)

func (c *D) afterStart() (err error) {
	proto := js.Global().Get("Object").Call("getPrototypeOf", c.mdc.Value)
	ogGetter := js.Global().Get("Object").Call("getOwnPropertyDescriptor",
		proto, "open").Get("get")

	return base.DefineSetGet(c, "open",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			v := args[0]
			var b, ok bool
			if v.Truthy() {
				ok = true
				b = v.Bool()
			} else {
				ok = false
			}

			if !ok {
				print("WARNING: Non-bool set on dialog.Open")
				return nil
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

			return nil
		}),
		ogGetter,
	)
}
