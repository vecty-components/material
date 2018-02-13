package linearprogress

import (
	"agamigo.io/gojs"
	"github.com/gopherjs/gopherjs/js"
)

// AfterStart implements the material.AfterStarter interface. It defines missing
// getters for the MDCLinearProgress properties, so that we can use our struct
// fields as one would normally expect.
func (c *LP) AfterStart() error {
	var err error
	gojs.CatchException(&err)
	o := c.Component()
	js.Global.Get("Object").Call("defineProperty",
		c, "determinate",
		js.M{
			"set": func(v bool) {
				o.Get("foundation_").Call("setDeterminate", v)
			},
			"get": func() bool {
				return o.Get("foundation_").Get("determinate_").Bool()
			},
		},
	)
	js.Global.Get("Object").Call("defineProperty",
		c, "progress",
		js.M{
			"set": func(v float64) {
				o.Get("foundation_").Call("setProgress", v)
			},
			"get": func() float64 {
				return o.Get("foundation_").Get("progress_").Float()
			},
		},
	)
	js.Global.Get("Object").Call("defineProperty",
		c, "buffer",
		js.M{
			"set": func(v float64) {
				o.Get("foundation_").Call("setBuffer", v)
				c.bufferCache = v
			},
			"get": func() float64 {
				return c.GetBufferCache()
			},
		},
	)
	js.Global.Get("Object").Call("defineProperty",
		c, "reverse",
		js.M{
			"set": func(v bool) {
				o.Get("foundation_").Call("setReverse", v)
			},
			"get": func() bool {
				return o.Get("foundation_").Get("reverse_").Bool()
			},
		},
	)
	return err
}

// GetBufferCache is a getter function for MDCLinearProgress.buffer
func (lp *LP) GetBufferCache() float64 {
	return lp.bufferCache
}
