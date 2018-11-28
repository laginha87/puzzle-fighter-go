// +build js,wasm

package main

import (
	"github.com/laginha87/puzzle-fighter-go/lib"
	"github.com/laginha87/puzzle-fighter-go/logic"
)

func main() {
	inputHandler := lib.NewInputHandler()

	inputHandler.On("KeyA", func() { println("PRESSED IT") })
	gl := logic.NewGameLoop(10, func(delta float64) {
		inputHandler.CheckInputs()
	})

	gl.Start()
	<-make(chan int, 1)
}
