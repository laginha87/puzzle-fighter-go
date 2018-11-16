package abstract

import "time"

// Loop loop interface
type Loop interface {
	GetTickRate() time.Duration

	// SetTickRate and restart game loop
	SetTickRate(tickRate time.Duration)

	// SetOnUpdate func
	SetOnUpdate(onUpdate func(float64))

	// Start game loop
	Start()

	// Stop game loop
	Stop()

	// Restart game loop
	Restart()
}
