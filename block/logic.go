package block

const (
	Red = iota
	Green
	Blue
	Black
	White
)

type Color = int16

type Logic struct {
	X, Y, Width, Height int16
	Color               Color
}
