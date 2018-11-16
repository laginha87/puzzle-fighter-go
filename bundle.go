// +build js,wasm

package main

import (
	"syscall/js"

	"github.com/laginha87/puzzle-fighter-go/client"
	"github.com/laginha87/puzzle-fighter-go/logic"
)

var (
	width  int
	height int
	gl     js.Value
)

func main() {
	// doc := js.Global().Get("document")
	// canvasEl := doc.Call("getElementById", "gocanvas")
	// width = doc.Get("body").Get("clientWidth").Int()
	// height = doc.Get("body").Get("clientHeight").Int()
	// canvasEl.Set("width", width)
	// canvasEl.Set("height", height)
	// gl = canvasEl.Call("getContext", "webgl")
	// gl.Call("clearColor", 0.0, 0.0, 0.0, 1.0)
	// gl.Call("clear", gl.Get("COLOR_BUFFER_BIT"))
	logger := client.NewLogger()
	game := logic.NewGame(logger)
	logger.Print("Starting Game Loop")
	game.Start()
	<-make(chan int, 1)
}
