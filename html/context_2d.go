package html

import (
	"syscall/js"
)

type Context2D struct {
	context js.Value
}

func NewContext2D(el js.Value) Context2D {
	return Context2D{
		context: el,
	}
}

func (c Context2D) Clear() {
	c.context.Call("clear")
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
