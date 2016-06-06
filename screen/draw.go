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

// DrawString modifies the screen buffer by writing the given UTF-8 string with
// fg/bg text/background colours at the given x, y coordinates. The string is
// length limited to the given max runes; pass 0 or negative to limit at the
// edge of the screen.
func DrawString(s string, fg, bg uint32, x, y, max int) {
	pos := y*buf.w + x
	if max <= 0 || max > buf.w-x {
		max = buf.w - x
	}

	var count int
	for _, r := range s {
		buf.runes[pos+count] = r
		buf.fg[pos+count] = fg
		buf.bg[pos+count] = bg
		count++
		if count >= max {
			break
		}
	}

	if x < buf.dmgX1 {
		buf.dmgX1 = x
	}
	if x+count >= buf.dmgX2 {
		buf.dmgX2 = x + count + 1
	}
	if y < buf.dmgY1 {
		buf.dmgY1 = y
	}
	if y >= buf.dmgY2 {
		buf.dmgY2 = y + 1
	}
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
