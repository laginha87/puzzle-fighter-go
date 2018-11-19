package logic

// Board The Board
type Board struct {
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

//Empty checks if board is empty at x,y
func (b *Board) Empty(x, y int) bool {
	return b.grid[x][y] == nil
}

//Get returns the block at x,y
func (b *Board) Get(x, y int) *Block {
	return b.grid[x][y]
}

//Set sets the block at x,y
func (b *Board) Set(x, y int, block *Block) {
	b.grid[x][y] = block
}
