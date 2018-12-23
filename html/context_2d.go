// +build js,wasm

package html

import (
	"syscall/js"
)

type Context2D struct {
	context js.Value
	width   int
	height  int
}

func NewContext2D(el js.Value, width, height int) Context2D {
	return Context2D{
		context: el,
		width:   width,
		height:  height,
	}
}

func (c Context2D) FillRect(x, y, w, h int16) {
	c.context.Call("fillRect", x, y, w, h)
}

func (c Context2D) FillColor(color string) {
	c.context.Set("fillColor", color)
}

func (c Context2D) FillStyle(color string) {
	c.context.Set("fillStyle", color)
}

func (c Context2D) Translate(x, y int16) {
	c.context.Call("translate", x, y)
}

func (c Context2D) Clear() {
	c.context.Call("clearRect", 0, 0, c.width, c.height)
}
