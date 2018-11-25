package board

import (
	"github.com/laginha87/puzzle-fighter-go/block"
	"github.com/laginha87/puzzle-fighter-go/html"
)

type Renderer struct {
	logic  *Logic
	blocks []*block.Renderer
	max    uint8
}

func NewRenderer(logic *Logic) *Renderer {
	renderer := &Renderer{logic: logic, blocks: make([]*block.Renderer, 200), max: 0}
	logic.On(Events.BLOCK_INSERTED, renderer.onInsert)
	return renderer
}

func (r *Renderer) Render(c html.Context2D) {
	for index := uint8(0); index < r.max; index++ {
		r.blocks[index].Render(c)
	}
}

func (r *Renderer) onInsert(payload interface{}) {
	logic := payload.(*block.Logic)
	renderer := block.NewRenderer(logic)
	r.blocks[r.max] = renderer
	r.max++
}

func (r *Renderer) onRemove(payload interface{}) {
	logic := payload.(*block.Logic)
	r.blocks[r.max] = renderer
	r.max++
}
