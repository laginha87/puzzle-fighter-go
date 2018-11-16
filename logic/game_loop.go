package logic

import (
	"runtime"
	"time"
)

type GameLoop struct {
	onUpdate func(float64)
	tickRate time.Duration
	quit     chan bool
}

// NewGameLoop Create new game loop
func NewGameLoop(tickRate time.Duration, onUpdate func(float64)) *GameLoop {
	return &GameLoop{
		onUpdate: onUpdate,
		tickRate: tickRate,
		quit:     make(chan bool),
	}
}

func (gl *GameLoop) startLoop() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	tickInterval := time.Second / gl.tickRate
	timeStart := time.Now().UnixNano()

	ticker := time.NewTicker(tickInterval)

	for {
		select {
		case <-ticker.C:
			now := time.Now().UnixNano()
			// DT in seconds
			delta := float64(now-timeStart) / 1000000000
			timeStart = now
			gl.onUpdate(delta)

		case <-gl.quit:
			ticker.Stop()
		}
	}
}

// GetTickRate get tick rate
func (gl *GameLoop) GetTickRate() time.Duration {
	return gl.tickRate
}

// SetTickRate and restart game loop
func (gl *GameLoop) SetTickRate(tickRate time.Duration) {
	gl.tickRate = tickRate
	gl.Restart()
}

// SetOnUpdate func
func (gl *GameLoop) SetOnUpdate(onUpdate func(float64)) {
	gl.onUpdate = onUpdate
}

// Start game loop
func (gl *GameLoop) Start() {
	go gl.startLoop()
}

// Stop game loop
func (gl *GameLoop) Stop() {
	gl.quit <- true
}

// Restart game loop
func (gl *GameLoop) Restart() {
	gl.Stop()
	gl.Start()
}
