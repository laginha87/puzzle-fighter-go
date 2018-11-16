package logic

// Board The Board
type Board = struct {
	grid [][]*Block
}

// NewBoard Creates a new board
func NewBoard(width, height int) Board {
	var grid = make([][]*Block, height)
	for i := range grid {
		grid[i] = make([]*Block, width)
	}
	return Board{
		grid: grid,
	}
}
