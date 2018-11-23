// +build js,wasm

package main

import (
	"github.com/laginha87/puzzle-fighter-go/block"
	"github.com/laginha87/puzzle-fighter-go/html"
)

func main() {
	canvas := html.NewCanvas()
	canvas.Init()
	context := canvas.Get2DContext()

	for i2 := 0; i2 < 10; i2++ {
		for index := int16(0); index < 5; index++ {
			b := block.NewRenderer(block.Logic{X: index, Y: i2, Width: 1, Height: 1, Color: index})
			b.Render(context)
		}
	}

}
