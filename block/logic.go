package block

type Color = uint8

var Colors = struct {
	RED   Color
	GREEN Color
	BLUE  Color
	BLACK Color
	WHITE Color
}{
	RED:   0,
	GREEN: 1,
	BLUE:  2,
	BLACK: 3,
	WHITE: 4,
}

type Logic struct {
	X, Y, Width, Height int16
	Color               Color
}
