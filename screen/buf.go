package screen

import (
	"bytes"
	"os"
	"unicode/utf8"

	"github.com/lwithers/terminal-go/screen/raw"
)

type buffer struct {
	w, h   int
	runes  []rune
	fg, bg []uint32

	dmgX1, dmgY1, dmgX2, dmgY2 int

	tmpbuf []byte
}

var (
	buf buffer
)

func (buf *buffer) clearBuffer() {
	for pos := range buf.runes {
		buf.runes[pos] = ' '
	}
	for pos := range buf.fg {
		buf.fg[pos] = White
	}
	for pos := range buf.bg {
		buf.bg[pos] = Black
	}

	buf.dmgX1 = 0
	buf.dmgY1 = 0
	buf.dmgX2 = buf.w
	buf.dmgY2 = buf.h
}

func (buf *buffer) resetBuffer() {
	buf.w, buf.h = GetSize()
	buf.runes = make([]rune, buf.w*buf.h)
	buf.fg = make([]uint32, buf.w*buf.h)
	buf.bg = make([]uint32, buf.w*buf.h)
	buf.tmpbuf = make([]byte, 8)

	buf.clearBuffer()
	buf.drawBuffer()
}

func (buf *buffer) drawBuffer() {
	cmd := bytes.NewBuffer(nil)

	for y := buf.dmgY1; y < buf.dmgY2; y++ {
		buf.drawBufferLine(y, cmd)
	}

	buf.dmgX1 = buf.w
	buf.dmgX2 = 0
	buf.dmgY1 = buf.h
	buf.dmgY2 = 0

	os.Stdout.Write(cmd.Bytes())
}

func (buf *buffer) drawBufferLine(y int, cmd *bytes.Buffer) {
	var fg1, bg1 uint32 = 0xFFFFFFFF, 0xFFFFFFFF
	pos := y*buf.w + buf.dmgX1

	raw.CursorMove(cmd, buf.dmgX1, y)

	for x := buf.dmgX1; x < buf.dmgX2; x++ {
		if fg1 != buf.fg[pos] {
			fg1 = buf.fg[pos]
			raw.ColourFG(cmd, uint8(fg1), uint8(fg1>>8), uint8(fg1>>16))
		}

		if bg1 != buf.bg[pos] {
			bg1 = buf.bg[pos]
			raw.ColourBG(cmd, uint8(bg1), uint8(bg1>>8), uint8(bg1>>16))
		}

		n := utf8.EncodeRune(buf.tmpbuf, buf.runes[pos])
		cmd.Write(buf.tmpbuf[:n])

		pos++
	}
}
