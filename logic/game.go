package logic

import (
	"github.com/laginha87/puzzle-fighter-go/abstract"
	"github.com/laginha87/puzzle-fighter-go/client"
)

// Game base class
type Game struct {
	logger abstract.Logger
}

func NewGame(logger abstract.Logger) Game {
	return Game{
		logger: logger,
	}
}

func (g *Game) Start() {
	gl := client.NewGameLoop(10, func(delta float64) {
		println(delta)
	})

	gl.Start()
}
