package screen

func Draw(sym Symbol, x, y int) {
	if x < buf.dmgX1 {
		buf.dmgX1 = x
	}
	if x >= buf.dmgX2 {
		buf.dmgX2 = x + 1
	}
	if y < buf.dmgY1 {
		buf.dmgY1 = y
	}
	if y >= buf.dmgY2 {
		buf.dmgY2 = y
	}

	pos := y*buf.w + x
	buf.runes[pos] = sym.R
	buf.fg[pos] = sym.Fg
	buf.bg[pos] = sym.Bg
}

func Flush() {
	buf.drawBuffer()
}

func Clear() {
	buf.clearBuffer()
}
