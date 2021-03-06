// +build js,wasm

package html

import (
	"syscall/js"
)

type Canvas struct {
	element       js.Value
	width, height int
}

func NewCanvas() *Canvas {
	return &Canvas{
		element: Document.getElementById("gocanvas"),
	}
}

func (c *Canvas) Init() {
	width := Document.doc.Get("body").Get("clientWidth").Int()
	height := Document.doc.Get("body").Get("clientHeight").Int()
	c.element.Set("width", width)
	c.element.Set("height", height)
	c.width = width
	c.height = height
}

func (c *Canvas) Get2DContext() Context2D {
	return NewContext2D(c.element.Call("getContext", "2d"), c.width, c.height)
}
