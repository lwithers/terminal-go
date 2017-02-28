package screen

import (
	"bytes"
	"os"
	"unicode/utf8"

	"github.com/lwithers/terminal-go/screen/raw"
)

// DrawRune draws a rune on the root window.
func DrawRune(r rune, fg, bg uint32, x, y int) {
	RootWindow.DrawRune(r, fg, bg, x, y)
}

// DrawRune modifies the window by placing the given Unicode rune
// with fg/bg text/background colours at the given x,y coordinates.
func (w *Window) DrawRune(r rune, fg, bg uint32, x, y int) {
	if x < 0 || x >= w.w {
		return
	}
	if y < 0 || y >= w.h {
		return
	}
	w.damaged(x, y, x+1, y+1)

	pos := y*w.w + x
	w.runes[pos] = r
	w.fg[pos] = fg
	w.bg[pos] = bg
}

// DrawRunes is an optimised version of DrawRune that writes multiple runes to
// the screen.
func (w *Window) DrawRunes(r []rune, fg, bg uint32, x, y int) {
	switch {
	case -x >= len(r), x >= w.w, y < 0, y >= w.h:
		return
	case x < 0:
		r = r[-x:]
		x = 0
	}
	w.damaged(x, y, x+len(r)+1, y+1)

	pos := y*w.w + x
	copy(w.runes[pos:], r)
	for i := 0; i < len(r); i++ {
		w.fg[pos+i] = fg
	}
	for i := 0; i < len(r); i++ {
		w.bg[pos+i] = bg
	}
}

// DrawString draws a string on the root window.
func DrawString(s string, fg, bg uint32, x, y, max int) {
	RootWindow.DrawString(s, fg, bg, x, y, max)
}

// DrawString modifies the window by writing the given UTF-8 string with
// fg/bg text/background colours at the given x, y coordinates. The string is
// length limited to the given max runes; pass 0 or negative to limit at the
// edge of the screen.
func (w *Window) DrawString(s string, fg, bg uint32, x, y, max int) {
	if x < 0 || x >= w.h {
		return
	}
	if y < 0 || y >= w.h {
		return
	}
	pos := y*w.w + x
	if max <= 0 || max > w.w-x {
		max = w.w - x
	}

	var count int
	for _, r := range s {
		w.runes[pos+count] = r
		w.fg[pos+count] = fg
		w.bg[pos+count] = bg
		count++
		if count >= max {
			break
		}
	}

	w.damaged(x, y, x+count+1, y+1)
}

// Flush will cause all window output to be updated.
func Flush() {
	w := RootWindow
	cmd := new(bytes.Buffer)
	for y := w.dy1; y < w.dy2; y++ {
		flushLine(y, cmd)
	}
	os.Stdout.Write(cmd.Bytes())
	w.clearDamage()
}

func flushLine(y int, cmd *bytes.Buffer) {
	w := RootWindow
	var (
		tmpbuf   [8]byte
		fg1, bg1 uint32 = 0xFFFFFFFF, 0xFFFFFFFF
	)

	raw.CursorMove(cmd, w.dx1, y)

	pos := y*w.w + w.dx1
	for x := w.dx1; x < w.dx2; x++ {
		r, fg, bg := w.resolveChar(x, y)

		if fg1 != fg {
			fg1 = fg
			raw.ColourFG(cmd, uint8(fg1), uint8(fg1>>8), uint8(fg1>>16))
		}

		if bg1 != bg {
			bg1 = bg
			raw.ColourBG(cmd, uint8(bg1), uint8(bg1>>8), uint8(bg1>>16))
		}

		n := utf8.EncodeRune(tmpbuf[:], r)
		cmd.Write(tmpbuf[:n])

		pos++
	}
}

// Clear clears the root window. The next flush will draw the root window using
// the default Background colour.
func Clear() {
	RootWindow.Clear()
}
