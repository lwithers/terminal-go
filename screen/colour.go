package screen

func Colour(r, g, b uint8) uint32 {
	return uint32(r) | uint32(g)<<8 | uint32(b)<<16
}

const (
	White = 0xFFFFFF
	Black = 0x000000
)
