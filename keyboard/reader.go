package keyboard

import (
	"time"
	"unicode/utf8"

	"github.com/lwithers/terminal-go/keyboard/kbraw"
)

type readerState int

const (
	keySequenceTimeout = time.Millisecond * 20
)

func StartReader() chan KeyPress {
	ch := make(chan KeyPress)
	go reader(ch)
	return ch
}

func bufConsume(buf []byte, count int) []byte {
	copy(buf, buf[count:])
	return buf[:len(buf)-count]
}

func reader(outch chan KeyPress) {
	buf := make([]byte, 0, 8)
	timeout := time.NewTimer(keySequenceTimeout)

	bch := kbraw.StartReader()

	for {
	charLoop:
		for len(buf) > 0 {
			switch buf[0] {
			case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
				14, 15, 16, 17, 18, 19, 20, 21, 22,
				23, 24, 25, 26:
				outch <- KeyPress(Key_Ctrl_A -
					KeyPress(buf[0]-1))
				buf = bufConsume(buf, 1)
			case 13:
				outch <- KeyPress(Key_Enter)
				buf = bufConsume(buf, 1)

			case 27:
				r, n := decodeEscapeSequence(buf)
				if n == 0 {
					timeout.Reset(keySequenceTimeout)
					break charLoop
				}
				outch <- KeyPress(r)
				buf = bufConsume(buf, n)

			case 127:
				outch <- KeyPress(Key_Backspace)
				buf = bufConsume(buf, 1)

			default:
				if !utf8.FullRune(buf) {
					timeout.Reset(keySequenceTimeout)
					break charLoop
				}
				r, n := utf8.DecodeRune(buf)
				outch <- KeyPress(r)
				buf = bufConsume(buf, n)
			}
		}

		select {
		case b := <-bch:
			buf = append(buf, b)
		case <-timeout.C:
			if len(buf) > 0 {
				if buf[0] == 27 {
					outch <- Key_Escape
				} else {
					outch <- KeyPress(utf8.RuneError)
				}
				buf = bufConsume(buf, 1)

				if len(buf) > 0 {
					timeout.Reset(keySequenceTimeout)
				}
			}
			break
		}
	}
}

func escapeFullSequence(buf []byte) bool {
	return true
}
