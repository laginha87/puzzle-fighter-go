package logic

// Game base class
type Game struct {
	board Board
}

func NewGame() Game {
	return Game{
		board: NewBoard(4, 20),
	}
}

func (g *Game) Start() {
	gl := NewGameLoop(10, func(delta float64) {
		println(delta)
	})

	gl.Start()
}
