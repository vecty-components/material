package linearprogress

import (
	"github.com/vecty-material/material/base"
)

// TODO: fix functions to be js.Func

// afterStart defines missing getters for MDCLinearProgress properties, so that
// we can use our struct fields as one would normally expect.
func (c *LP) afterStart() error {
	o := c.Component()
	err := base.DefineSetGet(c, "determinate",
		func(v interface{}) {
			o.Get("foundation_").Call("setDeterminate", v)
		},
		func() interface{} {
			return o.Get("foundation_").Get("determinate_").Bool()
		},
	)
	if err != nil {
		return err
	}
	err = base.DefineSetGet(c, "progress",
		func(v interface{}) {
			o.Get("foundation_").Call("setProgress", v)
		},
		func() interface{} {
			return o.Get("foundation_").Get("progress_").Float()
		},
	)
	if err != nil {
		return err
	}
	err = base.DefineSetGet(c, "buffer",
		func(v interface{}) {
			vFloat, ok := v.(float64)
			if !ok {
				panic("Unable to set buffer. Unable to parse float.")
			}
			o.Get("foundation_").Call("setBuffer", v)
			c.bufferCache = vFloat
		},
		func() interface{} {
			return c.GetBufferCache()
		},
	)
	if err != nil {
		return err
	}
	err = base.DefineSetGet(c, "reverse",
		func(v interface{}) {
			o.Get("foundation_").Call("setReverse", v)
		},
		func() interface{} {
			return o.Get("foundation_").Get("reverse_").Bool()
		},
	)
	return err
}

// GetBufferCache is a getter function for MDCLinearProgress.buffer
func (lp *LP) GetBufferCache() float64 {
	return lp.bufferCache
}
