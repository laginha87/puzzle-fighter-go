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

	for i2 := int16(0); i2 < 10; i2 += 2 {
		for index := int16(0); index < 5; index++ {
			b := block.NewRenderer(&block.Logic{X: index * 2, Y: i2, Width: 2, Height: 2, Color: uint8(index)})
			b.Render(context)
		}
	}

}
