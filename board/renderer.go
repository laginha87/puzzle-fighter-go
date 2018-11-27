package board

import (
	"github.com/laginha87/puzzle-fighter-go/block"
	"github.com/laginha87/puzzle-fighter-go/html"
)

type Renderer struct {
	logic  *Logic
	blocks map[*block.Logic]*block.Renderer
	origin []int16
}

func NewRenderer(logic *Logic, origin []int16) *Renderer {
	renderer := &Renderer{logic: logic, blocks: make(map[*block.Logic]*block.Renderer), origin: origin}
	logic.On(Events.BLOCK_INSERTED, renderer.onInsert)
	return renderer
}

func (r *Renderer) Render(c html.Context2D) {
	c.Translate(r.origin[0], r.origin[1])
	for _, renderer := range r.blocks {
		renderer.Render(c)
	}
	c.Translate(-r.origin[0], -r.origin[1])
}

func (r *Renderer) onInsert(payload interface{}) {
	logic := payload.(*block.Logic)
	renderer := block.NewRenderer(logic)
	r.blocks[logic] = renderer
}

func (r *Renderer) onRemove(payload interface{}) {
	logic := payload.(*block.Logic)
	r.blocks[logic] = nil
}
