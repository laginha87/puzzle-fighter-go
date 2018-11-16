package client

import (
	"syscall/js"
	"time"
)

// GameLoop loop interface
type GameLoop struct {
	onUpdate   js.Callback
	tickRate   time.Duration
	intervalID int
}

// NewGameLoop creates a new game loop
func NewGameLoop(tickRate time.Duration, onUpdate func(float64)) *GameLoop {
	callback := js.NewCallback(func(v []js.Value) {
		println(len(v))
		onUpdate(v[0].Float())
	})
	return &GameLoop{
		onUpdate: callback,
		tickRate: tickRate,
	}
}

// GetTickRate Comment
func (g *GameLoop) GetTickRate() time.Duration {
	return g.tickRate
}

// SetTickRate and restart game loop
func (g *GameLoop) SetTickRate(tickRate time.Duration) {
	g.tickRate = tickRate
}

// SetOnUpdate func
func (g *GameLoop) SetOnUpdate(onUpdate func(float64)) {
	// g.onUpdate = g
}

// Start game loop
func (g *GameLoop) Start() {
	g.intervalID = js.Global().Call("setInterval", g.onUpdate).Int()
}

// Stop game loop
func (g *GameLoop) Stop() {
	js.Global().Call("clearInterval", g.intervalID)
}

// Restart game loop
func (g *GameLoop) Restart() {

}
