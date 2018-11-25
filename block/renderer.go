package block

import "github.com/laginha87/puzzle-fighter-go/html"

var (
	scaleX int16 = 50
	scaleY int16 = 50
)

type Renderer struct {
	Logic *Logic
}

func NewRenderer(logic *Logic) *Renderer {
	return &Renderer{
		Logic: logic,
	}
}

var ColorMap = map[Color]string{
	Colors.RED:   "#ff0000",
	Colors.GREEN: "#00ff00",
	Colors.BLUE:  "#0000ff",
	Colors.WHITE: "#f0f0f0",
	Colors.BLACK: "#000000",
}

func (r Renderer) Render(c html.Context2D) {
	c.FillStyle(ColorMap[r.Logic.Color])
	c.FillRect(r.Logic.X*scaleX+2, r.Logic.Y*scaleY+2, r.Logic.Width*scaleX-2, r.Logic.Height*scaleY-2)
}
