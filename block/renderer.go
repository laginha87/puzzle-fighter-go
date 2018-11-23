package block

import "github.com/laginha87/puzzle-fighter-go/html"

var (
	scaleX int16 = 50
	scaleY int16 = 50
)

type Renderer struct {
	logic Logic
}

func NewRenderer(logic Logic) Renderer {
	return Renderer{
		logic: logic,
	}
}

var Colors = map[Color]string{
	Red:   "#ff0000",
	Green: "#00ff00",
	Blue:  "#0000ff",
	White: "#f0f0f0",
	Black: "#000000",
}

func (r Renderer) Render(c html.Context2D) {
	c.FillStyle(Colors[r.logic.Color])
	c.FillRect(r.logic.X*scaleX+2, r.logic.Y*scaleY+2, r.logic.Width*scaleX-2, r.logic.Height*scaleY-2)
}
