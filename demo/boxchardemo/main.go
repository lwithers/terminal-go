/*
boxchardemo draws box characters on the screen, with a nice grid.
*/
package main

import (
	"fmt"

	"github.com/lwithers/terminal-go/keyboard"
	"github.com/lwithers/terminal-go/screen"
)

var (
	AddrFg = screen.Colour(128, 255, 128)
	BoxFg  = screen.Colour(96, 96, 96)
	Bg     = screen.Colour(0, 0, 0)
	CharFg = screen.Colour(255, 255, 255)
)

func drawChars(base int) {
	// print characters
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			screen.DrawRune(rune(base+y<<4+x), CharFg, Bg,
				x*2+3, y*2+3)
		}
	}

}

func main() {
	screen.Init()
	keyboard.Init()

	// print character addresses
	for i := 0; i < 16; i++ {
		r := rune(i + '0')
		if i >= 10 {
			r = rune(i + 'A' - 10)
		}
		screen.DrawRune(r, AddrFg, screen.Black,
			i*2+3, 1)
		screen.DrawRune(r, AddrFg, screen.Black,
			1, i*2+3)
	}

	// draw surrounding box
	screen.DrawRune('┌', BoxFg, Bg, 0, 0)
	screen.DrawRune('┐', BoxFg, Bg, 34, 0)
	screen.DrawRune('└', BoxFg, Bg, 0, 34)
	screen.DrawRune('┘', BoxFg, Bg, 34, 34)
	for i := 1; i < 34; i++ {
		screen.DrawRune('─', BoxFg, Bg, i, 0)
		screen.DrawRune('─', BoxFg, Bg, i, 34)
		screen.DrawRune('│', BoxFg, Bg, 0, i)
		screen.DrawRune('│', BoxFg, Bg, 34, i)
	}

	// draw grid
	for y := 0; y < 16; y++ {
		screen.DrawRune('├', BoxFg, Bg, 0, y*2+2)
		screen.DrawRune('┤', BoxFg, Bg, 34, y*2+2)
		for x := 0; x < 16; x++ {
			screen.DrawRune('─', BoxFg, Bg, x*2+1, y*2+2)
			screen.DrawRune('┼', BoxFg, Bg, x*2+2, y*2+2)
			screen.DrawRune('│', BoxFg, Bg, x*2+2, y*2+1)
		}
		screen.DrawRune('─', BoxFg, Bg, 33, y*2+2)
	}
	for x := 0; x < 16; x++ {
		screen.DrawRune('┬', BoxFg, Bg, x*2+2, 0)
		screen.DrawRune('│', BoxFg, Bg, x*2+2, 33)
		screen.DrawRune('┴', BoxFg, Bg, x*2+2, 34)
	}

	screen.DrawString("Press ‘q’ to exit", CharFg, Bg, 1, 36, 0)
	keych := keyboard.StartReader()
	base := 0x2500

MainLoop:
	for {
		screen.DrawString(fmt.Sprintf("Base: 0x%05X", base),
			CharFg, Bg, 1, 37, 0)
		drawChars(base)
		screen.Flush()

		key := <-keych
		switch key {
		case keyboard.Key_Ctrl_C, 'q', 'Q':
			break MainLoop
		case keyboard.Key_PgUp:
			base -= 0x100
		case keyboard.Key_PgDown:
			base += 0x100
		}
	}

	keyboard.Stop()
	screen.Stop()
}
