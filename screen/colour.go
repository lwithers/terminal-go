package screen

// Colour returns a colour code suitable for use with the given RGB
// values.
func Colour(r, g, b uint8) uint32 {
	return uint32(r) | uint32(g)<<8 | uint32(b)<<16
}

const (
	White = 0xFFFFFF
	Black = 0x000000
)

var (
	// Background is the default background colour for newly-cleared
	// screens/windows.
	Background uint32 = Black //TODO: change to white
)
