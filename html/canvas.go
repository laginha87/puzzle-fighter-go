// +build js,wasm

package html

import (
	"syscall/js"
)

type Canvas struct {
	element js.Value
}

func NewCanvas() Canvas {
	return Canvas{
		element: document.getElementById("gocanvas"),
	}
}

func (c Canvas) Init() {
	width := document.doc.Get("body").Get("clientWidth").Int()
	height := document.doc.Get("body").Get("clientHeight").Int()
	c.element.Set("width", width)
	c.element.Set("height", height)
}

func (c Canvas) Get2DContext() Context2D {
	return NewContext2D(c.element.Call("getContext", "2d"))
}
