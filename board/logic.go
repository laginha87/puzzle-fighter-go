package board

import (
	"github.com/laginha87/puzzle-fighter-go/block"
	"github.com/laginha87/puzzle-fighter-go/lib"
)

type Logic struct {
	handler *lib.EventHandler
	grid    [][]*block.Logic
}

func NewLogic(width, height int) Logic {
	var grid = make([][]*block.Logic, height)
	for i := range grid {
		grid[i] = make([]*block.Logic, width)
	}
	return Logic{
		grid:    grid,
		handler: lib.NewEventHandler(1),
	}
}

//Empty checks if board is empty at x,y
func (b *Logic) Empty(x, y int) bool {
	return b.grid[x][y] == nil
}

//Get returns the block at x,y
func (b *Logic) Get(x, y int) *block.Logic {
	return b.grid[x][y]
}

//Insert a block to the grid
func (b *Logic) Insert(block *block.Logic) {
	b.grid[block.X][block.Y] = block
	b.Trigger(Events.BLOCK_INSERTED, block)
}

func (b *Logic) Remove(block *block.Logic) {
	b.grid[block.X][block.Y] = nil
	b.Trigger(Events.BLOCK_REMOVED, block)
}

func (b *Logic) Trigger(event uint8, payload interface{}) {
	b.handler.Trigger(event, payload)
}

func (b *Logic) On(event uint8, callback func(interface{})) {
	b.handler.On(event, callback)
}

type event = uint8

//Events that a board has
var Events = struct {
	BLOCK_INSERTED event
	BLOCK_REMOVED  event
}{
	BLOCK_INSERTED: 0,
	BLOCK_REMOVED:  1,
}
