package screen

// DrawRune modifies the screen buffer by placing the given Unicode rune
// with fg/bg text/background colours at the given x,y coordinates.
func DrawRune(r rune, fg, bg uint32, x, y int) {
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
		buf.dmgY2 = y + 1
	}

	pos := y*buf.w + x
	buf.runes[pos] = r
	buf.fg[pos] = fg
	buf.bg[pos] = bg
}

// Flush writes the contents of the buffer to the screen.
func Flush() {
	buf.drawBuffer()
}

// Clear clears the buffer (so the next flush will display an empty screen with
// a black blackground).
func Clear() {
	buf.clearBuffer()
}
