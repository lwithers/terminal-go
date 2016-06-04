package screen

type Symbol struct {
	// rune to draw
	R rune

	// foreground/background colours
	Fg, Bg uint32
}

var (
	EmptySymbol  = Symbol{' ', White, Black}
	PlayerSymbol = Symbol{'@', White, Black}
)
